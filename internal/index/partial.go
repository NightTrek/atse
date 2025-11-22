package index

import (
	"sync"

	sitter "github.com/smacker/go-tree-sitter"
)

// PartialIndex represents a subset of the codebase that has been parsed
// This is used for the hybrid ripgrep + Tree-sitter architecture
type PartialIndex struct {
	mu sync.RWMutex

	// Parsed files
	ParsedFiles map[string]*FileIndex

	// Quick lookup for symbols found so far
	Symbols map[string][]*SymbolNode
}

// FileIndex stores the parsed data for a single file
type FileIndex struct {
	Path    string
	Tree    *sitter.Tree
	Content []byte
	Symbols []*SymbolNode
}

// NewPartial creates a new empty partial index
func NewPartial() *PartialIndex {
	return &PartialIndex{
		ParsedFiles: make(map[string]*FileIndex),
		Symbols:     make(map[string][]*SymbolNode),
	}
}

// AddFile adds a parsed file to the index
func (idx *PartialIndex) AddFile(path string, tree *sitter.Tree, content []byte, symbols []*SymbolNode) {
	idx.mu.Lock()
	defer idx.mu.Unlock()

	// Store file info
	idx.ParsedFiles[path] = &FileIndex{
		Path:    path,
		Tree:    tree,
		Content: content,
		Symbols: symbols,
	}

	// Index symbols
	for _, sym := range symbols {
		idx.Symbols[sym.Symbol] = append(idx.Symbols[sym.Symbol], sym)
	}
}

// HasFile checks if a file is already in the index
func (idx *PartialIndex) HasFile(path string) bool {
	idx.mu.RLock()
	defer idx.mu.RUnlock()
	_, exists := idx.ParsedFiles[path]
	return exists
}

// FindSymbol looks up a symbol in the partial index
func (idx *PartialIndex) FindSymbol(name string) []*SymbolNode {
	idx.mu.RLock()
	defer idx.mu.RUnlock()

	// Return a copy to avoid race conditions
	if nodes, ok := idx.Symbols[name]; ok {
		result := make([]*SymbolNode, len(nodes))
		copy(result, nodes)
		return result
	}
	return nil
}

// Count returns stats about the index
func (idx *PartialIndex) Count() (files, symbols int) {
	idx.mu.RLock()
	defer idx.mu.RUnlock()
	return len(idx.ParsedFiles), len(idx.Symbols)
}
