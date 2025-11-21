package cli

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/NightTrek/atse/internal/parser"
	"github.com/NightTrek/atse/internal/util"
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
	Long: `Extract the complete source code for a feature and its dependencies.

This command uses graph traversal to discover all components of a feature,
then extracts their full source code. Perfect for understanding a feature
or preparing context for an LLM.

Examples:
  atse extract AuthService ./src --depth 2
  atse extract handleLogin ./src --depth 3 --output feature.txt
  atse extract UserService ./src --mode bfs`,
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

	// Create parser manager
	mgr := parser.New()

	// Collect files
	files, err := util.WalkFiles(path, recursiveFlag, includeFlag, excludeFlag, excludeDefaultsFlag)
	if err != nil {
		return fmt.Errorf("failed to collect files: %w", err)
	}

	if len(files) == 0 {
		fmt.Println("No files found matching criteria.")
		return nil
	}

	// Check for --rebuild-index flag
	if rebuildIndexFlag {
		mgr.InvalidateIndex()
	}

	// Build symbol index (or reuse cached one)
	index, err := buildSymbolIndex(mgr, files, verboseFlag)
	if err != nil {
		return fmt.Errorf("failed to build symbol index: %w", err)
	}

	// Find the entry point using the index (O(1) lookup)
	if verboseFlag {
		fmt.Fprintf(os.Stderr, "Finding entry point: %s...\n", symbol)
	}

	entryNodes := index.FindSymbol(symbol)
	if len(entryNodes) == 0 {
		fmt.Printf("Symbol '%s' not found in codebase.\n", symbol)
		return nil
	}

	entryPoint := entryNodes[0]
	entryPoint.Level = 0

	if verboseFlag {
		fmt.Fprintf(os.Stderr, "  ✓ Found %s in %s at line %d\n\n", symbol, entryPoint.FilePath, entryPoint.Line+1)
	}

	// Build the graph using index-based traversal
	graph := NewFeatureGraph(symbol, extractDepthFlag)
	graph.EntryPoint = entryPoint

	// Parse connection types (default to calls)
	connTypes, err := parseConnectionTypes("calls")
	if err != nil {
		return fmt.Errorf("invalid connection types: %w", err)
	}

	// Create connection finders
	finders := createConnectionFinders(connTypes)

	if verboseFlag {
		fmt.Fprintf(os.Stderr, "Traversing graph (mode=%s, depth=%d)...\n", extractModeFlag, extractDepthFlag)
	}

	// Traverse using index (much faster than old method)
	switch extractModeFlag {
	case "bfs":
		traverseBFSWithIndex(index, graph, finders, 0) // 0 = no symbol limit
	case "dfs":
		traverseDFSWithIndex(index, graph, finders, 0)
	default:
		return fmt.Errorf("invalid mode: %s (use 'bfs' or 'dfs')", extractModeFlag)
	}

	if verboseFlag {
		fmt.Fprintf(os.Stderr, "✓ Graph traversal complete: %d symbols found\n\n", len(graph.Nodes))
	}

	// Extract source code for all discovered components
	extractedCode := extractSourceCode(mgr, graph)

	// Output
	if outputFileFlag != "" {
		// TODO: Write to file
		fmt.Printf("Writing to %s...\n", outputFileFlag)
	}

	// Capture peak and end memory
	peakMem := captureMemoryStats()
	endMem := peakMem

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
		EndMemoryBytes:   endMem,
		FilesProcessed:   len(files),
		ResultsCount:     len(graph.Nodes),
	}

	logMetrics(metricsConfig, extractedCode)
	logBenchmark(metricsConfig, extractedCode)

	// Print to stdout
	fmt.Print(extractedCode)

	return nil
}

// extractSourceCode extracts the full source for all nodes in the graph
func extractSourceCode(mgr *parser.Manager, graph *FeatureGraph) string {
	var sb strings.Builder

	// Header
	sb.WriteString(fmt.Sprintf("# Feature: %s (%d components, %d files)\n\n",
		graph.EntryPoint.Symbol,
		len(graph.Nodes),
		countUniqueFiles(graph)))

	// Group nodes by level
	levelNodes := make(map[int][]*GraphNode)
	for _, node := range graph.Nodes {
		levelNodes[node.Level] = append(levelNodes[node.Level], node)
	}

	// Extract by level
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

		// Group by file
		fileNodes := make(map[string][]*GraphNode)
		for _, node := range nodes {
			fileNodes[node.FilePath] = append(fileNodes[node.FilePath], node)
		}

		// Extract source for each file
		for filePath, fileNodeList := range fileNodes {
			relPath := makeRelativePath(filePath)
			sb.WriteString(fmt.Sprintf("### %s\n\n", relPath))

			// Get file content
			content, err := mgr.GetContent(filePath)
			if err != nil {
				sb.WriteString(fmt.Sprintf("// Error reading file: %v\n\n", err))
				continue
			}

			// Parse to get node boundaries
			tree, err := mgr.ParseFile(filePath)
			if err != nil {
				sb.WriteString(fmt.Sprintf("// Error parsing file: %v\n\n", err))
				continue
			}

			// Extract each symbol's source
			for _, node := range fileNodeList {
				source := extractSymbolSource(tree, content, node)
				sb.WriteString(fmt.Sprintf("```\n%s\n```\n\n", source))
			}
		}
	}

	// Summary
	sb.WriteString("---\n")
	sb.WriteString(fmt.Sprintf("Total: %d components across %d files\n",
		len(graph.Nodes), countUniqueFiles(graph)))

	return sb.String()
}

// extractSymbolSource extracts the source code for a specific symbol
func extractSymbolSource(tree *sitter.Tree, content []byte, node *GraphNode) string {
	// Find the node in the tree that matches our symbol
	root := tree.RootNode()

	var findNode func(*sitter.Node) *sitter.Node
	findNode = func(n *sitter.Node) *sitter.Node {
		// Check if this node is at the right line
		if n.StartPoint().Row == node.Line {
			// Check if it's the right type
			nodeType := n.Type()
			if (node.Type == "function" && strings.Contains(nodeType, "function")) ||
				(node.Type == "class" && strings.Contains(nodeType, "class")) ||
				(node.Type == "method" && strings.Contains(nodeType, "method")) {
				return n
			}
		}

		// Recurse through children
		for i := 0; i < int(n.ChildCount()); i++ {
			if result := findNode(n.Child(i)); result != nil {
				return result
			}
		}

		return nil
	}

	targetNode := findNode(root)
	if targetNode == nil {
		return fmt.Sprintf("// Could not find source for %s at line %d", node.Symbol, node.Line+1)
	}

	// Extract the source code
	return targetNode.Content(content)
}

func init() {
	rootCmd.AddCommand(extractCmd)
	extractCmd.Flags().IntVar(&extractDepthFlag, "depth", 2, "Maximum traversal depth")
	extractCmd.Flags().StringVar(&extractModeFlag, "mode", "bfs", "Traversal mode: bfs or dfs")
	extractCmd.Flags().StringVar(&outputFileFlag, "output", "", "Output file (default: stdout)")
}
