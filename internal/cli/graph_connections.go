package cli

import (
	"fmt"
	"path/filepath"
	"strings"
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
	From     *GraphNode
	To       *GraphNode
	Metadata ConnectionMetadata
}

// ConnectionFinder is an interface for different connection strategies
type ConnectionFinder interface {
	FindConnections(node *GraphNode, idx *cliSymbolIndexAdapter) []Connection
	Type() ConnectionType
}

// CallsConnectionFinder finds function call relationships
type CallsConnectionFinder struct{}

// NewCallsConnectionFinder creates a new calls connection finder
func NewCallsConnectionFinder() ConnectionFinder {
	return &CallsConnectionFinder{}
}

// FindConnections finds all functions that this node calls
func (cf *CallsConnectionFinder) FindConnections(node *GraphNode, idx *cliSymbolIndexAdapter) []Connection {
	var connections []Connection

	// Find all call sites in this node's file where this symbol is the caller
	// NOTE: In the hybrid architecture, we don't have call sites pre-indexed.
	// We only have symbols.
	//
	// To make this work, we need to:
	// 1. Parse the file (already done if in partial index)
	// 2. Run a query to find calls within the function body
	// 3. Resolve those calls to symbols

	// For this initial hybrid implementation, we will implement a simplified version
	// that doesn't find calls yet (because `extractSymbolsForFile` only extracts declarations).

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
func (tf *TypesConnectionFinder) FindConnections(node *GraphNode, idx *cliSymbolIndexAdapter) []Connection {
	var connections []Connection
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
func (icf *ImportsConnectionFinder) FindConnections(node *GraphNode, idx *cliSymbolIndexAdapter) []Connection {
	var connections []Connection
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
	return resolvedPaths
}

// DataFlowConnectionFinder finds data flow relationships (basic implementation)
type DataFlowConnectionFinder struct{}

// NewDataFlowConnectionFinder creates a new data flow connection finder
func NewDataFlowConnectionFinder() ConnectionFinder {
	return &DataFlowConnectionFinder{}
}

// FindConnections finds data flow connections (simplified - just returns calls for now)
func (df *DataFlowConnectionFinder) FindConnections(node *GraphNode, idx *cliSymbolIndexAdapter) []Connection {
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
func (ef *EventsConnectionFinder) FindConnections(node *GraphNode, idx *cliSymbolIndexAdapter) []Connection {
	var connections []Connection
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
