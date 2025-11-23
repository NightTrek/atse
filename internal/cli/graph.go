package cli

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/NightTrek/atse/internal/index"
	"github.com/NightTrek/atse/internal/parser"
	"github.com/NightTrek/atse/internal/ripgrep"
	"github.com/spf13/cobra"
)

var (
	graphModeFlag        string
	graphDepthFlag       int
	bidirectionalFlag    bool
	showDependenciesFlag bool
	showDependentsFlag   bool
	connectionTypesFlag  string
	maxSymbolsFlag       int
)

var graphCmd = &cobra.Command{
	Use:   "graph <symbol> [path]",
	Short: "Generate a dependency graph for a symbol",
	Long: `Generate a dependency graph for a symbol using hybrid discovery.
Uses ripgrep to find the entry point, then traces dependencies by lazily
parsing files as they are discovered.

This approach is extremely fast even on large codebases because it only
parses files that are actually part of the dependency graph.

Examples:
  atse graph AuthService ./src
  atse graph validateUser ./src --depth 3
  atse graph User --mode dfs`,
	Args: cobra.RangeArgs(1, 2),
	RunE: runGraph,
}

func runGraph(cmd *cobra.Command, args []string) error {
	startTime := time.Now()
	startMem := captureMemoryStats()

	symbol := args[0]
	path := "."
	if len(args) > 1 {
		path = args[1]
	}

	// 1. Find entry point using ripgrep (Phase 1)
	if verboseFlag {
		fmt.Fprintf(os.Stderr, "🔍 Finding entry point '%s' using ripgrep...\n", symbol)
	}

	rgClient := ripgrep.New()
	if !rgClient.Available() {
		return fmt.Errorf("ripgrep (rg) not found in PATH")
	}

	// We search for the symbol definition specifically if possible, but for now
	// general search is safer to find usage + definition
	searchOpts := &ripgrep.SearchOptions{
		Includes: includeFlag,
		Excludes: excludeFlag,
	}

	candidates, err := rgClient.Search(symbol, path, searchOpts)
	if err != nil {
		return fmt.Errorf("ripgrep failed: %w", err)
	}

	if len(candidates) == 0 {
		fmt.Printf("Symbol '%s' not found in codebase.\n", symbol)
		return nil
	}

	// 2. Parse entry files to find the exact definition node
	mgr := parser.New()
	partialIndex := index.NewPartial()
	var entryPoint *GraphNode

	// We prioritize files where the symbol is likely defined (e.g. lines starting with func/class)
	// But for robust ness, parse unique candidate files until we find a definition
	filesParsed := 0
	uniqueFiles := make(map[string]bool)
	for _, c := range candidates {
		if uniqueFiles[c.Path] {
			continue
		}
		uniqueFiles[c.Path] = true

		// Parse this file
		tree, err := mgr.ParseFile(c.Path)
		if err != nil {
			continue
		}
		content, _ := mgr.GetContent(c.Path)
		langName, _ := mgr.InferLanguage(c.Path)

		symbols := extractSymbolsForFile(mgr, tree, langName, content, c.Path)
		partialIndex.AddFile(c.Path, tree, content, symbols)
		filesParsed++

		// Check if any symbol matches exactly
		for _, sym := range symbols {
			if sym.Symbol == symbol {
				// Found it!
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
		fmt.Printf("Could not find definition for '%s' in %d candidate files.\n", symbol, filesParsed)
		return nil
	}

	if verboseFlag {
		fmt.Fprintf(os.Stderr, "  ✓ Found entry point in %s\n", entryPoint.FilePath)
	}

	// 3. Build graph with lazy expansion
	graph := NewFeatureGraph(symbol, graphDepthFlag)
	graph.EntryPoint = entryPoint
	graph.Nodes[entryPoint.Symbol] = entryPoint

	// Parse connection types
	connTypes, err := parseConnectionTypes(connectionTypesFlag)
	if err != nil {
		return err
	}
	finders := createConnectionFinders(connTypes)

	// Traverse
	if verboseFlag {
		fmt.Fprintf(os.Stderr, "🕸️  Traversing graph (mode=%s, depth=%d)...\n", graphModeFlag, graphDepthFlag)
	}

	queue := []*GraphNode{entryPoint}
	visited := make(map[string]bool)
	visited[entryPoint.Symbol] = true

	processedCount := 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		processedCount++

		if current.Level >= graphDepthFlag {
			continue
		}

		// Find connections
		for _, finder := range finders {
			connections := finder.FindConnections(current, (*cliSymbolIndexAdapter)(partialIndex))

			for _, conn := range connections {
				if !visited[conn.To.Symbol] {
					conn.To.Level = current.Level + 1
					visited[conn.To.Symbol] = true
					graph.Nodes[conn.To.Symbol] = conn.To
					graph.Dependencies[current.Symbol] = append(graph.Dependencies[current.Symbol], conn.To.Symbol)
					queue = append(queue, conn.To)
				}
			}
		}
	}

	if verboseFlag {
		fmt.Fprintf(os.Stderr, "  ✓ Graph complete: %d nodes\n\n", len(graph.Nodes))
	}

	// Capture output
	formatted := captureGraph(graph, verboseFlag)

	// Metrics
	peakMem := captureMemoryStats()
	metricsConfig := MetricsLogConfig{
		Enabled:          logMetricsFlag,
		LogFile:          metricsLogFile,
		TokenModel:       tokenModelFlag,
		Command:          "graph",
		Args:             os.Args[1:],
		Format:           formatFlag,
		ExitCode:         0,
		StartTime:        startTime,
		EndTime:          time.Now(),
		StartMemoryBytes: startMem,
		PeakMemoryBytes:  peakMem,
		EndMemoryBytes:   peakMem,
		FilesProcessed:   filesParsed,
		ResultsCount:     len(graph.Nodes),
	}
	logMetrics(metricsConfig, formatted)
	logBenchmark(metricsConfig, formatted)

	fmt.Print(formatted)
	return nil
}

// Adapter to make PartialIndex look like SymbolIndex to the existing finders
type cliSymbolIndexAdapter index.PartialIndex

// ToLegacyIndex converts partial index to legacy index for compatibility
func (idx *cliSymbolIndexAdapter) ToLegacyIndex() *SymbolIndex {
	legacy := NewSymbolIndex()
	// Copy symbols
	for name, nodes := range idx.Symbols {
		var graphNodes []*GraphNode
		for _, n := range nodes {
			graphNodes = append(graphNodes, &GraphNode{
				Symbol:   n.Symbol,
				Type:     n.Type,
				FilePath: n.FilePath,
				Line:     n.Line,
			})
		}
		legacy.Symbols[name] = graphNodes
	}
	return legacy
}

// FeatureGraph represents the dependency graph
type FeatureGraph struct {
	EntryPoint   *GraphNode
	Nodes        map[string]*GraphNode
	Dependencies map[string][]string
	Dependents   map[string][]string
	MaxDepth     int
}

type GraphNode struct {
	Symbol   string
	Type     string
	FilePath string
	Line     uint32
	Level    int
}

func NewFeatureGraph(entrySymbol string, maxDepth int) *FeatureGraph {
	return &FeatureGraph{
		Nodes:        make(map[string]*GraphNode),
		Dependencies: make(map[string][]string),
		Dependents:   make(map[string][]string),
		MaxDepth:     maxDepth,
	}
}

// Reuse captureGraph from original implementation
func captureGraph(graph *FeatureGraph, verbose bool) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Feature Graph for %s (depth=%d):\n\n",
		graph.EntryPoint.Symbol, graph.MaxDepth))

	nodesByLevel := make(map[int][]*GraphNode)
	for _, node := range graph.Nodes {
		nodesByLevel[node.Level] = append(nodesByLevel[node.Level], node)
	}

	for i := 0; i <= graph.MaxDepth; i++ {
		nodes := nodesByLevel[i]
		if len(nodes) == 0 {
			continue
		}

		if i == 0 {
			sb.WriteString("Level 0 (Entry Point):\n")
		} else {
			sb.WriteString(fmt.Sprintf("Level %d:\n", i))
		}

		for _, node := range nodes {
			sb.WriteString(fmt.Sprintf("  %s (%s) - %s:%d\n",
				node.Symbol, node.Type, makeRelativePath(node.FilePath), node.Line+1))

			if verbose {
				if deps, ok := graph.Dependencies[node.Symbol]; ok {
					sb.WriteString(fmt.Sprintf("    → calls: %s\n", strings.Join(deps, ", ")))
				}
			}
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

func countUniqueFiles(graph *FeatureGraph) int {
	files := make(map[string]bool)
	for _, node := range graph.Nodes {
		files[node.FilePath] = true
	}
	return len(files)
}

func init() {
	rootCmd.AddCommand(graphCmd)
	graphCmd.Flags().StringVar(&graphModeFlag, "mode", "bfs", "Traversal mode: bfs or dfs")
	graphCmd.Flags().IntVar(&graphDepthFlag, "depth", 2, "Maximum traversal depth")
	graphCmd.Flags().BoolVar(&bidirectionalFlag, "bidirectional", false, "Show both dependencies and dependents")
	graphCmd.Flags().StringVar(&connectionTypesFlag, "connection-types", "calls", "Connection types to follow")
	graphCmd.Flags().IntVar(&maxSymbolsFlag, "max-symbols", 500, "Limit symbols")
}
