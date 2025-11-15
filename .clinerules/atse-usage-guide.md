## Brief overview

This rule provides comprehensive guidelines for AI agents using the ATSE (Agent Tree Search Engine) CLI tool for code analysis. ATSE is a Tree-sitter powered command-line tool that provides syntax-aware, zero-false-positive code analysis across multiple languages (TypeScript, JavaScript, Python, Go). This guide incorporates Anthropic's prompting best practices to maximize effectiveness when analyzing codebases.

## Quick reference guide

**Available Commands (in recommended order):**
1. `atse search <query> [path]` - Fuzzy search for symbols (functions, classes, methods)
2. `atse graph <symbol> [path]` - Build dependency graph with BFS/DFS traversal
3. `atse extract <symbol> [path]` - Extract full source code for a feature
4. `atse find-fn <function> [path]` - Find all calls to a specific function
5. `atse list-fns [path]` - List all function declarations
6. `atse deps [path]` - Show dependencies (imports)
7. `atse context <file:line:col>` - Get code context at a position
8. `atse query <tree-sitter-query> [path]` - Execute raw Tree-sitter queries

**Global Flags:**
- `--format <text|json|locations>` - Output format
- `--limit <n>` - Limit number of results
- `--depth <n>` - Maximum traversal depth (graph/extract)
- `--mode <bfs|dfs>` - Traversal mode (graph/extract)
- `--verbose, -v` - Include additional context
- `--recursive, -r` - Recursively search directories (default: true)
- `--include <pattern>` - Include file patterns (e.g., "*.ts")
- `--exclude <pattern>` - Exclude file patterns (e.g., "*.test.ts")

**Performance Monitoring Flags:**
- `--benchmark` - Display performance summary to stderr (duration, memory, files processed, results count, tokens)
- `--log-metrics` - Log performance metrics to JSONL file for historical tracking
- `--metrics-log-file <path>` - Custom path for metrics log (default: benchmark/results/raw/token_metrics.jsonl)
- `--token-model <model>` - Token counting model name (default: gpt-4o)

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
- `deps` for understanding import relationships
- `context` for getting surrounding code at specific locations
- `query` for advanced Tree-sitter pattern matching

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
atse deps ./src/services/payment.ts --verbose
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

**Key Metrics Tracked:**
- **Duration**: Execution time in seconds (ms precision)
- **Memory Usage**: Start, peak, and end memory in MB
- **Files Processed**: Number of files analyzed
- **Results Found**: Number of symbols/matches returned
- **Output Tokens**: Token count for LLM context sizing
- **Output Chars**: Character count

### Using --log-metrics flag

The `--log-metrics` flag appends performance data to a JSONL file for historical analysis:

```bash
atse graph AuthService ./src --log-metrics
```

**JSONL Entry Format:**
```json
{
  "timestamp": "2025-11-15T08:10:29.296742Z",
  "command": "graph",
  "args": ["graph", "AuthService", "./src"],
  "exit_code": 0,
  "execution_duration_ms": 1247,
  "memory_start_mb": 12.3,
  "memory_peak_mb": 52.1,
  "memory_end_mb": 45.3,
  "files_processed": 127,
  "results_count": 23,
  "output_token_count": 42,
  "output_char_count": 145,
  "model": "gpt-4o",
  "output_format": "text",
  "version": "0.1.0"
}
```

### Combined Usage

Use both flags together for immediate feedback and historical tracking:

```bash
atse graph AuthService ./src --benchmark --log-metrics
```

### Benchmarking Workflow

**For comparing command performance:**
```bash
# Benchmark different depth settings
atse graph AuthService ./src --depth 1 --benchmark
atse graph AuthService ./src --depth 2 --benchmark
atse graph AuthService ./src --depth 3 --benchmark

# Compare BFS vs DFS traversal
atse graph AuthService ./src --mode bfs --benchmark
atse graph AuthService ./src --mode dfs --benchmark
```

**For tracking improvements over time:**
```bash
# Log metrics before optimization
atse graph AuthService ./src --log-metrics

# ... make code improvements ...

# Log metrics after optimization
atse graph AuthService ./src --log-metrics

# Analyze the JSONL file to compare
jq 'select(.command=="graph") | {duration: .execution_duration_ms, memory: .memory_peak_mb}' \
  benchmark/results/raw/token_metrics.jsonl
```

### AI Agent Recommendations

When using ATSE as an AI agent:

1. **Use --benchmark for immediate insight**: When exploring a codebase, add `--benchmark` to understand command performance and adjust parameters accordingly
2. **Monitor memory usage**: If memory usage is high (>100 MB), reduce `--depth` or scope to specific directories
3. **Track token counts**: Use token count to estimate LLM context window usage before sending results
4. **Compare alternatives**: Benchmark different approaches (e.g., search vs. graph) to find the most efficient solution

### Performance Analysis Examples

**Analyzing slow commands:**
```bash
# If a graph command is slow, reduce depth
atse graph LargeService ./src --depth 1 --benchmark  # Try depth 1 first

# Or scope to specific files
atse graph LargeService ./src --include "*.ts" --exclude "*.test.*" --benchmark
```

**Tracking optimization impact:**
```bash
# Before optimization
atse search authenticate ./src --log-metrics --benchmark

# After adding indexes or caching
atse search authenticate ./src --log-metrics --benchmark

# Compare results
tail -2 benchmark/results/raw/token_metrics.jsonl | jq .execution_duration_ms
```

## Performance optimization tips

**For large codebases (10,000+ files):**
- Always use `--include` to filter by file type: `--include "*.go"`
- Use `--exclude` to skip test files: `--exclude "*.test.*"`
- Keep `--depth` low (1-2) for graph and extract commands
- Use `--limit` to cap results at reasonable numbers (10-50)
- Scope searches to specific directories rather than project root

**For graph/extract timeouts:**
- Reduce `--depth` parameter
- Use more specific symbol names
- Search within single files rather than entire directories
- Consider using `find-fn` or `search` as alternatives for simpler queries

## Integration with chain of thought

When using ATSE in complex analysis tasks, structure your thinking:

```
<thinking>
1. What is the core question I need to answer?
2. What ATSE command best addresses this question?
3. What parameters optimize for my specific need?
4. What do the results tell me?
5. Do I need follow-up commands?
</thinking>

<action>
[Execute ATSE command with specific parameters]
</action>

<analysis>
[Interpret results and determine next steps]
</analysis>
```

## Error handling and debugging

**Common issues and solutions:**

1. **Command times out:**
   - Reduce --depth or --limit
   - Scope to specific directories
   - Use simpler commands (search instead of graph)

2. **No results found:**
   - Check capitalization (symbol names are case-sensitive)
   - Try --fuzzy flag with search
   - Verify file patterns with --include match your target language
   - Use --verbose to see what files are being processed

3. **Too many results:**
   - Add --limit flag
   - Use --include/--exclude for file filtering
   - Be more specific with symbol names
   - Use --type flag with search command

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

### 2025-11-13 - Initial Comprehensive Testing Complete
**Context**: Tested all 8 ATSE commands against dgraph repository (large production Go codebase)
**Pattern**: Following the recommended workflow (search → graph → extract → detailed analysis) provided the most effective understanding of features
**Result**: All commands work reliably with appropriate parameter tuning
**Recommendation**: 
- For graph command on large codebases, start with --depth 1-2
- search command with --limit 5-10 is perfect for initial exploration
- extract command requires modest depth (1-2) to avoid extremely large outputs
- Always scope large codebase analysis to specific directories when possible

### 2025-11-15 - Performance Monitoring System Added
**Context**: Added comprehensive performance monitoring to track execution time, memory usage, and workload metrics
**Pattern**: Using `--benchmark` flag provides immediate performance feedback; `--log-metrics` enables historical tracking
**Result**: Can now benchmark command performance and identify optimization opportunities
**Recommendation**:
- Always use `--benchmark` when testing different parameter combinations (depth, mode, filters)
- Use `--log-metrics` to track performance improvements over time
- Monitor memory peak to detect when commands might timeout or use excessive resources
- Token counts help estimate LLM context window usage before sending large outputs
- Fast commands (<100ms) are ideal for iterative exploration; slower commands (>1s) may need parameter tuning

### 2025-11-15 - Graph Command Enhanced with Symbol Indexing (CRITICAL FIX)
**Context**: Graph command was hanging indefinitely on large codebases (500+ files) due to O(n²) file searching complexity
**Pattern**: Implemented upfront symbol indexing that parses all files once, then uses O(1) lookups for graph traversal
**Result**: Graph command now handles any sized codebase efficiently:
- **Before**: Hung on dgraph repo (547 files)
- **After**: Completes in 2.5s with full progress visibility
**New Features**:
- `--connection-types <calls|types|imports|dataflow|events|all>` - Choose which relationships to follow (default: calls)
- `--max-symbols <n>` - Limit symbol discovery to prevent runaway (default: 500, use 0 for unlimited)
- `--verbose` now shows index building progress (every 50 files) and traversal progress (each symbol)
**Recommendation**:
- **Always use `--verbose` on large codebases** to see progress and confirm it's working
- For initial exploration: `atse graph <symbol> <path> --depth 1 --verbose --benchmark`
- For deep analysis: increase depth gradually while monitoring with `--benchmark`
- Use `--max-symbols` to control scope: smaller for quick checks, larger for comprehensive analysis
- On huge codebases (1000+ files): scope to specific directories or use `--include "*.go"` filters
- The command no longer hangs - if it seems slow, `--verbose` will show it's making progress
**Performance**: 547 files indexed in ~2.5s, 31MB peak memory, scales linearly

### [Add your learnings below]

---
