package index

import (
	"encoding/gob"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/NightTrek/atse/internal/util"
	sitter "github.com/smacker/go-tree-sitter"
)

// GraphNode represents a node in the feature graph
type GraphNode struct {
	Symbol   string
	Type     string // "function", "class", "method"
	FilePath string
	Line     uint32
	Level    int // Distance from entry point
}

// FileMeta stores metadata for cache validation
type FileMeta struct {
	LastModified int64
	Size         int64
}

// Structs moved from parser to avoid cycle

// TypeUsage represents where a type is used
type TypeUsage struct {
	TypeName  string
	UsageType string // "parameter", "return", "field", "variable"
	FilePath  string
	Line      uint32
}

// Import represents an import statement
type Import struct {
	ImportPath string
	Alias      string
	FilePath   string
	Line       uint32
}

// CallSite represents where a function is called
type CallSite struct {
	CallerSymbol   string
	CallerFilePath string
	CallerLine     uint32
	CalledSymbol   string
}

// EventUsage represents event emitter or listener usage
type EventUsage struct {
	EventName string
	Type      string // "emit", "listen", "on"
	FilePath  string
	Line      uint32
}

// QueryMatch represents a single match from a Tree-sitter query
// (Needed for Parser interface if used in Build, though Build uses high level extractors mostly)
// Build uses Query directly for indexSymbols.
type QueryMatch struct {
	Name          string
	Text          string
	StartPosition sitter.Point
	EndPosition   sitter.Point
	StartByte     uint32
	EndByte       uint32
}

// Parser interface defines what the indexer needs from the parser
type Parser interface {
	ParseFile(filePath string) (*sitter.Tree, error)
	GetContent(filePath string) ([]byte, error)
	InferLanguage(filePath string) (string, error)
	Query(tree *sitter.Tree, queryString string, langName string, content []byte) ([]*QueryMatch, error)
	ExtractTypeUsages(tree *sitter.Tree, langName string, content []byte, filePath string) []*TypeUsage
	ExtractImports(tree *sitter.Tree, langName string, content []byte, filePath string) []Import
	ExtractCallSites(tree *sitter.Tree, langName string, content []byte, filePath string) []CallSite
	ExtractEventPatterns(tree *sitter.Tree, langName string, content []byte, filePath string) []EventUsage
}

// SymbolIndex provides fast lookup of symbol definitions across all files
type SymbolIndex struct {
	Symbols  map[string][]*GraphNode // Symbol name → locations
	Types    map[string][]*TypeUsage // Type name → usages
	Imports  map[string][]Import     // File path → imports
	Calls    map[string][]CallSite   // Function name → call sites
	Events   map[string][]EventUsage // Event name → emit/listen sites
	FileMeta map[string]FileMeta     // File path → metadata
}

// NewSymbolIndex creates a new empty symbol index
func NewSymbolIndex() *SymbolIndex {
	return &SymbolIndex{
		Symbols:  make(map[string][]*GraphNode),
		Types:    make(map[string][]*TypeUsage),
		Imports:  make(map[string][]Import),
		Calls:    make(map[string][]CallSite),
		Events:   make(map[string][]EventUsage),
		FileMeta: make(map[string]FileMeta),
	}
}

// Build parses all files once and builds comprehensive indices
func Build(mgr Parser, files []util.FileMatch, verbose bool) (*SymbolIndex, error) {
	idx := NewSymbolIndex()

	if verbose {
		fmt.Fprintf(os.Stderr, "\nBuilding symbol index...\n")
	}

	for i, file := range files {
		if verbose && (i+1)%50 == 0 {
			fmt.Fprintf(os.Stderr, "  Indexed: %d/%d files\n", i+1, len(files))
		}

		// Stat file for metadata
		info, err := os.Stat(file.Path)
		if err != nil {
			if verbose {
				fmt.Fprintf(os.Stderr, "  Warning: failed to stat %s: %v\n", file.Path, err)
			}
			continue
		}
		idx.FileMeta[file.Path] = FileMeta{
			LastModified: info.ModTime().Unix(),
			Size:         info.Size(),
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

// Save persists the index to disk
func (idx *SymbolIndex) Save(path string) error {
	// Create directory if needed
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return fmt.Errorf("failed to create cache directory: %w", err)
	}

	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create cache file: %w", err)
	}
	defer f.Close()

	enc := gob.NewEncoder(f)
	if err := enc.Encode(idx); err != nil {
		return fmt.Errorf("failed to encode index: %w", err)
	}
	return nil
}

// Load retrieves the index from disk
func Load(path string) (*SymbolIndex, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var idx SymbolIndex
	dec := gob.NewDecoder(f)
	if err := dec.Decode(&idx); err != nil {
		return nil, fmt.Errorf("failed to decode index: %w", err)
	}
	return &idx, nil
}

// IsStale checks if the index is outdated compared to current files
func (idx *SymbolIndex) IsStale(files []util.FileMatch) bool {
	// Check if files count differs significantly or any file is outdated.
	// To be strictly correct, we should check every file.
	// If a file is in 'files' but not in 'FileMeta', it's new -> Stale.
	// If a file is in 'files' and has different mtime -> Stale.

	for _, file := range files {
		meta, exists := idx.FileMeta[file.Path]
		if !exists {
			return true // New file found
		}

		info, err := os.Stat(file.Path)
		if err != nil {
			return true // File deleted or inaccessible
		}

		if info.ModTime().Unix() != meta.LastModified || info.Size() != meta.Size {
			return true // Modified
		}
	}

	// Note: This doesn't detect deleted files if 'files' doesn't include them.
	// But if they are deleted, they won't be in 'files' list returned by WalkFiles.
	// If 'idx' has files that are NOT in 'files', it means they were deleted.
	// We could check len(idx.FileMeta) vs len(files).
	if len(idx.FileMeta) != len(files) {
		return true
	}

	return false
}

// indexFile parses a single file and adds all its symbols to the indices
func (idx *SymbolIndex) indexFile(mgr Parser, filePath string) error {
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
func (idx *SymbolIndex) indexSymbols(mgr Parser, tree *sitter.Tree, langName string, content []byte, filePath string) error {
	queries := BuildSymbolQueries(langName)

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
func (idx *SymbolIndex) FindTypeUsages(typeName string) []*TypeUsage {
	return idx.Types[typeName]
}

// FindImportsFor performs O(1) lookup for imports in a file
func (idx *SymbolIndex) FindImportsFor(filePath string) []Import {
	return idx.Imports[filePath]
}

// FindCallSites performs O(1) lookup for call sites of a function
func (idx *SymbolIndex) FindCallSites(functionName string) []CallSite {
	return idx.Calls[functionName]
}

// FindEventUsages performs O(1) lookup for event usages
func (idx *SymbolIndex) FindEventUsages(eventName string) []EventUsage {
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
