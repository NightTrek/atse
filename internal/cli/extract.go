package cli

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/NightTrek/atse/internal/index"
	"github.com/NightTrek/atse/internal/parser"
	"github.com/NightTrek/atse/internal/ripgrep"
	sitter "github.com/smacker/go-tree-sitter"
	"github.com/spf13/cobra"
)

var (
	extractDepthFlag int
	extractModeFlag  string
	outputFileFlag   string
)

var extractCmd = &cobra.Command{
	Use:   "extract <symbol> [path]",
	Short: "Extract full source code for a feature",
	Long: `Extract complete source code for a feature using hybrid discovery.
Uses ripgrep to find the entry point, then traces dependencies to extract
all related code components.

Ideal for:
- Understanding how a feature is implemented across files
- Providing context to AI agents
- Documentation generation

Examples:
  atse extract AuthService ./src
  atse extract handleLogin ./src --depth 3
  atse extract User --output user_feature.txt`,
	Args: cobra.RangeArgs(1, 2),
	RunE: runExtract,
}

func runExtract(cmd *cobra.Command, args []string) error {
	startTime := time.Now()
	startMem := captureMemoryStats()

	symbol := args[0]
	path := "."
	if len(args) > 1 {
		path = args[1]
	}

	// Reuse graph logic to find components
	// This builds the dependency graph using hybrid discovery
	// We'll basically re-implement runGraph logic but stop before printing
	// and instead extract source code.

	// 1. Find entry point
	if verboseFlag {
		fmt.Fprintf(os.Stderr, "🔍 Finding entry point '%s'...\n", symbol)
	}

	rgClient := ripgrep.New()
	if !rgClient.Available() {
		return fmt.Errorf("ripgrep (rg) not found in PATH")
	}

	searchOpts := &ripgrep.SearchOptions{
		Includes: includeFlag,
		Excludes: excludeFlag,
	}

	candidates, err := rgClient.Search(symbol, path, searchOpts)
	if err != nil {
		return err
	}

	if len(candidates) == 0 {
		fmt.Printf("Symbol '%s' not found in codebase.\n", symbol)
		return nil
	}

	// 2. Parse entry files
	mgr := parser.New()
	partialIndex := index.NewPartial()
	var entryPoint *GraphNode

	uniqueFiles := make(map[string]bool)
	for _, c := range candidates {
		if uniqueFiles[c.Path] {
			continue
		}
		uniqueFiles[c.Path] = true

		tree, err := mgr.ParseFile(c.Path)
		if err != nil {
			continue
		}
		content, _ := mgr.GetContent(c.Path)
		langName, _ := mgr.InferLanguage(c.Path)

		symbols := extractSymbolsForFile(mgr, tree, langName, content, c.Path)
		partialIndex.AddFile(c.Path, tree, content, symbols)

		for _, sym := range symbols {
			if sym.Symbol == symbol {
				entryPoint = &GraphNode{
					Symbol:   sym.Symbol,
					Type:     sym.Type,
					FilePath: sym.FilePath,
					Line:     sym.Line,
					Level:    0,
				}
				break
			}
		}
		if entryPoint != nil {
			break
		}
	}

	if entryPoint == nil {
		fmt.Printf("Could not find definition for '%s'.\n", symbol)
		return nil
	}

	// 3. Build graph
	graph := NewFeatureGraph(symbol, extractDepthFlag)
	graph.EntryPoint = entryPoint
	graph.Nodes[entryPoint.Symbol] = entryPoint

	connTypes, _ := parseConnectionTypes("calls") // Default to calls for extract
	finders := createConnectionFinders(connTypes)

	queue := []*GraphNode{entryPoint}
	visited := make(map[string]bool)
	visited[entryPoint.Symbol] = true

	if verboseFlag {
		fmt.Fprintf(os.Stderr, "🕸️  Tracing dependencies (depth=%d)...\n", extractDepthFlag)
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.Level >= extractDepthFlag {
			continue
		}

		for _, finder := range finders {
			connections := finder.FindConnections(current, (*cliSymbolIndexAdapter)(partialIndex))
			for _, conn := range connections {
				if !visited[conn.To.Symbol] {
					conn.To.Level = current.Level + 1
					visited[conn.To.Symbol] = true
					graph.Nodes[conn.To.Symbol] = conn.To
					queue = append(queue, conn.To)
				}
			}
		}
	}

	if verboseFlag {
		fmt.Fprintf(os.Stderr, "  ✓ Found %d components\n", len(graph.Nodes))
	}

	// 4. Extract source code
	extractedCode := extractSourceCode(mgr, graph)

	// Output
	if outputFileFlag != "" {
		if err := os.WriteFile(outputFileFlag, []byte(extractedCode), 0644); err != nil {
			return fmt.Errorf("failed to write output file: %w", err)
		}
		fmt.Printf("Extracted code written to %s\n", outputFileFlag)
	} else {
		fmt.Print(extractedCode)
	}

	// Metrics
	peakMem := captureMemoryStats()
	metricsConfig := MetricsLogConfig{
		Enabled:          logMetricsFlag,
		LogFile:          metricsLogFile,
		TokenModel:       tokenModelFlag,
		Command:          "extract",
		Args:             os.Args[1:],
		Format:           formatFlag,
		ExitCode:         0,
		StartTime:        startTime,
		EndTime:          time.Now(),
		StartMemoryBytes: startMem,
		PeakMemoryBytes:  peakMem,
		EndMemoryBytes:   peakMem,
		FilesProcessed:   len(uniqueFiles),
		ResultsCount:     len(graph.Nodes),
	}
	logMetrics(metricsConfig, extractedCode)
	logBenchmark(metricsConfig, extractedCode)

	return nil
}

func extractSourceCode(mgr *parser.Manager, graph *FeatureGraph) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("# Feature: %s (%d components, %d files)\n\n",
		graph.EntryPoint.Symbol,
		len(graph.Nodes),
		countUniqueFiles(graph)))

	levelNodes := make(map[int][]*GraphNode)
	for _, node := range graph.Nodes {
		levelNodes[node.Level] = append(levelNodes[node.Level], node)
	}

	for level := 0; level <= graph.MaxDepth; level++ {
		nodes := levelNodes[level]
		if len(nodes) == 0 {
			continue
		}

		if level == 0 {
			sb.WriteString("## Entry Point\n\n")
		} else {
			sb.WriteString(fmt.Sprintf("## Level %d\n\n", level))
		}

		fileNodes := make(map[string][]*GraphNode)
		for _, node := range nodes {
			fileNodes[node.FilePath] = append(fileNodes[node.FilePath], node)
		}

		for filePath, fileNodeList := range fileNodes {
			relPath := makeRelativePath(filePath)
			sb.WriteString(fmt.Sprintf("### %s\n\n", relPath))

			content, err := mgr.GetContent(filePath)
			if err != nil {
				continue
			}
			tree, err := mgr.ParseFile(filePath)
			if err != nil {
				continue
			}

			for _, node := range fileNodeList {
				source := extractSymbolSource(tree, content, node)
				sb.WriteString(fmt.Sprintf("```\n%s\n```\n\n", source))
			}
		}
	}

	sb.WriteString("---\n")
	sb.WriteString(fmt.Sprintf("Total: %d components across %d files\n",
		len(graph.Nodes), countUniqueFiles(graph)))

	return sb.String()
}

func extractSymbolSource(tree *sitter.Tree, content []byte, node *GraphNode) string {
	root := tree.RootNode()
	var findNode func(*sitter.Node) *sitter.Node
	findNode = func(n *sitter.Node) *sitter.Node {
		if n.StartPoint().Row == node.Line {
			nodeType := n.Type()
			if (node.Type == "function" && strings.Contains(nodeType, "function")) ||
				(node.Type == "class" && strings.Contains(nodeType, "class")) ||
				(node.Type == "method" && strings.Contains(nodeType, "method")) {
				return n
			}
		}
		for i := 0; i < int(n.ChildCount()); i++ {
			if result := findNode(n.Child(i)); result != nil {
				return result
			}
		}
		return nil
	}

	targetNode := findNode(root)
	if targetNode == nil {
		// Fallback: extract just that line if node not found structurally
		lines := strings.Split(string(content), "\n")
		if int(node.Line) < len(lines) {
			return lines[node.Line]
		}
		return fmt.Sprintf("// Could not find source for %s at line %d", node.Symbol, node.Line+1)
	}

	return targetNode.Content(content)
}

func init() {
	rootCmd.AddCommand(extractCmd)
	extractCmd.Flags().IntVar(&extractDepthFlag, "depth", 2, "Maximum traversal depth")
	extractCmd.Flags().StringVar(&extractModeFlag, "mode", "bfs", "Traversal mode: bfs or dfs")
	extractCmd.Flags().StringVar(&outputFileFlag, "output", "", "Output file (default: stdout)")
}
