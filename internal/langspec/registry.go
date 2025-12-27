package langspec

import (
    "fmt"
    "strings"

    sitter "github.com/smacker/go-tree-sitter"
    "github.com/smacker/go-tree-sitter/bash"
    "github.com/smacker/go-tree-sitter/c"
    "github.com/smacker/go-tree-sitter/csharp"
    "github.com/smacker/go-tree-sitter/cpp"
    "github.com/smacker/go-tree-sitter/golang"
    "github.com/smacker/go-tree-sitter/java"
    "github.com/smacker/go-tree-sitter/javascript"
    "github.com/smacker/go-tree-sitter/php"
    "github.com/smacker/go-tree-sitter/python"
    "github.com/smacker/go-tree-sitter/ruby"
    "github.com/smacker/go-tree-sitter/rust"
    "github.com/smacker/go-tree-sitter/typescript/typescript"
)

// LanguageConfig defines how a language is loaded and queried.
type LanguageConfig struct {
    Name               string
    Aliases            []string
    Extensions         []string
    Load               func() *sitter.Language
    SymbolQueries      map[string]string
    TypeUsageQueries   []string
    ImportQueries      []string
    CallCaptureQueries []string
    CallQueryTemplate  string
    EventQueries       []string
}

// Registry manages language configurations and extension resolution.
type Registry struct {
    configs        map[string]*LanguageConfig
    extensionIndex map[string]string
    aliasIndex     map[string]string
}

var defaultRegistry *Registry

// DefaultRegistry returns the singleton registry populated with common Tree-sitter languages.
func DefaultRegistry() *Registry {
    if defaultRegistry != nil {
        return defaultRegistry
    }

    r := &Registry{
        configs:        make(map[string]*LanguageConfig),
        extensionIndex: make(map[string]string),
        aliasIndex:     make(map[string]string),
    }

    for _, cfg := range defaultLanguages() {
        r.Register(cfg)
    }

    defaultRegistry = r
    return defaultRegistry
}

// Register adds a language configuration to the registry.
func (r *Registry) Register(cfg *LanguageConfig) {
    if cfg == nil || cfg.Name == "" {
        return
    }

    key := strings.ToLower(cfg.Name)
    r.configs[key] = cfg

    for _, alias := range cfg.Aliases {
        r.aliasIndex[strings.ToLower(alias)] = key
    }

    for _, ext := range cfg.Extensions {
        r.extensionIndex[strings.ToLower(ext)] = key
    }
}

// Resolve returns a language configuration by name or alias.
func (r *Registry) Resolve(name string) (*LanguageConfig, bool) {
    if name == "" {
        return nil, false
    }

    key := strings.ToLower(name)
    if cfg, ok := r.configs[key]; ok {
        return cfg, true
    }

    if mapped, ok := r.aliasIndex[key]; ok {
        cfg, ok := r.configs[mapped]
        return cfg, ok
    }

    return nil, false
}

// ResolveByExtension returns a language configuration based on file extension (including the dot).
func (r *Registry) ResolveByExtension(ext string) (*LanguageConfig, bool) {
    if ext == "" {
        return nil, false
    }
    key := strings.ToLower(ext)
    if name, ok := r.extensionIndex[key]; ok {
        return r.configs[name], true
    }
    return nil, false
}

// SymbolQueries returns symbol queries for the given language, or nil if not found.
func (r *Registry) SymbolQueries(name string) map[string]string {
    if cfg, ok := r.Resolve(name); ok {
        return cfg.SymbolQueries
    }
    return nil
}

// CallCaptureQueries returns call extraction queries for the given language.
func (r *Registry) CallCaptureQueries(name string) []string {
    if cfg, ok := r.Resolve(name); ok {
        return cfg.CallCaptureQueries
    }
    return nil
}

// TypeUsageQueries returns type-usage queries for the language.
func (r *Registry) TypeUsageQueries(name string) []string {
    if cfg, ok := r.Resolve(name); ok {
        return cfg.TypeUsageQueries
    }
    return nil
}

// ImportQueries returns import queries for the language.
func (r *Registry) ImportQueries(name string) []string {
    if cfg, ok := r.Resolve(name); ok {
        return cfg.ImportQueries
    }
    return nil
}

// EventQueries returns event pattern queries for the language.
func (r *Registry) EventQueries(name string) []string {
    if cfg, ok := r.Resolve(name); ok {
        return cfg.EventQueries
    }
    return nil
}

// BuildFunctionCallQuery renders a call query for a specific function name, if supported.
func (r *Registry) BuildFunctionCallQuery(langName, funcName string) string {
    cfg, ok := r.Resolve(langName)
    if !ok || cfg.CallQueryTemplate == "" {
        return ""
    }
    return fmt.Sprintf(cfg.CallQueryTemplate, funcName)
}

func defaultLanguages() []*LanguageConfig {
    return []*LanguageConfig{
        {
            Name:       "go",
            Extensions: []string{".go"},
            Load:       golang.GetLanguage,
            SymbolQueries: map[string]string{
                "function": `(function_declaration name: (identifier) @name)`,
                "method":   `(method_declaration name: (field_identifier) @name)`,
                "type":     `(type_spec name: (type_identifier) @name)`,
            },
            TypeUsageQueries: []string{
                `(type_identifier) @type`,
                `(qualified_type package: (package_identifier) name: (type_identifier) @type)`,
            },
            ImportQueries: []string{
                `(import_spec path: (interpreted_string_literal) @path)`,
            },
            CallCaptureQueries: []string{
                `(call_expression function: (identifier) @name)`,
            },
            CallQueryTemplate: `(call_expression function: (identifier) @fn (#eq? @fn "%s"))`,
        },
        {
            Name:       "javascript",
            Aliases:    []string{"js"},
            Extensions: []string{".js", ".jsx", ".mjs", ".cjs"},
            Load:       javascript.GetLanguage,
            SymbolQueries: map[string]string{
                "function": `(function_declaration name: (identifier) @name)` ,
                "class":    `(class_declaration name: (type_identifier) @name)`,
                "method":   `(method_definition name: (property_identifier) @name)`,
                "const":    `(lexical_declaration (variable_declarator name: (identifier) @name)) @const`,
            },
            TypeUsageQueries: []string{},
            ImportQueries: []string{
                `(import_statement source: (string) @path)`,
            },
            CallCaptureQueries: []string{
                `(call_expression function: (identifier) @name)`,
            },
            CallQueryTemplate: `(call_expression function: (identifier) @fn (#eq? @fn "%s"))`,
            EventQueries: []string{
                `(call_expression
                    function: (member_expression
                        property: (property_identifier) @method)
                    arguments: (arguments (string) @event))`,
            },
        },
        {
            Name:       "typescript",
            Aliases:    []string{"ts", "tsx"},
            Extensions: []string{".ts", ".tsx"},
            Load:       typescript.GetLanguage,
            SymbolQueries: map[string]string{
                "function": `(function_declaration name: (identifier) @name)` ,
                "class":    `(class_declaration name: (type_identifier) @name)`,
                "method":   `(method_definition name: (property_identifier) @name)`,
                "const":    `(lexical_declaration (variable_declarator name: (identifier) @name)) @const`,
            },
            TypeUsageQueries: []string{
                `(type_annotation (type_identifier) @type)`,
                `(type_annotation (generic_type name: (type_identifier) @type))`,
            },
            ImportQueries: []string{
                `(import_statement source: (string) @path)`,
            },
            CallCaptureQueries: []string{
                `(call_expression function: (identifier) @name)`,
            },
            CallQueryTemplate: `(call_expression function: (identifier) @fn (#eq? @fn "%s"))`,
            EventQueries: []string{
                `(call_expression
                    function: (member_expression
                        property: (property_identifier) @method)
                    arguments: (arguments (string) @event))`,
            },
        },
        {
            Name:       "python",
            Aliases:    []string{"py"},
            Extensions: []string{".py"},
            Load:       python.GetLanguage,
            SymbolQueries: map[string]string{
                "function": `(function_definition name: (identifier) @name)`,
                "class":    `(class_definition name: (identifier) @name)`,
            },
            TypeUsageQueries: []string{
                `(type (identifier) @type)`,
            },
            ImportQueries: []string{
                `(import_statement name: (dotted_name) @path)`,
                `(import_from_statement module: (dotted_name) @path)`,
            },
            CallCaptureQueries: []string{
                `(call function: (identifier) @name)`,
            },
            CallQueryTemplate: `(call function: (identifier) @fn (#eq? @fn "%s"))`,
        },
        {
            Name:       "java",
            Extensions: []string{".java"},
            Load:       java.GetLanguage,
            SymbolQueries: map[string]string{
                "class":       `(class_declaration name: (identifier) @name)`,
                "interface":   `(interface_declaration name: (identifier) @name)`,
                "method":      `(method_declaration name: (identifier) @name)`,
                "constructor": `(constructor_declaration name: (identifier) @name)`,
            },
            TypeUsageQueries: []string{
                `(type_identifier) @type`,
            },
            ImportQueries: []string{
                `(import_declaration (scoped_identifier) @path)`,
            },
            CallCaptureQueries: []string{
                `(method_invocation name: (identifier) @name)`,
            },
            CallQueryTemplate: `(method_invocation name: (identifier) @fn (#eq? @fn "%s"))`,
        },
        {
            Name:       "rust",
            Extensions: []string{".rs"},
            Load:       rust.GetLanguage,
            SymbolQueries: map[string]string{
                "function": `(function_item name: (identifier) @name)`,
                "method":   `(impl_item (function_item name: (identifier) @name))`,
                "trait":    `(trait_item name: (identifier) @name)`,
            },
            TypeUsageQueries: []string{
                `(type_identifier) @type`,
            },
            ImportQueries: []string{
                `(use_declaration (scoped_identifier) @path)`,
                `(use_declaration (use_wildcard (scoped_identifier) @path))`,
            },
            CallCaptureQueries: []string{
                `(call_expression function: (identifier) @name)`,
                `(call_expression function: (scoped_identifier) @name)`,
            },
            CallQueryTemplate: `(call_expression function: (identifier) @fn (#eq? @fn "%s"))`,
        },
        {
            Name:       "c",
            Extensions: []string{".c", ".h"},
            Load:       c.GetLanguage,
            SymbolQueries: map[string]string{
                "function": `(function_definition declarator: (function_declarator declarator: (identifier) @name))`,
            },
            TypeUsageQueries: []string{},
            ImportQueries: []string{
                `(preproc_include path: (string_literal) @path)`,
            },
            CallCaptureQueries: []string{
                `(call_expression function: (identifier) @name)`,
            },
            CallQueryTemplate: `(call_expression function: (identifier) @fn (#eq? @fn "%s"))`,
        },
        {
            Name:       "cpp",
            Aliases:    []string{"c++", "cplusplus"},
            Extensions: []string{".cc", ".cpp", ".cxx", ".hpp", ".hh", ".hxx"},
            Load:       cpp.GetLanguage,
            SymbolQueries: map[string]string{
                "function": `(function_definition declarator: (function_declarator declarator: (identifier) @name))`,
                "class":    `(class_specifier name: (type_identifier) @name)`,
                "struct":   `(struct_specifier name: (type_identifier) @name)`,
            },
            TypeUsageQueries: []string{},
            ImportQueries: []string{
                `(preproc_include path: (string_literal) @path)`,
            },
            CallCaptureQueries: []string{
                `(call_expression function: (identifier) @name)`,
            },
            CallQueryTemplate: `(call_expression function: (identifier) @fn (#eq? @fn "%s"))`,
        },
        {
            Name:       "csharp",
            Aliases:    []string{"c#", "cs", "csharp"},
            Extensions: []string{".cs"},
            Load:       csharp.GetLanguage,
            SymbolQueries: map[string]string{
                "class":  `(class_declaration name: (identifier) @name)`,
                "method": `(method_declaration name: (identifier) @name)`,
                "interface": `(interface_declaration name: (identifier) @name)`,
            },
            TypeUsageQueries: []string{
                `(identifier) @type`,
            },
            ImportQueries: []string{
                `(using_directive (qualified_name) @path)`,
            },
            CallCaptureQueries: []string{
                `(invocation_expression function: (identifier) @name)`,
            },
            CallQueryTemplate: `(invocation_expression function: (identifier) @fn (#eq? @fn "%s"))`,
        },
        {
            Name:       "php",
            Extensions: []string{".php"},
            Load:       php.GetLanguage,
            SymbolQueries: map[string]string{
                "function": `(function_definition name: (name) @name)`,
                "method":   `(method_declaration name: (name) @name)`,
                "class":    `(class_declaration name: (name) @name)`,
            },
            TypeUsageQueries: []string{},
            ImportQueries: []string{
                `(use_declaration (qualified_name) @path)`,
            },
            CallCaptureQueries: []string{
                `(function_call_expression function: (name) @name)`,
            },
            CallQueryTemplate: `(function_call_expression function: (name) @fn (#eq? @fn "%s"))`,
        },
        {
            Name:       "ruby",
            Extensions: []string{".rb"},
            Load:       ruby.GetLanguage,
            SymbolQueries: map[string]string{
                "method": `(method name: (identifier) @name)`,
                "class":  `(class name: (constant) @name)`,
                "module": `(module name: (constant) @name)`,
            },
            TypeUsageQueries: []string{},
            ImportQueries: []string{
                `(call receiver: (identifier) @prefix method: (identifier) @method
                    arguments: (argument_list (string) @path))`,
            },
            CallCaptureQueries: []string{
                `(call method: (identifier) @name)`,
            },
            CallQueryTemplate: `(call method: (identifier) @fn (#eq? @fn "%s"))`,
        },
        {
            Name:       "bash",
            Extensions: []string{".sh"},
            Load:       bash.GetLanguage,
            SymbolQueries: map[string]string{
                "function": `(function_definition name: (word) @name)`,
            },
            TypeUsageQueries: []string{},
            ImportQueries: []string{},
            CallCaptureQueries: []string{
                `(command name: (word) @name)`,
            },
            CallQueryTemplate: `(command name: (word) @fn (#eq? @fn "%s"))`,
        },
    }
}
