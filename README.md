# ATSE - Agent Tree Search CLI

> A Tree-sitter powered command-line tool for fast, accurate code analysis at scale.

**atse** helps developers and AI agents understand massive codebases instantly through structural, syntax-aware code analysis. 

It uses a **hybrid architecture** combining ripgrep (for instant file discovery) and Tree-sitter (for structural parsing) to handle repositories with 100,000+ files with zero latency.

## 🚀 Core Features

- **Hybrid Search**: Instant structural search, powered by ripgrep + Tree-sitter
- **Lazy Graph**: Dependency analysis that scales to millions of lines of code
- **Smart Extract**: Context-aware code extraction for AI agents
- **Zero Config**: No database, no server setup required

## 📦 Installation

```bash
# Install via shell script
curl -fsSL https://raw.githubusercontent.com/NightTrek/atse/main/scripts/install.sh | bash

# Or build from source (requires Go 1.23+)
go install github.com/NightTrek/atse/cmd/atse@latest
```

**Prerequisites**: 
- `rg` (ripgrep) must be installed and in your PATH
- C compiler (gcc/clang) for building Tree-sitter bindings

## ⚡️ Usage

### 1. Search (Entry Points)
Find symbols instantly without false positives.

```bash
atse search AuthService ./src
atse search "login" ./src --fuzzy
atse search User --type class
```

### 2. Graph (Understand Relationships)
Trace dependencies lazily. Only parses what is connected.

```bash
atse graph AuthService ./src
atse graph validateUser --depth 3
atse graph User --mode dfs
```

### 3. Extract (Get Context)
Get full source code for a feature and its dependencies.

```bash
atse extract AuthService ./src
atse extract handleRequest --depth 2 --output context.txt
```

### 4. Find References
Find function calls with zero false positives.

```bash
atse find-fn authenticate
```

## 🔄 Migration Guide (v0.2.0)

We've streamlined ATSE to focus on what it does best. Some commands have been removed in favor of using `ripgrep` directly, which is faster for simple tasks.

| Old Command | New Best Practice (using `rg`) |
|-------------|--------------------------------|
| `atse list-fns` | `rg "^func " --type go` |
| `atse deps` | `rg "^import" --type go` |
| `atse query` | Use `ast-grep` or specialized tools |

##  Architecture

ATSE uses a **lazy-loading partial index**:
1. **Discovery**: Uses `ripgrep` to find candidate files (~50ms)
2. **Parsing**: Parses only relevants files with Tree-sitter (~100ms)
3. **Expansion**: Follows imports/calls to discover new files on-demand
4. **Result**: Instant structural analysis without indexing the whole repo

This allows ATSE to work on 100k+ file monorepos where traditional LSP servers might struggle or burn gigabytes of RAM.

## License

MIT
