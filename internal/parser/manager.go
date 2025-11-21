package parser

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/NightTrek/atse/internal/index"
	"github.com/NightTrek/atse/internal/util"
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
	contents  map[string][]byte  // Store file contents for node extraction
	Index     *index.SymbolIndex // Shared symbol index
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

// GetOrLoadIndex retrieves the persistent index or builds it if missing/stale
func (m *Manager) GetOrLoadIndex(rootPath string, files []util.FileMatch, verbose bool) (*index.SymbolIndex, error) {
	m.mu.RLock()
	if m.Index != nil {
		defer m.mu.RUnlock()
		return m.Index, nil
	}
	m.mu.RUnlock()

	// Determine cache path
	absRoot, err := filepath.Abs(rootPath)
	if err != nil {
		return nil, err
	}
	// Try project root .atse/index.cache
	cacheDir := filepath.Join(absRoot, ".atse")
	cachePath := filepath.Join(cacheDir, "index.cache")

	// Try to load from disk
	idx, err := index.Load(cachePath)
	if err == nil && idx != nil {
		// Validate cache (check mtimes)
		if !idx.IsStale(files) {
			if verbose {
				fmt.Fprintf(os.Stderr, "Loaded index from cache: %s\n", cachePath)
			}
			m.mu.Lock()
			m.Index = idx
			m.mu.Unlock()
			return idx, nil
		}
		if verbose {
			fmt.Fprintf(os.Stderr, "Index cache is stale, rebuilding...\n")
		}
	} else if verbose {
		fmt.Fprintf(os.Stderr, "No index cache found at %s, building...\n", cachePath)
	}

	// Build index
	idx, err = index.Build(m, files, verbose)
	if err != nil {
		return nil, err
	}

	// Save index
	if err := idx.Save(cachePath); err != nil {
		if verbose {
			fmt.Fprintf(os.Stderr, "Warning: failed to save index cache: %v\n", err)
		}
	} else if verbose {
		fmt.Fprintf(os.Stderr, "Saved index to %s\n", cachePath)
	}

	m.mu.Lock()
	m.Index = idx
	m.mu.Unlock()
	return idx, nil
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
func (m *Manager) Query(tree *sitter.Tree, queryString string, langName string, content []byte) ([]*index.QueryMatch, error) {
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

	var matches []*index.QueryMatch
	for {
		match, ok := cursor.NextMatch()
		if !ok {
			break
		}

		for _, capture := range match.Captures {
			nodeContent := capture.Node.Content(content)

			matches = append(matches, &index.QueryMatch{
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
// Note: NodeInfo is not in index package yet, but used by list-fns.
// I should probably move NodeInfo to index package too if needed by Parser interface.
// But Parser interface only needs Extract* methods and Query.
// ListNodesByType is used by list-fns command directly.
// I will define NodeInfo here for now (if not moved).
type NodeInfo struct {
	Name          string
	Type          string
	Text          string
	StartPosition sitter.Point
	EndPosition   sitter.Point
}

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

// ExtractTypeUsages finds all type references in a file
func (m *Manager) ExtractTypeUsages(tree *sitter.Tree, langName string, content []byte, filePath string) []*index.TypeUsage {
	var usages []*index.TypeUsage

	// Build queries for type usages based on language
	var queries []string
	switch langName {
	case "typescript", "javascript":
		queries = []string{
			`(type_annotation (type_identifier) @type)`,
			`(type_annotation (generic_type name: (type_identifier) @type))`,
		}
	case "go":
		queries = []string{
			`(type_identifier) @type`,
			`(qualified_type package: (package_identifier) name: (type_identifier) @type)`,
		}
	case "python":
		queries = []string{
			`(type (identifier) @type)`,
		}
	}

	for _, queryString := range queries {
		results, err := m.Query(tree, queryString, langName, content)
		if err != nil {
			continue
		}

		for _, result := range results {
			typeName := strings.TrimSpace(result.Text)
			if typeName != "" {
				usages = append(usages, &index.TypeUsage{
					TypeName:  typeName,
					UsageType: "reference",
					FilePath:  filePath,
					Line:      result.StartPosition.Row,
				})
			}
		}
	}

	return usages
}

// ExtractImports parses import/require/include statements
func (m *Manager) ExtractImports(tree *sitter.Tree, langName string, content []byte, filePath string) []index.Import {
	var imports []index.Import

	// Build queries for imports based on language
	var queryString string
	switch langName {
	case "typescript", "javascript":
		queryString = `(import_statement source: (string) @path)`
	case "go":
		queryString = `(import_spec path: (interpreted_string_literal) @path)`
	case "python":
		queryString = `(import_statement name: (dotted_name) @path)`
	}

	if queryString == "" {
		return imports
	}

	results, err := m.Query(tree, queryString, langName, content)
	if err != nil {
		return imports
	}

	for _, result := range results {
		importPath := strings.Trim(strings.TrimSpace(result.Text), "\"'`")
		if importPath != "" {
			imports = append(imports, index.Import{
				ImportPath: importPath,
				FilePath:   filePath,
				Line:       result.StartPosition.Row,
			})
		}
	}

	return imports
}

// ExtractCallSites finds all function call expressions
func (m *Manager) ExtractCallSites(tree *sitter.Tree, langName string, content []byte, filePath string) []index.CallSite {
	var callSites []index.CallSite

	// Build queries for function calls based on language
	var queryString string
	switch langName {
	case "typescript", "javascript":
		queryString = `(call_expression function: (identifier) @name)`
	case "go":
		queryString = `(call_expression function: (identifier) @name)`
	case "python":
		queryString = `(call function: (identifier) @name)`
	}

	if queryString == "" {
		return callSites
	}

	results, err := m.Query(tree, queryString, langName, content)
	if err != nil {
		return callSites
	}

	// For each call, try to find the containing function
	for _, result := range results {
		calledName := strings.TrimSpace(result.Text)
		if calledName == "" {
			continue
		}

		// Find containing function
		caller := m.findContainingFunction(tree, result.StartPosition.Row, content)

		callSites = append(callSites, index.CallSite{
			CallerSymbol:   caller,
			CallerFilePath: filePath,
			CallerLine:     result.StartPosition.Row,
			CalledSymbol:   calledName,
		})
	}

	return callSites
}

// findContainingFunction finds the function that contains a given line
func (m *Manager) findContainingFunction(tree *sitter.Tree, line uint32, content []byte) string {
	point := sitter.Point{Row: line, Column: 0}
	node := tree.RootNode().NamedDescendantForPointRange(point, point)

	if node == nil {
		return "(global)"
	}

	// Walk up the tree to find a function
	blockTypes := map[string]bool{
		"function_declaration": true,
		"method_definition":    true,
		"method_declaration":   true,
		"function_definition":  true,
	}

	current := node
	for current != nil {
		if blockTypes[current.Type()] {
			// Try to get function name
			nameNode := current.ChildByFieldName("name")
			if nameNode != nil {
				return nameNode.Content(content)
			}
		}
		current = current.Parent()
	}

	return "(anonymous)"
}

// ExtractEventPatterns finds event emitter/listener patterns
func (m *Manager) ExtractEventPatterns(tree *sitter.Tree, langName string, content []byte, filePath string) []index.EventUsage {
	var events []index.EventUsage

	// For now, use a simple pattern: look for .on(), .emit(), addEventListener calls
	// This is language-agnostic at the call site level
	var queryString string
	switch langName {
	case "typescript", "javascript":
		queryString = `(call_expression
			function: (member_expression
				property: (property_identifier) @method)
			arguments: (arguments (string) @event))`
	default:
		return events
	}

	results, err := m.Query(tree, queryString, langName, content)
	if err != nil {
		return events
	}

	// Process results in pairs (method, event)
	for i := 0; i+1 < len(results); i += 2 {
		methodName := strings.TrimSpace(results[i].Text)
		eventName := strings.Trim(strings.TrimSpace(results[i+1].Text), "\"'`")

		// Check if it's an event-related method
		eventType := ""
		switch methodName {
		case "on", "addEventListener", "addListener", "once":
			eventType = "listen"
		case "emit", "dispatchEvent", "fire", "trigger":
			eventType = "emit"
		default:
			continue
		}

		if eventName != "" && eventType != "" {
			events = append(events, index.EventUsage{
				EventName: eventName,
				Type:      eventType,
				FilePath:  filePath,
				Line:      results[i].StartPosition.Row,
			})
		}
	}

	return events
}
