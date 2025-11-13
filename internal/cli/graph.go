package cli

import (
	"fmt"
	"strings"

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
	symbol := args[0]
	path := "."
	if len(args) > 1 {
		path = args[1]
	}

	// Create parser manager
	mgr := parser.New()

	// Collect files
	files, err := util.WalkFiles(path, recursiveFlag, includeFlag, excludeFlag)
	if err != nil {
		return fmt.Errorf("failed to collect files: %w", err)
	}

	if len(files) == 0 {
		fmt.Println("No files found matching criteria.")
		return nil
	}

	// Build the graph
	graph := NewFeatureGraph(symbol, graphDepthFlag)

	// Find the entry point symbol
	entryPoint := findEntryPoint(mgr, files, symbol)
	if entryPoint == nil {
		fmt.Printf("Symbol '%s' not found in codebase.\n", symbol)
		return nil
	}

	graph.EntryPoint = entryPoint

	// Traverse based on mode
	switch graphModeFlag {
	case "bfs":
		traverseBFS(mgr, files, graph, bidirectionalFlag)
	case "dfs":
		traverseDFS(mgr, files, graph, bidirectionalFlag)
	default:
		return fmt.Errorf("invalid mode: %s (use 'bfs' or 'dfs')", graphModeFlag)
	}

	// Format and output
	formatGraph(graph, verboseFlag)

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

// traverseBFS performs breadth-first search to discover feature boundaries
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

// formatGraph formats and outputs the feature graph
func formatGraph(graph *FeatureGraph, verbose bool) {
	fmt.Printf("Feature Graph for %s (%s, depth=%d):\n\n",
		graph.EntryPoint.Symbol, graphModeFlag, graph.MaxDepth)

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
			fmt.Printf("Level 0 (Entry Point):\n")
		} else {
			fmt.Printf("Level %d:\n", level)
		}

		for _, node := range nodes {
			relPath := makeRelativePath(node.FilePath)
			fmt.Printf("  %s (%s) - %s:%d\n", node.Symbol, node.Type, relPath, node.Line+1)

			if verbose {
				// Show dependencies
				if deps, ok := graph.Dependencies[node.Symbol]; ok && len(deps) > 0 {
					fmt.Printf("    → calls: %s\n", strings.Join(deps, ", "))
				}
				// Show dependents
				if dependents, ok := graph.Dependents[node.Symbol]; ok && len(dependents) > 0 {
					fmt.Printf("    ← called by: %s\n", strings.Join(dependents, ", "))
				}
			}
		}
		fmt.Println()
	}

	// Summary
	fileCount := countUniqueFiles(graph)
	fmt.Printf("Found %d components across %d files\n", len(graph.Nodes), fileCount)
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
}
