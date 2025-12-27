package index

import (
    "fmt"
    "os"
    "strings"

    "github.com/NightTrek/atse/internal/langspec"
    sitter "github.com/smacker/go-tree-sitter"
)

// ParserInterface defines the methods needed from parser.Manager
// This prevents circular dependencies
type ParserInterface interface {
	ParseFile(filePath string) (*sitter.Tree, error)
	GetContent(filePath string) ([]byte, error)
	InferLanguage(filePath string) (string, error)
	Query(tree *sitter.Tree, queryString string, langName string, content []byte) ([]*QueryMatch, error)
	ExtractTypeUsages(tree *sitter.Tree, langName string, content []byte, filePath string) []*TypeUsage
	ExtractImports(tree *sitter.Tree, langName string, content []byte, filePath string) []Import
	ExtractCallSites(tree *sitter.Tree, langName string, content []byte, filePath string) []CallSite
	ExtractEventPatterns(tree *sitter.Tree, langName string, content []byte, filePath string) []EventUsage
}

// FileInfo represents a file to be indexed
type FileInfo struct {
	Path     string
	Category string // "production", "test", "generated", "config"
}

// BuildIndex parses all files and builds a comprehensive symbol index
func BuildIndex(parser ParserInterface, files []FileInfo, verbose bool) (*SymbolIndex, error) {
    idx := NewSymbolIndex()
    registry := langspec.DefaultRegistry()

	if verbose {
		fmt.Fprintf(os.Stderr, "\nBuilding symbol index...\n")
	}

	for i, file := range files {
		if verbose && (i+1)%50 == 0 {
			fmt.Fprintf(os.Stderr, "  Indexed: %d/%d files\n", i+1, len(files))
		}

        if err := indexFile(parser, idx, file.Path, registry); err != nil {
            if verbose {
                fmt.Fprintf(os.Stderr, "  Warning: failed to index %s: %v\n", file.Path, err)
            }
			continue
		}

		// Record file modification time for cache invalidation
		info, err := os.Stat(file.Path)
		if err == nil {
			idx.IndexedFiles[file.Path] = info.ModTime()
		}
	}

	if verbose {
		fmt.Fprintf(os.Stderr, "  ✓ Index complete: %d files, %d symbols, %d types, %d imports\n\n",
			len(files), len(idx.Symbols), len(idx.Types), len(idx.Imports))
	}

	return idx, nil
}

// indexFile parses a single file and adds all its symbols to the indices
func indexFile(parser ParserInterface, idx *SymbolIndex, filePath string, registry *langspec.Registry) error {
	// Parse the file
	tree, err := parser.ParseFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to parse file: %w", err)
	}

	content, err := parser.GetContent(filePath)
	if err != nil {
		return fmt.Errorf("failed to get content: %w", err)
	}

	langName, err := parser.InferLanguage(filePath)
	if err != nil {
		return fmt.Errorf("failed to infer language: %w", err)
	}

	// Index symbols (functions, classes, methods)
    if err := indexSymbols(parser, idx, tree, langName, content, filePath, registry); err != nil {
        return fmt.Errorf("failed to index symbols: %w", err)
    }

	// Index type usages
	typeUsages := parser.ExtractTypeUsages(tree, langName, content, filePath)
	for _, usage := range typeUsages {
		idx.Types[usage.TypeName] = append(idx.Types[usage.TypeName], usage)
	}

	// Index imports
	imports := parser.ExtractImports(tree, langName, content, filePath)
	idx.Imports[filePath] = imports

	// Index call sites
	callSites := parser.ExtractCallSites(tree, langName, content, filePath)
	for _, callSite := range callSites {
		idx.Calls[callSite.CalledSymbol] = append(idx.Calls[callSite.CalledSymbol], callSite)
	}

	// Index event patterns
	eventUsages := parser.ExtractEventPatterns(tree, langName, content, filePath)
	for _, event := range eventUsages {
		idx.Events[event.EventName] = append(idx.Events[event.EventName], event)
	}

	return nil
}

// indexSymbols indexes all symbol definitions in a file
func indexSymbols(parser ParserInterface, idx *SymbolIndex, tree *sitter.Tree, langName string, content []byte, filePath string, registry *langspec.Registry) error {
    queries := registry.SymbolQueries(langName)

	for symbolType, queryString := range queries {
		results, err := parser.Query(tree, queryString, langName, content)
		if err != nil {
			continue
		}

		for _, result := range results {
			name := strings.TrimSpace(result.Text)
			if name == "" {
				continue
			}

			node := &SymbolNode{
				Symbol:   name,
				Type:     symbolType,
				FilePath: filePath,
				Line:     result.StartPosition.Row,
			}

			idx.Symbols[name] = append(idx.Symbols[name], node)
		}
	}

	return nil
}

// BuildSymbolQueries creates language-specific queries for finding symbols
func BuildSymbolQueries(langName string) map[string]string {
	queries := make(map[string]string)

	switch langName {
	case "typescript", "javascript":
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
