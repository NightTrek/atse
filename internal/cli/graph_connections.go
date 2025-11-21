package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/NightTrek/atse/internal/index"
)

// ConnectionType represents the type of relationship between symbols
type ConnectionType string

const (
	ConnectionCalls    ConnectionType = "calls"    // Function/method invocations
	ConnectionTypes    ConnectionType = "types"    // Shared type usage
	ConnectionImports  ConnectionType = "imports"  // Module dependencies
	ConnectionDataFlow ConnectionType = "dataflow" // Data passed between functions
	ConnectionEvents   ConnectionType = "events"   // Event emitter/listener patterns
)

// ConnectionMetadata stores information about why symbols are connected
type ConnectionMetadata struct {
	Type        ConnectionType
	Description string // Human-readable description of the connection
	FilePath    string // Where the connection was discovered
	Line        uint32 // Line number of the connection
}

// Connection represents an edge in the graph
type Connection struct {
	From     *index.GraphNode
	To       *index.GraphNode
	Metadata ConnectionMetadata
}

// ConnectionFinder is an interface for different connection strategies
type ConnectionFinder interface {
	FindConnections(node *index.GraphNode, idx *index.SymbolIndex) []Connection
	Type() ConnectionType
}

// CallsConnectionFinder finds function call relationships
type CallsConnectionFinder struct{}

// NewCallsConnectionFinder creates a new calls connection finder
func NewCallsConnectionFinder() ConnectionFinder {
	return &CallsConnectionFinder{}
}

// FindConnections finds all functions that this node calls
func (cf *CallsConnectionFinder) FindConnections(node *index.GraphNode, idx *index.SymbolIndex) []Connection {
	var connections []Connection

	// Find all call sites in this node's file where this symbol is the caller
	for calledSymbol, callSites := range idx.Calls {
		for _, site := range callSites {
			// Check if this node is the caller
			if site.CallerSymbol == node.Symbol && site.CallerFilePath == node.FilePath {
				// Find the called symbol's definition
				calledNodes := idx.FindSymbol(calledSymbol)
				for _, calledNode := range calledNodes {
					connections = append(connections, Connection{
						From: node,
						To:   calledNode,
						Metadata: ConnectionMetadata{
							Type:        ConnectionCalls,
							Description: fmt.Sprintf("%s calls %s", node.Symbol, calledSymbol),
							FilePath:    site.CallerFilePath,
							Line:        site.CallerLine,
						},
					})
				}
			}
		}
	}

	if verboseFlag && len(connections) > 0 {
		fmt.Fprintf(os.Stderr, "    Found %d call connections from %s\n", len(connections), node.Symbol)
	}

	return connections
}

// Type returns the connection type
func (cf *CallsConnectionFinder) Type() ConnectionType {
	return ConnectionCalls
}

// TypesConnectionFinder finds shared type usage relationships
type TypesConnectionFinder struct{}

// NewTypesConnectionFinder creates a new types connection finder
func NewTypesConnectionFinder() ConnectionFinder {
	return &TypesConnectionFinder{}
}

// FindConnections finds symbols that share types with this node
func (tf *TypesConnectionFinder) FindConnections(node *index.GraphNode, idx *index.SymbolIndex) []Connection {
	var connections []Connection

	// Find types used in this node's file
	usedTypes := make(map[string]bool)
	for typeName, usages := range idx.Types {
		for _, usage := range usages {
			if usage.FilePath == node.FilePath {
				usedTypes[typeName] = true
			}
		}
	}

	// Find other symbols that use the same types
	for typeName := range usedTypes {
		usages := idx.FindTypeUsages(typeName)
		for _, usage := range usages {
			if usage.FilePath != node.FilePath {
				// Find symbols defined in this file
				targetSymbols := idx.FindSymbol(usage.FilePath)
				for _, targetNode := range targetSymbols {
					if targetNode.FilePath == usage.FilePath {
						connections = append(connections, Connection{
							From: node,
							To:   targetNode,
							Metadata: ConnectionMetadata{
								Type:        ConnectionTypes,
								Description: fmt.Sprintf("Both use type %s", typeName),
								FilePath:    usage.FilePath,
								Line:        usage.Line,
							},
						})
					}
				}
			}
		}
	}

	if verboseFlag && len(connections) > 0 {
		fmt.Fprintf(os.Stderr, "    Found %d type connections from %s\n", len(connections), node.Symbol)
	}

	return connections
}

// Type returns the connection type
func (tf *TypesConnectionFinder) Type() ConnectionType {
	return ConnectionTypes
}

// ImportsConnectionFinder finds import/module dependency relationships
type ImportsConnectionFinder struct{}

// NewImportsConnectionFinder creates a new imports connection finder
func NewImportsConnectionFinder() ConnectionFinder {
	return &ImportsConnectionFinder{}
}

// FindConnections finds symbols in imported modules
func (icf *ImportsConnectionFinder) FindConnections(node *index.GraphNode, idx *index.SymbolIndex) []Connection {
	var connections []Connection

	// Get imports for this node's file
	imports := idx.FindImportsFor(node.FilePath)

	// For each import, find symbols in the imported modules
	for _, imp := range imports {
		// Try to resolve import path to file paths
		importedFiles := resolveImportPath(imp.ImportPath, node.FilePath)

		for _, importedFile := range importedFiles {
			// Find all symbols in the imported file
			for _, symbols := range idx.Symbols {
				for _, symbol := range symbols {
					if symbol.FilePath == importedFile {
						connections = append(connections, Connection{
							From: node,
							To:   symbol,
							Metadata: ConnectionMetadata{
								Type:        ConnectionImports,
								Description: fmt.Sprintf("%s imports %s", filepath.Base(node.FilePath), imp.ImportPath),
								FilePath:    node.FilePath,
								Line:        imp.Line,
							},
						})
					}
				}
			}
		}
	}

	if verboseFlag && len(connections) > 0 {
		fmt.Fprintf(os.Stderr, "    Found %d import connections from %s\n", len(connections), node.Symbol)
	}

	return connections
}

// Type returns the connection type
func (icf *ImportsConnectionFinder) Type() ConnectionType {
	return ConnectionImports
}

// resolveImportPath attempts to resolve an import path to actual file paths
func resolveImportPath(importPath string, currentFile string) []string {
	var resolvedPaths []string

	// Simple resolution: for relative paths, resolve relative to current file
	if strings.HasPrefix(importPath, ".") {
		dir := filepath.Dir(currentFile)
		resolved := filepath.Join(dir, importPath)
		// Try common extensions
		for _, ext := range []string{".ts", ".js", ".go", ".py", ""} {
			candidate := resolved + ext
			resolvedPaths = append(resolvedPaths, candidate)
		}
	}

	// For absolute/package imports, would need more sophisticated resolution
	// For now, return empty - can be enhanced later

	return resolvedPaths
}

// DataFlowConnectionFinder finds data flow relationships (basic implementation)
type DataFlowConnectionFinder struct{}

// NewDataFlowConnectionFinder creates a new data flow connection finder
func NewDataFlowConnectionFinder() ConnectionFinder {
	return &DataFlowConnectionFinder{}
}

// FindConnections finds data flow connections (simplified - just returns calls for now)
func (df *DataFlowConnectionFinder) FindConnections(node *index.GraphNode, idx *index.SymbolIndex) []Connection {
	// For now, data flow is approximated by call relationships
	// A full implementation would track variables passed between functions
	// This is a placeholder that can be enhanced later
	callFinder := NewCallsConnectionFinder()
	connections := callFinder.FindConnections(node, idx)

	// Update metadata to indicate data flow
	for i := range connections {
		connections[i].Metadata.Type = ConnectionDataFlow
		connections[i].Metadata.Description = strings.Replace(
			connections[i].Metadata.Description,
			"calls",
			"passes data to",
			1,
		)
	}

	return connections
}

// Type returns the connection type
func (df *DataFlowConnectionFinder) Type() ConnectionType {
	return ConnectionDataFlow
}

// EventsConnectionFinder finds event emitter/listener relationships
type EventsConnectionFinder struct{}

// NewEventsConnectionFinder creates a new events connection finder
func NewEventsConnectionFinder() ConnectionFinder {
	return &EventsConnectionFinder{}
}

// FindConnections finds event-based connections
func (ef *EventsConnectionFinder) FindConnections(node *index.GraphNode, idx *index.SymbolIndex) []Connection {
	var connections []Connection

	// Find events emitted or listened to in this node's file
	eventsInFile := make(map[string][]index.EventUsage)
	for eventName, usages := range idx.Events {
		for _, usage := range usages {
			if usage.FilePath == node.FilePath {
				eventsInFile[eventName] = append(eventsInFile[eventName], usage)
			}
		}
	}

	// For each event, find other files that emit/listen to the same event
	for eventName := range eventsInFile {
		allUsages := idx.FindEventUsages(eventName)
		for _, usage := range allUsages {
			if usage.FilePath != node.FilePath {
				// Find symbols in the other file
				for _, symbols := range idx.Symbols {
					for _, symbol := range symbols {
						if symbol.FilePath == usage.FilePath {
							connections = append(connections, Connection{
								From: node,
								To:   symbol,
								Metadata: ConnectionMetadata{
									Type:        ConnectionEvents,
									Description: fmt.Sprintf("Connected via event '%s'", eventName),
									FilePath:    usage.FilePath,
									Line:        usage.Line,
								},
							})
						}
					}
				}
			}
		}
	}

	if verboseFlag && len(connections) > 0 {
		fmt.Fprintf(os.Stderr, "    Found %d event connections from %s\n", len(connections), node.Symbol)
	}

	return connections
}

// Type returns the connection type
func (ef *EventsConnectionFinder) Type() ConnectionType {
	return ConnectionEvents
}

// parseConnectionTypes parses a comma-separated string of connection types
func parseConnectionTypes(flagValue string) ([]ConnectionType, error) {
	if flagValue == "" {
		return []ConnectionType{ConnectionCalls}, nil
	}

	parts := strings.Split(flagValue, ",")
	var types []ConnectionType

	for _, part := range parts {
		part = strings.TrimSpace(strings.ToLower(part))
		switch part {
		case "calls":
			types = append(types, ConnectionCalls)
		case "types":
			types = append(types, ConnectionTypes)
		case "imports":
			types = append(types, ConnectionImports)
		case "dataflow", "data-flow":
			types = append(types, ConnectionDataFlow)
		case "events":
			types = append(types, ConnectionEvents)
		case "all":
			return []ConnectionType{
				ConnectionCalls,
				ConnectionTypes,
				ConnectionImports,
				ConnectionDataFlow,
				ConnectionEvents,
			}, nil
		default:
			return nil, fmt.Errorf("unknown connection type: %s", part)
		}
	}

	return types, nil
}

// createConnectionFinders creates ConnectionFinder instances for given types
func createConnectionFinders(types []ConnectionType) []ConnectionFinder {
	finders := make([]ConnectionFinder, 0, len(types))

	for _, connType := range types {
		switch connType {
		case ConnectionCalls:
			finders = append(finders, NewCallsConnectionFinder())
		case ConnectionTypes:
			finders = append(finders, NewTypesConnectionFinder())
		case ConnectionImports:
			finders = append(finders, NewImportsConnectionFinder())
		case ConnectionDataFlow:
			finders = append(finders, NewDataFlowConnectionFinder())
		case ConnectionEvents:
			finders = append(finders, NewEventsConnectionFinder())
		}
	}

	return finders
}
