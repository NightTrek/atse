package parser

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/golang"
	"github.com/smacker/go-tree-sitter/javascript"
	"github.com/smacker/go-tree-sitter/python"
	"github.com/smacker/go-tree-sitter/typescript/typescript"
)

// Manager handles Tree-sitter parsing, language loading, and caching
type Manager struct {
	languages map[string]*sitter.Language
	trees     map[string]*sitter.Tree
	contents  map[string][]byte // Store file contents for node extraction
	mu        sync.RWMutex
}

// New creates a new parser manager
func New() *Manager {
	return &Manager{
		languages: make(map[string]*sitter.Language),
		trees:     make(map[string]*sitter.Tree),
		contents:  make(map[string][]byte),
	}
}

// LoadLanguage loads a Tree-sitter language grammar
func (m *Manager) LoadLanguage(langName string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Check if already loaded
	if _, exists := m.languages[langName]; exists {
		return nil
	}

	var lang *sitter.Language
	switch langName {
	case "go":
		lang = golang.GetLanguage()
	case "javascript", "js":
		lang = javascript.GetLanguage()
	case "typescript", "ts":
		lang = typescript.GetLanguage()
	case "python", "py":
		lang = python.GetLanguage()
	default:
		return fmt.Errorf("unsupported language: %s", langName)
	}

	m.languages[langName] = lang
	return nil
}

// InferLanguage determines the language from a file extension
func (m *Manager) InferLanguage(filePath string) (string, error) {
	ext := strings.ToLower(filepath.Ext(filePath))
	switch ext {
	case ".go":
		return "go", nil
	case ".js", ".jsx", ".mjs", ".cjs":
		return "javascript", nil
	case ".ts", ".tsx":
		return "typescript", nil
	case ".py":
		return "python", nil
	default:
		return "", fmt.Errorf("unsupported file extension: %s", ext)
	}
}

// ParseFile parses a file and caches the syntax tree
func (m *Manager) ParseFile(filePath string) (*sitter.Tree, error) {
	// Check cache first
	m.mu.RLock()
	if tree, exists := m.trees[filePath]; exists {
		m.mu.RUnlock()
		return tree, nil
	}
	m.mu.RUnlock()

	// Read file content
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	// Infer and load language
	langName, err := m.InferLanguage(filePath)
	if err != nil {
		return nil, err
	}

	if err := m.LoadLanguage(langName); err != nil {
		return nil, err
	}

	// Parse the file
	m.mu.Lock()
	defer m.mu.Unlock()

	lang := m.languages[langName]
	parser := sitter.NewParser()
	parser.SetLanguage(lang)

	tree, err := parser.ParseCtx(context.Background(), nil, content)
	if err != nil {
		return nil, fmt.Errorf("failed to parse file %s: %w", filePath, err)
	}

	if tree == nil {
		return nil, fmt.Errorf("failed to parse file: %s (tree is nil)", filePath)
	}

	// Cache the tree and content
	m.trees[filePath] = tree
	m.contents[filePath] = content
	return tree, nil
}

// GetTree retrieves a cached tree or parses the file
func (m *Manager) GetTree(filePath string) (*sitter.Tree, error) {
	return m.ParseFile(filePath)
}

// GetContent retrieves the cached content for a file
func (m *Manager) GetContent(filePath string) ([]byte, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	content, exists := m.contents[filePath]
	if !exists {
		return nil, fmt.Errorf("content not found for file: %s", filePath)
	}
	return content, nil
}

// Query executes a Tree-sitter query on a syntax tree
func (m *Manager) Query(tree *sitter.Tree, queryString string, langName string, content []byte) ([]*QueryMatch, error) {
	m.mu.RLock()
	lang, exists := m.languages[langName]
	m.mu.RUnlock()

	if !exists {
		return nil, fmt.Errorf("language %s not loaded", langName)
	}

	query, err := sitter.NewQuery([]byte(queryString), lang)
	if err != nil {
		return nil, fmt.Errorf("invalid query: %w", err)
	}
	defer query.Close()

	cursor := sitter.NewQueryCursor()
	defer cursor.Close()

	cursor.Exec(query, tree.RootNode())

	var matches []*QueryMatch
	for {
		match, ok := cursor.NextMatch()
		if !ok {
			break
		}

		for _, capture := range match.Captures {
			nodeContent := capture.Node.Content(content)

			matches = append(matches, &QueryMatch{
				Name:          query.CaptureNameForId(capture.Index),
				Text:          nodeContent,
				StartPosition: capture.Node.StartPoint(),
				EndPosition:   capture.Node.EndPoint(),
				StartByte:     capture.Node.StartByte(),
				EndByte:       capture.Node.EndByte(),
			})
		}
	}

	return matches, nil
}

// ListNodesByType finds all nodes of a specific type in a tree
func (m *Manager) ListNodesByType(tree *sitter.Tree, nodeType string, content []byte) []*NodeInfo {
	var nodes []*NodeInfo
	var traverse func(*sitter.Node)

	traverse = func(node *sitter.Node) {
		if node.Type() == nodeType {
			nameNode := node.ChildByFieldName("name")
			name := "(anonymous)"
			if nameNode != nil {
				name = nameNode.Content(content)
			}

			nodes = append(nodes, &NodeInfo{
				Name:          name,
				Type:          node.Type(),
				Text:          node.Content(content),
				StartPosition: node.StartPoint(),
				EndPosition:   node.EndPoint(),
			})
		}

		for i := 0; i < int(node.ChildCount()); i++ {
			traverse(node.Child(i))
		}
	}

	traverse(tree.RootNode())
	return nodes
}

// GetContextAtPosition finds the containing function/class at a position
func (m *Manager) GetContextAtPosition(tree *sitter.Tree, row, col uint32, content []byte) string {
	point := sitter.Point{Row: row, Column: col}
	node := tree.RootNode().NamedDescendantForPointRange(point, point)

	if node == nil {
		return ""
	}

	// Walk up the tree to find a block-level parent
	blockTypes := map[string]bool{
		"function_declaration": true,
		"method_definition":    true,
		"class_declaration":    true,
		"arrow_function":       true,
		"function_definition":  true, // Python
		"class_definition":     true, // Python
		"method_declaration":   true, // Go
		"type_declaration":     true, // Go
	}

	current := node
	for current != nil {
		if blockTypes[current.Type()] {
			return current.Content(content)
		}
		current = current.Parent()
	}

	return node.Content(content)
}

// ClearCache removes all cached parse trees
func (m *Manager) ClearCache() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.trees = make(map[string]*sitter.Tree)
	m.contents = make(map[string][]byte)
}

// InvalidateFile removes a specific file from the cache
func (m *Manager) InvalidateFile(filePath string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.trees, filePath)
	delete(m.contents, filePath)
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

// NodeInfo contains information about a syntax tree node
type NodeInfo struct {
	Name          string
	Type          string
	Text          string
	StartPosition sitter.Point
	EndPosition   sitter.Point
}
