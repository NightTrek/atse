# ATSE - Agent Tree Search CLI

> A Tree-sitter powered command-line tool for fast, accurate code analysis

**atse** (pronounced "at-see" or "at-ess-ee") helps developers and AI agents understand large codebases through structural, syntax-aware code analysis.

## 🎯 Why ATSE?

Traditional text-based search tools like `grep` produce many false positives and don't understand code structure. ATSE uses [Tree-sitter](https://tree-sitter.github.io/tree-sitter/) to parse code into syntax trees, enabling:

- ✅ **Zero false positives** - searches understand code structure
- ✅ **Multi-language support** - TypeScript, JavaScript, Python, Go, Java
- ✅ **Fast analysis** - parallel processing with Go
- ✅ **Token-efficient output** - optimized for AI agents
- ✅ **Composable** - Unix-style CLI that works with pipes and scripts

## 🚀 Quick Start

### Installation

```bash
# Clone the repository
git clone https://github.com/NightTrek/atse.git
cd atse

# Build using Make (recommended)
make build

# Or build manually with correct CGO flags
CGO_ENABLED=1 go build -ldflags="-s -w" -o atse ./cmd/atse/main.go

# Verify installation
./atse --version
./atse list-fns tests/fixtures/simple.go
```

**Note:** ATSE requires CGO to be enabled for Tree-sitter bindings. The Makefile handles this automatically.

#### Make Commands

```bash
make build    # Build the binary
make clean    # Clean build artifacts
make check    # Build and run quick tests
make install  # Install to $GOPATH/bin
```

### Basic Usage

```bash
# Find all calls to a function
atse find-fn authenticate ./src

# List all functions in a file
atse list-fns ./src/api.ts

# Show dependencies (imports)
atse deps ./src

# Get context around a specific location
atse context ./src/api.ts:42:10

# Execute a raw Tree-sitter query
atse query "(function_declaration)" ./src
```

## 📚 Commands

### `find-fn` - Find Function Calls
Locate all places where a specific function is called.

```bash
atse find-fn authenticate ./src
atse find-fn login ./src --format json
atse find-fn processUser . --include "*.ts" --exclude "*.test.ts"
```

### `list-fns` - List Functions
Enumerate all function declarations in files.

```bash
atse list-fns ./src/api.ts
atse list-fns ./src --format json
atse list-fns . --include "*.go" --verbose
```

### `deps` - Show Dependencies
Analyze import statements and dependencies.

```bash
atse deps ./src/api.ts
atse deps ./src --format json
atse deps . --include "*.go"
```

### `context` - Get Code Context
Retrieve the surrounding context (function, class) for a position.

```bash
atse context ./src/api.ts:42:10
atse context ./src/api.ts:42:10 --verbose
atse context ./src/api.ts:42:10 --format json
```

### `query` - Raw Tree-sitter Query
Execute custom Tree-sitter queries (power users).

```bash
atse query "(function_declaration)" ./src
atse query "(call_expression function: (identifier) @fn)" ./src --format json
atse query "(class_declaration name: (identifier) @name)" .
```

## 🎛️ Global Flags

All commands support these flags:

- `--format <text|json|locations>` - Output format (default: text)
- `--verbose, -v` - Include additional context
- `--recursive, -r` - Recursively search directories (default: true)
- `--include <pattern>` - Include files matching pattern (e.g., `"*.ts"`)
- `--exclude <pattern>` - Exclude files matching pattern (e.g., `"*.test.ts"`)
- `--limit <n>` - Limit number of results (0 = no limit)
- `--no-cache` - Bypass parse tree cache

## 🏗️ Architecture

```
atse/
├── cmd/atse/           # CLI entry point
├── internal/
│   ├── cli/            # Cobra command definitions
│   ├── parser/         # Tree-sitter parser manager
│   ├── search/         # Search and query logic
│   ├── analysis/       # Dependency and reference analysis
│   ├── output/         # Output formatting
│   ├── config/         # Configuration management
│   └── util/           # Utilities (paths, workers)
└── tests/              # Test fixtures and integration tests
```

## 🔧 Development Status

**Current Phase:** Phase 2 Complete - All Core Commands Implemented ✅

- [x] Project structure and Go module setup
- [x] Cobra CLI framework with all core commands
- [x] Global flags and help documentation
- [x] Build system with proper CGO configuration
- [x] Parser manager implementation (fully functional)
- [x] All command implementations (`list-fns`, `find-fn`, `deps`, `context`, `query`)
- [x] Output formatting (text, JSON, locations)
- [x] Multi-language support (TypeScript, JavaScript, Python, Go)
- [x] Parse tree caching
- [ ] Comprehensive test suite
- [ ] Configuration file support

## 🎯 Design Philosophy

Based on research from [Mario Zechner's "MCP vs CLI" analysis](https://mariozechner.at/posts/2025-08-15-mcp-vs-cli/):

1. **CLI First** - Simple, portable, composable
2. **Token Efficient** - Clean output, no unnecessary JSON
3. **Single Command** - One binary with subcommands
4. **Progressive Disclosure** - Minimal by default, verbose when needed
5. **Composable** - Works with Unix pipes and tools

## 🤖 AI Agent Integration

ATSE is designed to work seamlessly with AI coding agents:

```typescript
// Example: Using atse from an AI agent
const result = await execute_command({
    command: `atse find-fn authenticate ./src --format json`,
    requires_approval: false
});

const matches = JSON.parse(result.stdout);
```

Benefits for agents:
- No MCP server setup required
- Fast execution (native binary)
- Token-efficient output
- Works in any environment

## 📖 Resources

- **Tree-sitter MCP Server** - Original TypeScript implementation that inspired this project
- **Blog Post** - [MCP vs CLI: Benchmarking Tools for Coding Agents](https://mariozechner.at/posts/2025-08-15-mcp-vs-cli/)
- **Tree-sitter** - [Official Documentation](https://tree-sitter.github.io/tree-sitter/)

## 🗺️ Roadmap

### Phase 1: Core Implementation ✅ COMPLETE
- [x] Parser manager with language detection
- [x] Implement `find-fn`, `list-fns`, `deps`, `context`, `query`
- [x] Basic caching
- [ ] Comprehensive unit and integration tests

### Phase 2: Advanced Features
- [ ] Reference and definition resolution
- [ ] Project-wide dependency graphs
- [ ] Configuration file support (`.atse.toml`)
- [ ] Performance optimizations

### Phase 3: Distribution
- [ ] Release binaries (macOS, Linux, Windows)
- [ ] Homebrew formula
- [ ] Comprehensive documentation
- [ ] Optional MCP wrapper

## 📄 License

MIT License - see LICENSE file for details

## 🤝 Contributing

Contributions welcome! This project is in active development.

---

**Author:** Daniel Steigman ([@NightTrek](https://github.com/NightTrek))

**Status:** 🚧 Under Active Development
