package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/NightTrek/atse/internal/parser"
	"github.com/NightTrek/atse/internal/util"
	sitter "github.com/smacker/go-tree-sitter"
)

// SymbolIndex provides fast lookup of symbol definitions across all files
type SymbolIndex struct {
	Symbols map[string][]*GraphNode        // Symbol name → locations
	Types   map[string][]*parser.TypeUsage // Type name → usages
	Imports map[string][]parser.Import     // File path → imports
	Calls   map[string][]parser.CallSite   // Function name → call sites
	Events  map[string][]parser.EventUsage // Event name → emit/listen sites
}

// NewSymbolIndex creates a new empty symbol index
func NewSymbolIndex() *SymbolIndex {
	return &SymbolIndex{
		Symbols: make(map[string][]*GraphNode),
		Types:   make(map[string][]*parser.TypeUsage),
		Imports: make(map[string][]parser.Import),
		Calls:   make(map[string][]parser.CallSite),
		Events:  make(map[string][]parser.EventUsage),
	}
}

// buildSymbolIndex parses all files once and builds comprehensive indices
func buildSymbolIndex(mgr *parser.Manager, files []util.FileMatch, verbose bool) (*SymbolIndex, error) {
	idx := NewSymbolIndex()

	if verbose {
		fmt.Fprintf(os.Stderr, "\nBuilding symbol index...\n")
	}

	for i, file := range files {
		if verbose && (i+1)%50 == 0 {
			fmt.Fprintf(os.Stderr, "  Indexed: %d/%d files\n", i+1, len(files))
		}

		if err := idx.indexFile(mgr, file.Path); err != nil {
			if verbose {
				fmt.Fprintf(os.Stderr, "  Warning: failed to index %s: %v\n", file.Path, err)
			}
			continue
		}
	}

	if verbose {
		fmt.Fprintf(os.Stderr, "  ✓ Index complete: %d files, %d symbols, %d types, %d imports\n\n",
			len(files), len(idx.Symbols), len(idx.Types), len(idx.Imports))
	}

	return idx, nil
}

// indexFile parses a single file and adds all its symbols to the indices
func (idx *SymbolIndex) indexFile(mgr *parser.Manager, filePath string) error {
	// Parse the file
	tree, err := mgr.ParseFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to parse file: %w", err)
	}

	content, err := mgr.GetContent(filePath)
	if err != nil {
		return fmt.Errorf("failed to get content: %w", err)
	}

	langName, err := mgr.InferLanguage(filePath)
	if err != nil {
		return fmt.Errorf("failed to infer language: %w", err)
	}

	// Index symbols (functions, classes, methods)
	if err := idx.indexSymbols(mgr, tree, langName, content, filePath); err != nil {
		return fmt.Errorf("failed to index symbols: %w", err)
	}

	// Index type usages
	typeUsages := mgr.ExtractTypeUsages(tree, langName, content, filePath)
	for _, usage := range typeUsages {
		idx.Types[usage.TypeName] = append(idx.Types[usage.TypeName], usage)
	}

	// Index imports
	imports := mgr.ExtractImports(tree, langName, content, filePath)
	idx.Imports[filePath] = imports

	// Index call sites
	callSites := mgr.ExtractCallSites(tree, langName, content, filePath)
	for _, callSite := range callSites {
		idx.Calls[callSite.CalledSymbol] = append(idx.Calls[callSite.CalledSymbol], callSite)
	}

	// Index event patterns
	eventUsages := mgr.ExtractEventPatterns(tree, langName, content, filePath)
	for _, event := range eventUsages {
		idx.Events[event.EventName] = append(idx.Events[event.EventName], event)
	}

	return nil
}

// indexSymbols indexes all symbol definitions in a file
func (idx *SymbolIndex) indexSymbols(mgr *parser.Manager, tree *sitter.Tree, langName string, content []byte, filePath string) error {
	queries := buildSymbolQueries(langName)

	for symbolType, queryString := range queries {
		results, err := mgr.Query(tree, queryString, langName, content)
		if err != nil {
			continue
		}

		for _, result := range results {
			name := strings.TrimSpace(result.Text)
			if name == "" {
				continue
			}

			node := &GraphNode{
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

// FindSymbol performs O(1) lookup for symbol definitions
func (idx *SymbolIndex) FindSymbol(symbolName string) []*GraphNode {
	return idx.Symbols[symbolName]
}

// FindTypeUsages performs O(1) lookup for type usages
func (idx *SymbolIndex) FindTypeUsages(typeName string) []*parser.TypeUsage {
	return idx.Types[typeName]
}

// FindImportsFor performs O(1) lookup for imports in a file
func (idx *SymbolIndex) FindImportsFor(filePath string) []parser.Import {
	return idx.Imports[filePath]
}

// FindCallSites performs O(1) lookup for call sites of a function
func (idx *SymbolIndex) FindCallSites(functionName string) []parser.CallSite {
	return idx.Calls[functionName]
}

// FindEventUsages performs O(1) lookup for event usages
func (idx *SymbolIndex) FindEventUsages(eventName string) []parser.EventUsage {
	return idx.Events[eventName]
}

// GetAllSymbolNames returns all symbol names in the index
func (idx *SymbolIndex) GetAllSymbolNames() []string {
	names := make([]string, 0, len(idx.Symbols))
	for name := range idx.Symbols {
		names = append(names, name)
	}
	return names
}

// GetStats returns statistics about the index
func (idx *SymbolIndex) GetStats() map[string]int {
	totalSymbols := 0
	for _, nodes := range idx.Symbols {
		totalSymbols += len(nodes)
	}

	totalTypes := 0
	for _, usages := range idx.Types {
		totalTypes += len(usages)
	}

	totalCalls := 0
	for _, sites := range idx.Calls {
		totalCalls += len(sites)
	}

	return map[string]int{
		"unique_symbols":   len(idx.Symbols),
		"total_symbols":    totalSymbols,
		"unique_types":     len(idx.Types),
		"total_type_usage": totalTypes,
		"unique_imports":   len(idx.Imports),
		"total_calls":      totalCalls,
		"unique_events":    len(idx.Events),
	}
}
