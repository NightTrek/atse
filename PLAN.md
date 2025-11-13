# ATSE (Agent Tree Search CLI) – Project Blueprint

## Vision
Build a Tree-sitter powered command-line interface that helps developers (humans and AI agents) understand large codebases faster and more efficiently. The tool should balance powerful structural analysis with a simple CLI experience, provide token-efficient output, and serve as a foundation for optional MCP integration later.

## Goals
- Deliver high-level and targeted code analysis commands that work across entire projects.
- Provide clean, composable CLI output formats optimized for both humans and agents.
- Support raw Tree-sitter queries for power users while exposing ergonomic high-level commands.
- Design a modular architecture that can evolve to include caching, configuration, and an MCP wrapper.
- Ensure high code quality through comprehensive tests for each command.

## Resources & Inspiration
- **Tree-sitter MCP Server (current project)**
  - `src/manager.ts`: Parser lifecycle, language loading, search, node listing, context extraction.
  - `src/tools/handlers.ts`: Argument validation, tree retrieval, response formatting.
  - `src/middleware.ts`: Token counting pattern.
- **Blog Post:** [Mario Zechner – "MCP vs CLI: Benchmarking Tools for Coding Agents" (2025-08-15)](https://mariozechner.at/posts/2025-08-15-mcp-vs-cli/)
  - Key insight: the protocol matters less than tool design; CLI + thin MCP wrapper is a winning combo.
  - Emphasis on token-efficient output, minimal round trips, and a single well-defined tool surface.
- **Go Ecosystem Documentation**
  - [`github.com/spf13/cobra`](https://github.com/spf13/cobra): CLI framework.
  - [`github.com/smacker/go-tree-sitter`](https://github.com/smacker/go-tree-sitter): Go bindings for Tree-sitter.
  - Tree-sitter language grammars for TypeScript, JavaScript, Python, Go, Java (matching MCP server coverage).
- **Existing Workflow Patterns & Ideas**
  - CLI composability (`grep`, `sort`, `jq`) for post-processing results.
  - Unix philosophy: single binary with subcommands.
  - Token-efficient defaults inspired by terminalcp’s evaluation.

## Target Users & Use Cases
- **Software engineers** onboarding onto large repos.
- **AI agents** needing fast structural insights without noisy output.
- **Power users** requiring raw Tree-sitter query capabilities.
- **Automation scripts** that consume structured JSON when needed.

## Top-Level Command Design
```
atse
├── find-fn        # Locate function call sites across files/directories
├── find-class     # Locate class usages/instantiations
├── find-var       # Locate variable references
├── list-fns       # Enumerate function declarations
├── list-classes   # Enumerate class declarations
├── list-imports   # Enumerate imports/exports
├── deps           # Report dependencies/import graph for file(s)
├── exports        # Report exported symbols
├── refs           # Find references given a symbol
├── context        # Retrieve contextual snippet around position
├── def            # Jump to definition for symbol at position
├── query          # Execute raw Tree-sitter query (power users)
├── init           # Initialize project configuration (.atse.toml)
├── clear-cache    # Purge cached parse trees
├── stats          # Display cache/parser metrics
├── version        # Show version/build info
└── help           # Usage documentation (Cobra-provided)
```

### Global Flags (applied where relevant)
- `--recursive / -r` (default true) – traverse directories
- `--include` – glob(s) to include (e.g., `"*.ts"`, `"src/**/*.go"`)
- `--exclude` – glob(s) to exclude (e.g., tests, vendor directories)
- `--format` (`text` | `json` | `locations`) – output style (default `text`)
- `--verbose / -v` – include additional context (function bodies, etc.)
- `--limit` – max results to return
- `--no-cache` – bypass cache

### Output Philosophy
- Default human/agent-friendly plain text (locations only when possible).
- Optional JSON for machine consumption (`--format json`).
- `locations` format for piping to editors or scripts (`file:line:column`).
- Verbose mode enriches context without overwhelming default output.

## Architecture Overview
```
atse/
├── cmd/atse/main.go          # Cobra root command definition
├── internal/
│   ├── parser/
│   │   ├── manager.go        # Tree-sitter parser lifecycle + caching
│   │   ├── languages.go      # Grammar loading / detection
│   │   └── cache.go          # In-memory + optional disk cache plumbing
│   ├── search/
│   │   ├── structural.go     # Generic query execution
│   │   ├── highlevel.go      # find-fn/find-class wrappers
│   │   └── batch.go          # Parallel directory traversal utilities
│   ├── analysis/
│   │   ├── dependencies.go   # Imports/exports/dependency graph
│   │   ├── references.go     # Reference & definition resolution (stretch)
│   │   └── context.go        # Context extraction around positions
│   ├── output/
│   │   ├── formatter.go      # Text/JSON/location formatting
│   │   └── templates.go      # Verbose output helpers (if needed)
│   ├── config/
│   │   ├── config.go         # Project config (.atse.toml) management
│   │   └── defaults.go       # Default settings + flag wiring
│   └── util/
│       ├── paths.go          # Absolute path resolution, glob filters
│       └── workers.go        # Goroutine pools / concurrency helpers
├── grammars/                 # Embedded Tree-sitter grammars (.so / .dll / .dylib)
└── tests/
    ├── fixtures/             # Sample codebases for testing
    └── integration/          # CLI-level integration tests
```

### Key Components
- **Parser Manager** (`internal/parser`)
  - Maintains parsers per language.
  - Automatically detects language by file extension.
  - Caches syntax trees (with invalidation/refresh strategies).
  - Provides core operations: parse, search, list nodes, get snippet.

- **Search/Analysis Modules**
  - Build on manager primitives to implement high-level commands.
  - Provide language-specific heuristics for common tasks (imports, classes, etc.).

- **CLI Layer (Cobra)**
  - Wires commands to internal packages.
  - Handles flag parsing, output formatting selection, error messaging.

- **Output Formatting**
  - Centralized logic to ensure consistency and token efficiency.
  - Supports text, JSON, and location-only outputs.

- **Configuration & Caching**
  - `.atse.toml` for project defaults (languages, include/exclude patterns).
  - Cache directory (e.g., `.atse-cache/`) for persisted parse trees (future work).

## Testing Strategy
- **Unit Tests**
  - Parser manager (language loading, parse caching, invalidation).
  - High-level analyzers (dependencies, references, context extraction).
  - Output formatters.
- **Command Tests** (Cobra-level)
  - Validate flag parsing, error handling, and output contract for each command.
- **Integration Tests**
  - Run commands against fixtures (TypeScript, Go, Python, etc.).
  - Ensure consistent output across platforms (normalize paths/newlines).
- **Performance Benchmarks** (future)
  - Assess large directory traversal and caching behavior.

## Roadmap (Initial Focus)
1. **Scaffold Project**
   - Initialize Go module.
   - Wire Cobra root command with placeholder subcommands.
   - Add CI skeleton (lint/test).

2. **Parser Infrastructure**
   - Implement parser manager with language detection.
   - Load grammars (consider embedding).
   - Provide structural search helper (raw queries).

3. **Core Commands + Tests**
   - `find-fn`, `list-fns`, `deps`, `context`, `query`.
   - Unit tests for each command handler.
   - Integration tests covering TypeScript fixture.

4. **Output & Config Enhancements**
   - Implement `--format`, `--limit`, `--include`/`--exclude` support.
   - Add `.atse.toml` init and load.
   - Provide verbose mode and location-only format.

5. **Caching & Performance** (Subsequent iteration)
   - In-memory caching baseline.
   - Investigate on-disk cache for large repos.
   - Parallel file traversal and progress indicators.

6. **Extended Analysis (Future)**
   - Reference/definition resolution.
   - Project-wide dependency graphs.
   - Batch operations and summarization.

7. **Distribution & Docs**
   - Comprehensive README with walkthroughs.
   - Packaging: `go install`, Homebrew formula, release binaries.
   - Optional: thin MCP wrapper reusing CLI binary.

## Open Questions / Considerations
- How to manage embedded grammars (static linking vs. runtime download).
- Strategy for cross-language symbol resolution (requires semantic analysis).
- Handling mixed-language repositories (per-command language hints?).
- Balancing performance with accuracy for large codebases (e.g., incremental parsing).
- CLI UX for specifying positions (`file:line:column` vs separate flags).

## Next Steps
- Commit this `PLAN.md` as the foundation.
- Scaffold the Cobra CLI entry point.
- Implement parser manager core with go-tree-sitter.
- Develop core commands with corresponding tests and fixtures.

---

*Authored: 2025-11-12*

*Maintainer: Daniel Steigman (NightTrek)*
