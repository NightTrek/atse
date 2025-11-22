## Brief overview

This rule provides comprehensive guidelines for AI agents using the ATSE (Agent Tree Search Engine) CLI tool for code analysis. ATSE is a Tree-sitter powered command-line tool that provides syntax-aware, zero-false-positive code analysis across multiple languages (TypeScript, JavaScript, Python, Go). This guide incorporates Anthropic's prompting best practices to maximize effectiveness when analyzing codebases.

## Quick reference guide

**Available Commands (in recommended order):**
1. `atse search <query> [path]` - Hybrid search for symbols (ripgrep + tree-sitter)
2. `atse graph <symbol> [path]` - Build dependency graph with lazy expansion
3. `atse extract <symbol> [path]` - Extract full source code for a feature
4. `atse find-fn <function> [path]` - Find all calls to a specific function
5. `atse context <file:line:col>` - Get code context at a position

**Simplified Global Flags:**
- `--format <text|json|locations>` - Output format
- `--limit <n>` - Limit number of results
- `--verbose, -v` - Include additional context
- `--include <pattern>` - Include file patterns (e.g., "*.ts")
- `--exclude <pattern>` - Exclude file patterns (e.g., "*.test.ts")
- `--benchmark` - Display performance summary

**Removed Flags:**
- `--recursive` - Always recursive
- `--production-only` - Use exclude
- `--include-tests` - Use include
- `--rebuild-index` - No longer needed with hybrid architecture

## Recommended workflow

**Step 1: Discovery (search)**
- Always start with `atse search` to find entry points
- Use fuzzy matching (`--fuzzy`) for exploratory searches
- Filter by type with `--type` flag when you know what you're looking for

**Step 2: Understand Scope (graph)**
- Use `atse graph` to discover feature boundaries
- Start with `--mode bfs --depth 2` for broad understanding
- Use `--mode dfs --depth 3+` for deep call chains
- Add `--bidirectional` to see both dependencies and dependents

**Step 3: Extract Code (extract)**
- Use `atse extract` to get full source code of discovered features
- Keep depth low (1-2) for large codebases to avoid timeouts
- Use `--output` flag to save to file for further analysis

**Step 4: Detailed Analysis (other commands)**
- `find-fn` for finding specific function calls
- `context` for getting surrounding code at specific locations

## System prompt approach

**Role Definition:**
When using ATSE, adopt the role of a senior software architect analyzing a codebase. Think step-by-step about what information you need and which ATSE command best provides that information.

**Chain of Thought Process:**
1. **Identify the task**: What am I trying to understand about this codebase?
2. **Choose the right command**: Which ATSE command(s) will give me that information?
3. **Start broad, then narrow**: Begin with search/graph, then use specific commands
4. **Adjust parameters**: If results are too many/few, adjust --limit, --depth, or --include/--exclude
5. **Verify and iterate**: Check if results answer the question; if not, try different approach

## Use case examples

### Use Case 1: Understanding a new feature
**Goal**: Understand how authentication works in a codebase

**Approach with thinking:**
```
<thinking>
I need to understand the authentication feature. Let me:
1. Search for authentication-related symbols
2. Build a dependency graph to see the full scope
3. Extract the source code to analyze implementation
</thinking>

Commands:
1. atse search "authenticate" ./src --limit 10
2. atse graph "AuthService" ./src --mode bfs --depth 2
3. atse extract "AuthService" ./src --depth 2 --output auth-feature.txt
```

### Use Case 2: Finding where a function is called
**Goal**: Track down all usages of a deprecated function

**Approach:**
```
atse find-fn "deprecatedMethod" ./src --format json
```
Use JSON format for programmatic processing of results.

### Use Case 3: Analyzing dependencies
**Goal**: Understand what external libraries a module depends on

**Approach:**
```
rg "^import" ./src/services/payment.ts
```

### Use Case 4: Code refactoring preparation
**Goal**: Identify all components that need updating when changing an interface

**Approach with structured output:**
```
<thinking>
To prepare for refactoring the UserService interface:
1. Find the interface definition location
2. Find all implementations and usages
3. Build a dependency graph to see impact scope
</thinking>

Commands:
1. atse search "UserService" ./src --type class
2. atse find-fn "UserService" ./src --limit 50
3. atse graph "UserService" ./src --bidirectional --depth 2
```

## Performance monitoring and benchmarking

**NEW: Performance Tracking (Added 2025-11-15)**

ATSE now includes comprehensive performance monitoring to help you benchmark command execution and track improvements over time.

### Using --benchmark flag

The `--benchmark` flag displays a beautiful performance summary to stderr without interfering with stdout output:

```bash
atse graph AuthService ./src --benchmark
```

**Output:**
```
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
⚡ Performance Benchmark
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Command:          graph
Duration:         1.247s
Memory (Start):   12.3 MB
Memory (Peak):    52.1 MB
Memory (End):     45.3 MB
Files Processed:  127
Results Found:    23
Output Tokens:    42 (gpt-4o)
Output Chars:     145
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
```

## Iterative learning section

**How to add learnings:**
When you discover a pattern that works particularly well or avoid a pitfall, add it below with the date and a brief description. This helps future iterations of AI agents using this tool.

**Learning Template:**
```
### [Date] - [Brief Title]
**Context**: [What were you trying to do?]
**Pattern**: [What approach worked/didn't work?]
**Result**: [What was the outcome?]
**Recommendation**: [What should future agents do?]
```

---

### 2025-11-21 - Hybrid Architecture Upgrade (Massive Scale Support)
**Context**: Scaling to 4,000+ files and monorepos.
**Pattern**: Replaced full indexing with hybrid ripgrep + Tree-sitter lazy loading.
**Result**: 
- **Search**: Instant (500ms) results from ripgrep, structured parse on-demand
- **Graph/Extract**: Lazy expansion means only relevant files are parsed
- **Memory**: Dropped from GBs to MBs for large repos
**Recommendation**: 
- Ensure `rg` is installed
- Use ATSE for structural analysis, but prefer `rg` for simple text search
- Hybrid mode is automatic - just use commands as normal

### [Add your learnings below]

---
