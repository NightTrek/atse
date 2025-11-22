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
	Short: "Discover feature boundaries via graph traversal (BFS/DFS)",
	Long: `Build a dependency graph starting from a symbol to discover feature boundaries.

This command helps you understand the full scope of a feature by traversing
the call graph and dependency relationships. Use BFS to find all related
components at each level, or DFS to follow deep call chains.

Examples:
  atse graph AuthService ./src --mode bfs --depth 3
  atse graph handleLogin ./src --mode dfs --depth 5
  atse graph UserService ./src --bidirectional
  atse graph authenticate ./src --show-dependencies`,
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

	if verboseFlag {
		fmt.Fprintf(os.Stderr, "Collected %d files to analyze\n", len(files))
	}

	// BUILD SYMBOL INDEX (NEW!)
	index, err := buildSymbolIndex(mgr, files, verboseFlag)
	if err != nil {
		return fmt.Errorf("failed to build symbol index: %w", err)
	}

	// Find the entry point using the index (O(1) lookup!)
	if verboseFlag {
		fmt.Fprintf(os.Stderr, "Finding entry point: %s...\n", symbol)
	}

	entryNodes := index.FindSymbol(symbol)
	if len(entryNodes) == 0 {
		fmt.Printf("Symbol '%s' not found in codebase.\n", symbol)
		return nil
	}

	entryPoint := entryNodes[0] // Use first match
	entryPoint.Level = 0

	if verboseFlag {
		fmt.Fprintf(os.Stderr, "  ✓ Found %s in %s at line %d\n\n", symbol, entryPoint.FilePath, entryPoint.Line+1)
	}

	// Parse connection types
	connTypes, err := parseConnectionTypes(connectionTypesFlag)
	if err != nil {
		return fmt.Errorf("invalid connection types: %w", err)
	}

	// Create connection finders
	finders := createConnectionFinders(connTypes)

	if verboseFlag {
		typeNames := make([]string, len(connTypes))
		for i, t := range connTypes {
			typeNames[i] = string(t)
		}
		fmt.Fprintf(os.Stderr, "Traversing graph (mode=%s, depth=%d, max_symbols=%d, connections=%s)...\n",
			graphModeFlag, graphDepthFlag, maxSymbolsFlag, strings.Join(typeNames, ","))
	}

	// Build the graph using NEW index-based traversal
	graph := NewFeatureGraph(symbol, graphDepthFlag)
	graph.EntryPoint = entryPoint

	// Traverse based on mode (uses index + connection finders)
	switch graphModeFlag {
	case "bfs":
		traverseBFSWithIndex(index, graph, finders, maxSymbolsFlag)
	case "dfs":
		traverseDFSWithIndex(index, graph, finders, maxSymbolsFlag)
	default:
		return fmt.Errorf("invalid mode: %s (use 'bfs' or 'dfs')", graphModeFlag)
	}

	if verboseFlag {
		fmt.Fprintf(os.Stderr, "\n✓ Graph traversal complete: %d symbols found\n\n", len(graph.Nodes))
	}

	// Format and output
	formatted := captureGraph(graph, verboseFlag)

	// Capture peak and end memory
	peakMem := captureMemoryStats()
	endMem := peakMem

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
		EndMemoryBytes:   endMem,
		FilesProcessed:   len(files),
		ResultsCount:     len(graph.Nodes),
	}

	logMetrics(metricsConfig, formatted)
	logBenchmark(metricsConfig, formatted)

	fmt.Print(formatted)

	return nil
}

// GraphNode represents a node in the feature graph
type GraphNode struct {
	Symbol   string
	Type     string // "function", "class", "method"
	FilePath string
	Line     uint32
	Level    int // Distance from entry point
}

// FeatureGraph represents the discovered feature boundary
type FeatureGraph struct {
	EntryPoint   *GraphNode
	Nodes        map[string]*GraphNode // Key: symbol name
	Dependencies map[string][]string   // Key: from symbol, Value: to symbols
	Dependents   map[string][]string   // Key: to symbol, Value: from symbols
	MaxDepth     int
}

// NewFeatureGraph creates a new feature graph
func NewFeatureGraph(entrySymbol string, maxDepth int) *FeatureGraph {
	return &FeatureGraph{
		Nodes:        make(map[string]*GraphNode),
		Dependencies: make(map[string][]string),
		Dependents:   make(map[string][]string),
		MaxDepth:     maxDepth,
	}
}

// findEntryPoint locates the starting symbol in the codebase
func findEntryPoint(mgr *parser.Manager, files []util.FileMatch, symbol string) *GraphNode {
	for _, file := range files {
		tree, err := mgr.ParseFile(file.Path)
		if err != nil {
			continue
		}

		content, err := mgr.GetContent(file.Path)
		if err != nil {
			continue
		}

		langName, err := mgr.InferLanguage(file.Path)
		if err != nil {
			continue
		}

		// Search for the symbol
		queries := buildSymbolQueries(langName)
		for symbolType, queryString := range queries {
			results, err := mgr.Query(tree, queryString, langName, content)
			if err != nil {
				continue
			}

			for _, result := range results {
				name := strings.TrimSpace(result.Text)
				if name == symbol {
					return &GraphNode{
						Symbol:   name,
						Type:     symbolType,
						FilePath: file.Path,
						Line:     result.StartPosition.Row,
						Level:    0,
					}
				}
			}
		}
	}

	return nil
}

// traverseBFSWithIndex performs BFS using the symbol index
func traverseBFSWithIndex(index *SymbolIndex, graph *FeatureGraph, finders []ConnectionFinder, maxSymbols int) {
	visited := make(map[string]bool)
	queue := []*GraphNode{graph.EntryPoint}
	visited[graph.EntryPoint.Symbol] = true
	graph.Nodes[graph.EntryPoint.Symbol] = graph.EntryPoint

	if verboseFlag {
		fmt.Fprintf(os.Stderr, "\nLevel 0 (Entry Point): %s\n", graph.EntryPoint.Symbol)
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// Check max symbols limit
		if maxSymbols > 0 && len(graph.Nodes) >= maxSymbols {
			if verboseFlag {
				fmt.Fprintf(os.Stderr, "\n⚠️  Reached max symbols limit (%d)\n", maxSymbols)
			}
			break
		}

		// Stop if we've reached max depth
		if current.Level >= graph.MaxDepth {
			continue
		}

		if verboseFlag {
			fmt.Fprintf(os.Stderr, "  Processing: %s (level %d)\n", current.Symbol, current.Level)
		}

		// Find connections using all finders
		for _, finder := range finders {
			connections := finder.FindConnections(current, index)

			for _, conn := range connections {
				if !visited[conn.To.Symbol] {
					conn.To.Level = current.Level + 1
					visited[conn.To.Symbol] = true
					graph.Nodes[conn.To.Symbol] = conn.To
					queue = append(queue, conn.To)

					// Record edge
					graph.Dependencies[current.Symbol] = append(graph.Dependencies[current.Symbol], conn.To.Symbol)
				}
			}
		}
	}
}

// traverseDFSWithIndex performs DFS using the symbol index
func traverseDFSWithIndex(index *SymbolIndex, graph *FeatureGraph, finders []ConnectionFinder, maxSymbols int) {
	visited := make(map[string]bool)
	graph.Nodes[graph.EntryPoint.Symbol] = graph.EntryPoint

	if verboseFlag {
		fmt.Fprintf(os.Stderr, "\nStarting DFS from: %s\n", graph.EntryPoint.Symbol)
	}

	var dfs func(*GraphNode)
	dfs = func(current *GraphNode) {
		visited[current.Symbol] = true

		// Check max symbols limit
		if maxSymbols > 0 && len(graph.Nodes) >= maxSymbols {
			if verboseFlag {
				fmt.Fprintf(os.Stderr, "\n⚠️  Reached max symbols limit (%d)\n", maxSymbols)
			}
			return
		}

		// Stop if we've reached max depth
		if current.Level >= graph.MaxDepth {
			return
		}

		if verboseFlag {
			fmt.Fprintf(os.Stderr, "  Processing: %s (level %d)\n", current.Symbol, current.Level)
		}

		// Find connections using all finders
		for _, finder := range finders {
			connections := finder.FindConnections(current, index)

			for _, conn := range connections {
				if !visited[conn.To.Symbol] {
					conn.To.Level = current.Level + 1
					graph.Nodes[conn.To.Symbol] = conn.To
					graph.Dependencies[current.Symbol] = append(graph.Dependencies[current.Symbol], conn.To.Symbol)
					dfs(conn.To)
				}
			}
		}
	}

	dfs(graph.EntryPoint)
}

// traverseBFS performs breadth-first search to discover feature boundaries (OLD - DEPRECATED)
func traverseBFS(mgr *parser.Manager, files []util.FileMatch, graph *FeatureGraph, bidirectional bool) {
	visited := make(map[string]bool)
	queue := []*GraphNode{graph.EntryPoint}
	visited[graph.EntryPoint.Symbol] = true
	graph.Nodes[graph.EntryPoint.Symbol] = graph.EntryPoint

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// Stop if we've reached max depth
		if current.Level >= graph.MaxDepth {
			continue
		}

		// Find dependencies (what this symbol calls/uses)
		deps := findDependencies(mgr, files, current)
		for _, dep := range deps {
			if !visited[dep.Symbol] {
				dep.Level = current.Level + 1
				visited[dep.Symbol] = true
				graph.Nodes[dep.Symbol] = dep
				queue = append(queue, dep)
			}

			// Record edge
			graph.Dependencies[current.Symbol] = append(graph.Dependencies[current.Symbol], dep.Symbol)
		}

		// Find dependents (what calls/uses this symbol) if bidirectional
		if bidirectional {
			dependents := findDependents(mgr, files, current)
			for _, dependent := range dependents {
				if !visited[dependent.Symbol] {
					dependent.Level = current.Level + 1
					visited[dependent.Symbol] = true
					graph.Nodes[dependent.Symbol] = dependent
					queue = append(queue, dependent)
				}

				// Record edge
				graph.Dependents[current.Symbol] = append(graph.Dependents[current.Symbol], dependent.Symbol)
			}
		}
	}
}

// traverseDFS performs depth-first search to follow call chains
func traverseDFS(mgr *parser.Manager, files []util.FileMatch, graph *FeatureGraph, bidirectional bool) {
	visited := make(map[string]bool)
	graph.Nodes[graph.EntryPoint.Symbol] = graph.EntryPoint

	var dfs func(*GraphNode)
	dfs = func(current *GraphNode) {
		visited[current.Symbol] = true

		// Stop if we've reached max depth
		if current.Level >= graph.MaxDepth {
			return
		}

		// Find dependencies
		deps := findDependencies(mgr, files, current)
		for _, dep := range deps {
			if !visited[dep.Symbol] {
				dep.Level = current.Level + 1
				graph.Nodes[dep.Symbol] = dep
				graph.Dependencies[current.Symbol] = append(graph.Dependencies[current.Symbol], dep.Symbol)
				dfs(dep)
			}
		}

		// Find dependents if bidirectional
		if bidirectional {
			dependents := findDependents(mgr, files, current)
			for _, dependent := range dependents {
				if !visited[dependent.Symbol] {
					dependent.Level = current.Level + 1
					graph.Nodes[dependent.Symbol] = dependent
					graph.Dependents[current.Symbol] = append(graph.Dependents[current.Symbol], dependent.Symbol)
					dfs(dependent)
				}
			}
		}
	}

	dfs(graph.EntryPoint)
}

// findDependencies finds what a symbol depends on (calls, imports, etc.)
func findDependencies(mgr *parser.Manager, files []util.FileMatch, node *GraphNode) []*GraphNode {
	var deps []*GraphNode

	// Parse the file containing this symbol
	tree, err := mgr.ParseFile(node.FilePath)
	if err != nil {
		return deps
	}

	content, err := mgr.GetContent(node.FilePath)
	if err != nil {
		return deps
	}

	langName, err := mgr.InferLanguage(node.FilePath)
	if err != nil {
		return deps
	}

	// Find call expressions in this file
	callQuery := buildCallQuery(langName)
	if callQuery == "" {
		return deps
	}

	results, err := mgr.Query(tree, callQuery, langName, content)
	if err != nil {
		return deps
	}

	// Extract called function names
	for _, result := range results {
		calledName := strings.TrimSpace(result.Text)
		if calledName != "" && calledName != node.Symbol {
			// Try to find this symbol in the codebase
			depNode := findSymbolInFiles(mgr, files, calledName)
			if depNode != nil {
				deps = append(deps, depNode)
			}
		}
	}

	return deps
}

// findDependents finds what depends on this symbol (what calls it)
func findDependents(mgr *parser.Manager, files []util.FileMatch, node *GraphNode) []*GraphNode {
	var dependents []*GraphNode

	// Search all files for calls to this symbol
	for _, file := range files {
		tree, err := mgr.ParseFile(file.Path)
		if err != nil {
			continue
		}

		content, err := mgr.GetContent(file.Path)
		if err != nil {
			continue
		}

		langName, err := mgr.InferLanguage(file.Path)
		if err != nil {
			continue
		}

		// Find calls to our symbol
		callQuery := buildSpecificCallQuery(langName, node.Symbol)
		if callQuery == "" {
			continue
		}

		results, err := mgr.Query(tree, callQuery, langName, content)
		if err != nil {
			continue
		}

		if len(results) > 0 {
			// This file calls our symbol - find the containing function
			caller := findContainingFunction(mgr, tree, langName, content, results[0].StartPosition.Row)
			if caller != nil && caller.Symbol != node.Symbol {
				dependents = append(dependents, caller)
			}
		}
	}

	return dependents
}

// findSymbolInFiles searches for a symbol across all files
func findSymbolInFiles(mgr *parser.Manager, files []util.FileMatch, symbol string) *GraphNode {
	for _, file := range files {
		tree, err := mgr.ParseFile(file.Path)
		if err != nil {
			continue
		}

		content, err := mgr.GetContent(file.Path)
		if err != nil {
			continue
		}

		langName, err := mgr.InferLanguage(file.Path)
		if err != nil {
			continue
		}

		queries := buildSymbolQueries(langName)
		for symbolType, queryString := range queries {
			results, err := mgr.Query(tree, queryString, langName, content)
			if err != nil {
				continue
			}

			for _, result := range results {
				name := strings.TrimSpace(result.Text)
				if name == symbol {
					return &GraphNode{
						Symbol:   name,
						Type:     symbolType,
						FilePath: file.Path,
						Line:     result.StartPosition.Row,
					}
				}
			}
		}
	}

	return nil
}

// findContainingFunction finds the function that contains a given line
func findContainingFunction(mgr *parser.Manager, tree *sitter.Tree, langName string, content []byte, line uint32) *GraphNode {
	queries := buildSymbolQueries(langName)
	functionQuery := queries["function"]
	if functionQuery == "" {
		return nil
	}

	results, err := mgr.Query(tree, functionQuery, langName, content)
	if err != nil {
		return nil
	}

	// Find the function that contains this line
	for _, result := range results {
		if result.StartPosition.Row <= line && result.EndPosition.Row >= line {
			return &GraphNode{
				Symbol: strings.TrimSpace(result.Text),
				Type:   "function",
				Line:   result.StartPosition.Row,
			}
		}
	}

	return nil
}

// buildCallQuery creates a query to find function calls
func buildCallQuery(langName string) string {
	switch langName {
	case "typescript", "javascript":
		return `(call_expression function: (identifier) @name)`
	case "go":
		return `(call_expression function: (identifier) @name)`
	case "python":
		return `(call function: (identifier) @name)`
	default:
		return ""
	}
}

// buildSpecificCallQuery creates a query to find calls to a specific function
func buildSpecificCallQuery(langName, funcName string) string {
	switch langName {
	case "typescript", "javascript":
		return fmt.Sprintf(`(call_expression function: (identifier) @name (#eq? @name "%s"))`, funcName)
	case "go":
		return fmt.Sprintf(`(call_expression function: (identifier) @name (#eq? @name "%s"))`, funcName)
	case "python":
		return fmt.Sprintf(`(call function: (identifier) @name (#eq? @name "%s"))`, funcName)
	default:
		return ""
	}
}

// captureGraph formats the feature graph and returns as string
func captureGraph(graph *FeatureGraph, verbose bool) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Feature Graph for %s (%s, depth=%d):\n\n",
		graph.EntryPoint.Symbol, graphModeFlag, graph.MaxDepth))

	// Group nodes by level
	levelNodes := make(map[int][]*GraphNode)
	for _, node := range graph.Nodes {
		levelNodes[node.Level] = append(levelNodes[node.Level], node)
	}

	// Print by level
	for level := 0; level <= graph.MaxDepth; level++ {
		nodes := levelNodes[level]
		if len(nodes) == 0 {
			continue
		}

		if level == 0 {
			sb.WriteString("Level 0 (Entry Point):\n")
		} else {
			sb.WriteString(fmt.Sprintf("Level %d:\n", level))
		}

		for _, node := range nodes {
			relPath := makeRelativePath(node.FilePath)
			sb.WriteString(fmt.Sprintf("  %s (%s) - %s:%d\n", node.Symbol, node.Type, relPath, node.Line+1))

			if verbose {
				// Show dependencies
				if deps, ok := graph.Dependencies[node.Symbol]; ok && len(deps) > 0 {
					sb.WriteString(fmt.Sprintf("    → calls: %s\n", strings.Join(deps, ", ")))
				}
				// Show dependents
				if dependents, ok := graph.Dependents[node.Symbol]; ok && len(dependents) > 0 {
					sb.WriteString(fmt.Sprintf("    ← called by: %s\n", strings.Join(dependents, ", ")))
				}
			}
		}
		sb.WriteString("\n")
	}

	// Summary
	fileCount := countUniqueFiles(graph)
	sb.WriteString(fmt.Sprintf("Found %d components across %d files\n", len(graph.Nodes), fileCount))

	return sb.String()
}

// formatGraph formats and outputs the feature graph
func formatGraph(graph *FeatureGraph, verbose bool) {
	fmt.Print(captureGraph(graph, verbose))
}

// countUniqueFiles counts unique files in the graph
func countUniqueFiles(graph *FeatureGraph) int {
	files := make(map[string]bool)
	for _, node := range graph.Nodes {
		files[node.FilePath] = true
	}
	return len(files)
}

func init() {
	rootCmd.AddCommand(graphCmd)
	graphCmd.Flags().StringVar(&graphModeFlag, "mode", "bfs", "Traversal mode: bfs (breadth-first) or dfs (depth-first)")
	graphCmd.Flags().IntVar(&graphDepthFlag, "depth", 2, "Maximum traversal depth")
	graphCmd.Flags().BoolVar(&bidirectionalFlag, "bidirectional", false, "Show both dependencies and dependents")
	graphCmd.Flags().BoolVar(&showDependenciesFlag, "show-dependencies", true, "Show what this symbol depends on")
	graphCmd.Flags().BoolVar(&showDependentsFlag, "show-dependents", false, "Show what depends on this symbol")
	graphCmd.Flags().StringVar(&connectionTypesFlag, "connection-types", "calls", "Connection types to follow: calls, types, imports, dataflow, events, or 'all'")
	graphCmd.Flags().IntVar(&maxSymbolsFlag, "max-symbols", 500, "Maximum number of symbols to discover (0 = no limit)")
}
