package cli

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"encoding/json"

	"github.com/NightTrek/atse/internal/output"
	"github.com/NightTrek/atse/internal/parser"
	"github.com/NightTrek/atse/internal/util"
	sitter "github.com/smacker/go-tree-sitter"
	"github.com/spf13/cobra"
)

var (
	searchTypeFlag string
	fuzzyFlag      bool
)

var searchCmd = &cobra.Command{
	Use:   "search <query> [path]",
	Short: "Fuzzy search for symbols (functions, classes, variables)",
	Long: `Search for symbols across the codebase with fuzzy matching support.

This is typically the first command you run to find entry points into a feature.
It searches for functions, classes, and other symbols by name.

Examples:
  atse search auth ./src
  atse search AuthService ./src --type class
  atse search "login" ./src --fuzzy
  atse search handleRequest ./src --format json`,
	Args: cobra.RangeArgs(1, 2),
	RunE: runSearch,
}

func runSearch(cmd *cobra.Command, args []string) error {
	startTime := time.Now()
	startMem := captureMemoryStats()

	query := args[0]
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

	// Search for symbols
	var matches []SymbolMatch
	filesProcessed := 0

	for _, file := range files {
		// Parse the file
		tree, err := mgr.ParseFile(file.Path)
		if err != nil {
			if verboseFlag {
				fmt.Fprintf(cmd.ErrOrStderr(), "Warning: failed to parse %s: %v\n", file.Path, err)
			}
			continue
		}

		// Get file content
		content, err := mgr.GetContent(file.Path)
		if err != nil {
			if verboseFlag {
				fmt.Fprintf(cmd.ErrOrStderr(), "Warning: failed to get content for %s: %v\n", file.Path, err)
			}
			continue
		}

		// Infer language
		langName, err := mgr.InferLanguage(file.Path)
		if err != nil {
			if verboseFlag {
				fmt.Fprintf(cmd.ErrOrStderr(), "Warning: failed to infer language for %s: %v\n", file.Path, err)
			}
			continue
		}

		// Find symbols in this file
		fileMatches := findSymbols(mgr, tree, langName, content, file.Path, query)
		matches = append(matches, fileMatches...)
		filesProcessed++
	}

	// Filter by type if specified
	if searchTypeFlag != "" {
		matches = filterByType(matches, searchTypeFlag)
	}

	// Sort by relevance (exact matches first, then by score)
	sortMatches(matches, query)

	// Apply limit
	if limitFlag > 0 && len(matches) > limitFlag {
		matches = matches[:limitFlag]
	}

	// Format and output
	formatted := captureSearchResults(matches, output.Format(formatFlag), verboseFlag)

	// Capture peak and end memory
	peakMem := captureMemoryStats()
	endMem := peakMem

	metricsConfig := MetricsLogConfig{
		Enabled:          logMetricsFlag,
		LogFile:          metricsLogFile,
		TokenModel:       tokenModelFlag,
		Command:          "search",
		Args:             os.Args[1:],
		Format:           formatFlag,
		ExitCode:         0,
		StartTime:        startTime,
		EndTime:          time.Now(),
		StartMemoryBytes: startMem,
		PeakMemoryBytes:  peakMem,
		EndMemoryBytes:   endMem,
		FilesProcessed:   filesProcessed,
		ResultsCount:     len(matches),
	}

	logMetrics(metricsConfig, formatted)
	logBenchmark(metricsConfig, formatted)

	fmt.Print(formatted)

	return nil
}

// SymbolMatch represents a found symbol
type SymbolMatch struct {
	Name      string
	Type      string // "function", "class", "method", "variable"
	Signature string
	FilePath  string
	Line      uint32
	Column    uint32
	Score     int // Relevance score for sorting
}

// findSymbols searches for symbols in a file
func findSymbols(mgr *parser.Manager, tree *sitter.Tree, langName string, content []byte, filePath string, query string) []SymbolMatch {
	var matches []SymbolMatch

	// Build query for different symbol types based on language
	symbolQueries := buildSymbolQueries(langName)

	for symbolType, queryString := range symbolQueries {
		results, err := mgr.Query(tree, queryString, langName, content)
		if err != nil {
			continue
		}

		for _, result := range results {
			// The result.Text is just the name (from @name capture)
			name := strings.TrimSpace(result.Text)

			// Check if this matches our search query
			if matchesQuery(name, query, fuzzyFlag) {
				// For signature, we need to get more context
				// For now, just use the name
				signature := name
				score := calculateScore(name, query)

				matches = append(matches, SymbolMatch{
					Name:      name,
					Type:      symbolType,
					Signature: signature,
					FilePath:  filePath,
					Line:      result.StartPosition.Row,
					Column:    result.StartPosition.Column,
					Score:     score,
				})
			}
		}
	}

	return matches
}

// buildSymbolQueries creates language-specific queries for finding symbols
func buildSymbolQueries(langName string) map[string]string {
	queries := make(map[string]string)

	switch langName {
	case "typescript", "javascript":
		// Simplified queries that work
		queries["function"] = `(function_declaration name: (identifier) @name)`
		queries["class"] = `(class_declaration name: (type_identifier) @name)`
		queries["method"] = `(method_definition name: (property_identifier) @name)`

	case "go":
		queries["function"] = `(function_declaration name: (identifier) @name)`
		queries["method"] = `(method_declaration name: (field_identifier) @name)`

	case "python":
		queries["function"] = `(function_definition name: (identifier) @name)`
		queries["class"] = `(class_definition name: (identifier) @name)`
	}

	return queries
}

// matchesQuery checks if a symbol name matches the search query
func matchesQuery(name, query string, fuzzy bool) bool {
	nameLower := strings.ToLower(name)
	queryLower := strings.ToLower(query)

	// Exact match
	if nameLower == queryLower {
		return true
	}

	// Prefix match
	if strings.HasPrefix(nameLower, queryLower) {
		return true
	}

	// Contains match
	if strings.Contains(nameLower, queryLower) {
		return true
	}

	// Fuzzy match (if enabled)
	if fuzzy {
		return fuzzyMatch(nameLower, queryLower)
	}

	return false
}

// fuzzyMatch performs simple fuzzy matching
func fuzzyMatch(name, query string) bool {
	// Simple fuzzy: all query chars must appear in order in name
	nameIdx := 0
	for _, queryChar := range query {
		found := false
		for nameIdx < len(name) {
			if rune(name[nameIdx]) == queryChar {
				found = true
				nameIdx++
				break
			}
			nameIdx++
		}
		if !found {
			return false
		}
	}
	return true
}

// calculateScore assigns a relevance score to a match
func calculateScore(name, query string) int {
	nameLower := strings.ToLower(name)
	queryLower := strings.ToLower(query)

	// Exact match: highest score
	if nameLower == queryLower {
		return 1000
	}

	// Prefix match: high score
	if strings.HasPrefix(nameLower, queryLower) {
		return 500
	}

	// Contains match: medium score
	if strings.Contains(nameLower, queryLower) {
		return 250
	}

	// Fuzzy match: lower score
	return 100
}

// extractSignature extracts a readable signature from symbol text
func extractSignature(text, symbolType string) string {
	// For now, return first line (simplified)
	lines := strings.Split(text, "\n")
	if len(lines) > 0 {
		sig := strings.TrimSpace(lines[0])
		// Limit length
		if len(sig) > 80 {
			sig = sig[:77] + "..."
		}
		return sig
	}
	return text
}

// filterByType filters matches by symbol type
func filterByType(matches []SymbolMatch, typeFilter string) []SymbolMatch {
	var filtered []SymbolMatch
	for _, match := range matches {
		if match.Type == typeFilter {
			filtered = append(filtered, match)
		}
	}
	return filtered
}

// sortMatches sorts matches by relevance
func sortMatches(matches []SymbolMatch, query string) {
	sort.Slice(matches, func(i, j int) bool {
		// First by score (descending)
		if matches[i].Score != matches[j].Score {
			return matches[i].Score > matches[j].Score
		}
		// Then by name length (shorter first)
		if len(matches[i].Name) != len(matches[j].Name) {
			return len(matches[i].Name) < len(matches[j].Name)
		}
		// Finally alphabetically
		return matches[i].Name < matches[j].Name
	})
}

// captureSearchResults formats search results and returns as string
func captureSearchResults(matches []SymbolMatch, format output.Format, verbose bool) string {
	if len(matches) == 0 {
		return "No symbols found matching query.\n"
	}

	switch format {
	case output.FormatJSON:
		return captureSearchJSON(matches)
	case output.FormatLocations:
		return captureSearchLocations(matches)
	default:
		return captureSearchText(matches, verbose)
	}
}

// formatSearchResults formats and outputs search results
func formatSearchResults(matches []SymbolMatch, format output.Format, verbose bool) {
	fmt.Print(captureSearchResults(matches, format, verbose))
}

// captureSearchText formats results as human-readable text and returns string
func captureSearchText(matches []SymbolMatch, verbose bool) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Found %d symbol(s):\n\n", len(matches)))

	for _, match := range matches {
		// Make path relative
		relPath := makeRelativePath(match.FilePath)
		sb.WriteString(fmt.Sprintf("%s (%s) - %s:%d\n", match.Name, match.Type, relPath, match.Line+1))

		if verbose && match.Signature != "" {
			sb.WriteString(fmt.Sprintf("  %s\n", match.Signature))
		}
	}
	return sb.String()
}

// captureSearchJSON formats results as JSON and returns string
func captureSearchJSON(matches []SymbolMatch) string {
	// Convert to simple structure
	type JSONMatch struct {
		Name      string `json:"name"`
		Type      string `json:"type"`
		Signature string `json:"signature,omitempty"`
		File      string `json:"file"`
		Line      uint32 `json:"line"`
		Column    uint32 `json:"column"`
	}

	jsonMatches := make([]JSONMatch, len(matches))
	for i, match := range matches {
		jsonMatches[i] = JSONMatch{
			Name:      match.Name,
			Type:      match.Type,
			Signature: match.Signature,
			File:      match.FilePath,
			Line:      match.Line + 1,
			Column:    match.Column,
		}
	}

	data, _ := json.MarshalIndent(jsonMatches, "", "  ")
	return string(data) + "\n"
}

// captureSearchLocations formats results as simple locations and returns string
func captureSearchLocations(matches []SymbolMatch) string {
	var sb strings.Builder
	for _, match := range matches {
		sb.WriteString(fmt.Sprintf("%s:%d:%d\n", match.FilePath, match.Line+1, match.Column))
	}
	return sb.String()
}

// formatSearchText formats results as human-readable text
func formatSearchText(matches []SymbolMatch, verbose bool) {
	fmt.Print(captureSearchText(matches, verbose))
}

// formatSearchJSON formats results as JSON
func formatSearchJSON(matches []SymbolMatch) {
	fmt.Print(captureSearchJSON(matches))
}

// formatSearchLocations formats results as simple locations
func formatSearchLocations(matches []SymbolMatch) {
	fmt.Print(captureSearchLocations(matches))
}

// makeRelativePath converts absolute path to relative if possible
func makeRelativePath(path string) string {
	// Simple implementation - just get the last few components
	parts := strings.Split(path, "/")
	if len(parts) > 3 {
		return strings.Join(parts[len(parts)-3:], "/")
	}
	return path
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringVar(&searchTypeFlag, "type", "", "Filter by symbol type (function, class, method, variable)")
	searchCmd.Flags().BoolVar(&fuzzyFlag, "fuzzy", false, "Enable fuzzy matching")
}
