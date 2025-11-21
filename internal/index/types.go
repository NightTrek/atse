package index

import (
	"encoding/gob"
	"fmt"
	"os"
	"path/filepath"
	"time"

	sitter "github.com/smacker/go-tree-sitter"
)

// SymbolIndex provides fast lookup of symbol definitions across all files
type SymbolIndex struct {
	Symbols map[string][]*SymbolNode // Symbol name → locations
	Types   map[string][]*TypeUsage  // Type name → usages
	Imports map[string][]Import      // File path → imports
	Calls   map[string][]CallSite    // Function name → call sites
	Events  map[string][]EventUsage  // Event name → emit/listen sites

	// Metadata for cache invalidation
	IndexedFiles map[string]time.Time // File path → last modified time
	CreatedAt    time.Time            // When this index was created
	Version      string               // Index format version
}

// SymbolNode represents a symbol definition location
type SymbolNode struct {
	Symbol   string
	Type     string // "function", "class", "method"
	FilePath string
	Line     uint32
	Level    int // Distance from entry point (for graph traversal)
}

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

// NewSymbolIndex creates a new empty symbol index
func NewSymbolIndex() *SymbolIndex {
	return &SymbolIndex{
		Symbols:      make(map[string][]*SymbolNode),
		Types:        make(map[string][]*TypeUsage),
		Imports:      make(map[string][]Import),
		Calls:        make(map[string][]CallSite),
		Events:       make(map[string][]EventUsage),
		IndexedFiles: make(map[string]time.Time),
		CreatedAt:    time.Now(),
		Version:      "1.0",
	}
}

// FindSymbol performs O(1) lookup for symbol definitions
func (idx *SymbolIndex) FindSymbol(symbolName string) []*SymbolNode {
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

// IsStale checks if the index needs to be rebuilt based on file modification times
func (idx *SymbolIndex) IsStale() (bool, error) {
	for filePath, cachedMtime := range idx.IndexedFiles {
		info, err := os.Stat(filePath)
		if err != nil {
			if os.IsNotExist(err) {
				// File was deleted
				return true, nil
			}
			return true, fmt.Errorf("failed to stat %s: %w", filePath, err)
		}

		if info.ModTime().After(cachedMtime) {
			// File was modified
			return true, nil
		}
	}

	return false, nil
}

// SaveToFile serializes the index to disk using gob encoding
func (idx *SymbolIndex) SaveToFile(cachePath string) error {
	// Ensure directory exists
	dir := filepath.Dir(cachePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create cache directory: %w", err)
	}

	// Create file
	file, err := os.Create(cachePath)
	if err != nil {
		return fmt.Errorf("failed to create cache file: %w", err)
	}
	defer file.Close()

	// Encode
	encoder := gob.NewEncoder(file)
	if err := encoder.Encode(idx); err != nil {
		return fmt.Errorf("failed to encode index: %w", err)
	}

	return nil
}

// LoadFromFile deserializes the index from disk
func LoadFromFile(cachePath string) (*SymbolIndex, error) {
	file, err := os.Open(cachePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open cache file: %w", err)
	}
	defer file.Close()

	var idx SymbolIndex
	decoder := gob.NewDecoder(file)
	if err := decoder.Decode(&idx); err != nil {
		return nil, fmt.Errorf("failed to decode index: %w", err)
	}

	return &idx, nil
}

// GetCachePath returns the default cache path for a project
func GetCachePath(projectRoot string) string {
	// Use .atse directory in project root
	return filepath.Join(projectRoot, ".atse", "index.cache")
}

// GetUserCachePath returns cache path in user's home directory
func GetUserCachePath(projectRoot string) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	// Create a hash of the project path for uniqueness
	projHash := fmt.Sprintf("%x", []byte(projectRoot))[:16]

	return filepath.Join(home, ".cache", "atse", projHash, "index.cache"), nil
}

// QueryMatch represents a single match from a Tree-sitter query
type QueryMatch struct {
	Name          string
	Text          string
	StartPosition sitter.Point
	EndPosition   sitter.Point
	StartByte     uint32
	EndByte       uint32
}
