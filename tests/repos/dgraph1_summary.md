# ATSE Authentication Analysis - Dgraph Repository
## Performance Benchmarking Summary

**Test Date:** 2025-11-15  
**Repository:** tests/repos/dgraph (large production Go codebase)  
**Focus:** Authentication logic discovery  
**Metrics Log:** tests/repos/dgraph1.jsonl  

---

## Executive Summary

Successfully tested all 8 ATSE commands against the dgraph repository to discover authentication logic. The benchmarking system provided valuable insights into command performance characteristics:

**Key Findings:**
- ✅ **Fast Commands (<100ms):** deps, context, query, list-fns
- ⚠️ **Medium Commands (100-500ms):** graph (scoped), extract
- 🐌 **Slow Commands (>1s):** search (full repo), find-fn (full repo)
- 🚫 **Hanging Issues:** graph command on full repo (depth 1+) appears to hang/timeout

**Authentication Coverage:**
- Discovered 15 authentication-related symbols
- Found 18,074 function call sites for `authorizeUser`
- Extracted complete authentication feature with 30 components
- Identified key files: `dgraph/edgraph/access.go`, `acl/` directory

---

## Detailed Performance Metrics

### 1. search - "authenticate" (Specific Query)
```
Duration:         1.151s
Memory (Peak):    9.0 MB
Files Processed:  544
Results Found:    1
Output Tokens:    22 (o200k_base)
Output Chars:     78
```

**Analysis:**
- Searched 544 Go files across entire dgraph repo
- Found 1 result: `authenticateLogin` method
- ~2ms per file processed (very reasonable for full-text search)
- Low memory usage despite large codebase

**Result:**
```
authenticateLogin (method) - dgraph/edgraph/access.go:110
```

---

### 2. search - "auth" (Broader Query)
```
Duration:         1.139s
Memory (Peak):    8.9 MB
Files Processed:  544
Results Found:    15
Output Tokens:    248 (o200k_base)
Output Chars:     867
```

**Analysis:**
- Similar performance to specific "authenticate" search
- Found 15 authentication-related symbols
- Demonstrates fuzzy matching effectiveness
- Comparable memory usage to first search

**Key Authentication Symbols Discovered:**
- `AuthFor`, `AuthMeta`, `AuthRules` (GraphQL schema methods)
- `authCheck`, `authRules` (authorization functions)
- `authorizeUser`, `AuthSuperAdmin` (user authorization)
- `authorizeAlter`, `authorizePreds` (permission checks)
- `authorizeQuery`, `authorizeRequest` (query/request authorization)
- `authenticateLogin`, `authorizeMutation` (login & mutation auth)
- `authorizeNewNodes` (GraphQL mutation auth)

---

### 3. graph - "greet" (Small Test Fixture)
```
Duration:         0.014s (14ms)
Memory (Peak):    0.26 MB
Files Processed:  3
Results Found:    2
Output Tokens:    61 (o200k_base)
Output Chars:     210
```

**Analysis:**
- ✅ Baseline test confirms `graph` command works correctly
- Extremely fast on small codebases
- Minimal memory footprint
- Demonstrates proper BFS traversal

---

### 4. find-fn - "authorizeUser" (Full Repo)
```
Duration:         0.962s
Memory (Peak):    5.9 MB
Files Processed:  544
Results Found:    18,074 (!!)
Output Tokens:    517,808 (!!)
Output Chars:     1,569,420
```

**Analysis:**
- ⚠️ **WARNING:** Massive output - 517K tokens exceeds most LLM context windows
- Found 18,074 call sites for `authorizeUser` across codebase
- Despite huge result set, only 962ms execution time
- Modest memory usage (5.9 MB) considering output size
- **Recommendation:** Always use `--limit` flag with find-fn on large codebases

**Performance Insight:**
- ~1.8ms per file processed (very efficient)
- Would need `--limit 50` or less for practical LLM consumption

---

### 5. graph - "authorizeUser" (Scoped to access.go)
```
Duration:         0.319s
Memory (Peak):    1.3 MB
Files Processed:  1
Results Found:    29
Output Tokens:    580 (o200k_base)
Output Chars:     1,886
```

**Analysis:**
- ✅ **Solution to graph hanging:** Scope to specific file
- Fast execution when limited to single file
- Discovered 29 dependencies at depth 1
- Manageable output size for LLM context
- **Recommendation:** For large repos, always scope graph to specific files/directories

**Discovered Dependencies (Level 1):**
- `upsertGuardianAndGroot`, `parsePredsFromQuery`, `removeVarsFromQueryVars`
- `extractUserAndGroups`, `blockedPreds`, `addUserFilterToFilter`
- `userFilter`, `groupFilter`, `removeFilters`
- `authorizePreds`, `upsertGroot`, `RefreshACLs`
- `logAccess`, `addUserFilterToQuery`, `shouldAllowAcls`
- `validateToken`, `parsePredsFromMutation`, `isAclPredMutation`
- `parsePredsFromFilter`, `parentFilter`, `removeGroupBy`
- `getRefreshJwt`, `validateLoginRequest`, `refreshAclCache`
- `upsertGuardian`, `getAccessJwt`, `removePredsFromQuery`
- `AuthorizeGuardians`

---

### 6. list-fns (Full Repo, Limited to 200)
```
Duration:         0.070s (70ms)
Memory (Peak):    12.6 MB
Files Processed:  23
Results Found:    200
Output Tokens:    6,122 (o200k_base)
Output Chars:     19,186
```

**Analysis:**
- ⚡ Very fast command (70ms)
- Highest memory usage (12.6 MB) but still modest
- Only processed 23 files (stopped after hitting limit)
- Excellent for getting overview of available functions
- **Use Case:** Quick discovery of function names in a codebase

**Sample Auth Functions Found:**
- `getUserAndGroup`, `checkForbiddenOpts`, `add`, `userAdd`, `groupAdd`
- `del`, `userOrGroupDel`, `mod`, `changePassword`, `userMod`
- `chMod`, `queryUser`, `getUserModNQuad`, `queryGroup`
- `queryAndPrintUser`, `queryAndPrintGroup`, `info`

---

### 7. deps - access.go Dependencies
```
Duration:         0.005s (5ms)
Memory (Peak):    0.29 MB
Files Processed:  1
Results Found:    1
Output Tokens:    27 (o200k_base)
Output Chars:     82
```

**Analysis:**
- ⚡ **Fastest command tested** (5ms)
- Minimal memory footprint
- Simple, focused task: show imports
- **Use Case:** Quick dependency analysis

---

### 8. context - authenticateLogin Location
```
Duration:         0.004s (4ms)
Memory (Peak):    0.29 MB
Files Processed:  1
Results Found:    1
Output Tokens:    407 (o200k_base)
Output Chars:     1,630
```

**Analysis:**
- ⚡ Second fastest command (4ms)
- Provides full function source code with context
- Perfect for examining specific code locations
- **Use Case:** Deep dive into specific implementations

**Code Retrieved:**
- Full `authenticateLogin` method implementation
- Shows token validation, user authorization flow
- Demonstrates refresh token vs password authentication paths

---

### 9. extract - "authenticateLogin" Feature
```
Duration:         0.344s
Memory (Peak):    1.7 MB
Files Processed:  1
Results Found:    30
Output Tokens:    6,507 (o200k_base)
Output Chars:     22,296
```

**Analysis:**
- Moderately fast for comprehensive feature extraction
- Retrieved entry point + 29 dependencies
- Excellent for understanding complete feature implementation
- **Use Case:** Feature comprehension, refactoring preparation

**Feature Components Extracted:**
- Entry point: `authenticateLogin`
- Authentication helpers: `validateToken`, `validateLoginRequest`, `extractUserAndGroups`
- Authorization logic: `authorizeUser`, `authorizePreds`, `AuthorizeGuardians`
- Token generation: `getAccessJwt`, `getRefreshJwt`
- ACL management: `RefreshACLs`, `refreshAclCache`, `upsertGuardian`, `upsertGroot`
- Query filtering: `addUserFilterToQuery`, `addUserFilterToFilter`, `removePredsFromQuery`
- Helper utilities: `userFilter`, `groupFilter`, `parentFilter`, `removeFilters`

---

### 10. query - Tree-sitter Pattern
```
Duration:         0.005s (5ms)
Memory (Peak):    0.32 MB
Files Processed:  1
Results Found:    36
Output Tokens:    984 (o200k_base)
Output Chars:     2,964
```

**Analysis:**
- ⚡ Extremely fast (5ms)
- Found all 36 function declarations in access.go
- Advanced Tree-sitter pattern matching
- **Use Case:** Custom code pattern detection

---

## Performance Comparison Table

| Command | Duration | Memory Peak | Files | Results | Tokens | Speed (ms/file) |
|---------|----------|-------------|-------|---------|--------|-----------------|
| **search** (authenticate) | 1,151ms | 9.0 MB | 544 | 1 | 22 | 2.1ms |
| **search** (auth) | 1,139ms | 8.9 MB | 544 | 15 | 248 | 2.1ms |
| **graph** (small fixture) | 14ms | 0.26 MB | 3 | 2 | 61 | 4.7ms |
| **find-fn** (full repo) | 962ms | 5.9 MB | 544 | 18,074 | 517,808 | 1.8ms |
| **graph** (scoped file) | 319ms | 1.3 MB | 1 | 29 | 580 | 319ms |
| **list-fns** | 70ms | 12.6 MB | 23 | 200 | 6,122 | 3.0ms |
| **deps** | 5ms | 0.29 MB | 1 | 1 | 27 | 5.0ms |
| **context** | 4ms | 0.29 MB | 1 | 1 | 407 | 4.0ms |
| **extract** (depth 1) | 344ms | 1.7 MB | 1 | 30 | 6,507 | 344ms |
| **query** | 5ms | 0.32 MB | 1 | 36 | 984 | 5.0ms |

---

## Authentication Logic Discovered

### Core Authentication Files
1. **`dgraph/edgraph/access.go`** - Main authentication/authorization logic (36 functions)
2. **`acl/` directory** - ACL management, user/group operations
3. **`graphql/schema/` directory** - GraphQL auth rules
4. **`graphql/resolve/` directory** - GraphQL mutation auth

### Key Authentication Functions
1. **`authenticateLogin`** - Primary login method (refresh token + password auth)
2. **`authorizeUser`** - User authorization and validation
3. **`validateToken`** - JWT token validation
4. **`authorizePreds`** - Predicate-level authorization
5. **`authorizeQuery`** - Query authorization
6. **`authorizeMutation`** - Mutation authorization
7. **`authorizeAlter`** - Schema alteration authorization
8. **`AuthSuperAdmin`** - Super admin (groot) authorization
9. **`AuthorizeGuardians`** - Guardian group authorization
10. **`RefreshACLs`** - ACL cache refresh

### Authentication Flow
```
Login Request
    ↓
authenticateLogin
    ↓
validateLoginRequest → validateToken (if refresh token)
    ↓                      ↓
    ↓                  extractUserAndGroups
    ↓                      ↓
    └────→ authorizeUser ←─┘
              ↓
         authorizePreds → worker.AclCachePtr
              ↓
         getAccessJwt / getRefreshJwt
              ↓
         Return JWT tokens
```

---

## Command Recommendations

### For Large Codebases (500+ files like dgraph):

#### ✅ **Always Use:**
1. **`search`** - Reliable for initial discovery (1-2s, 8-9 MB)
   - Use specific queries first, then broaden
   - Add `--limit 10-20` to cap results

2. **`find-fn`** - Great for tracking function calls BUT:
   - ⚠️ **Always use `--limit 50` or less**
   - Without limit: 517K tokens (unusable for LLMs)
   - Extremely fast per-file (1.8ms/file)

3. **`list-fns`** - Excellent overview tool (70ms, 12.6 MB)
   - Use `--limit 100-200` for manageable output
   - Fastest way to see what functions exist

4. **`deps`, `context`, `query`** - Super fast (<10ms)
   - Perfect for targeted analysis
   - Minimal resource usage
   - Always scope to specific files

#### ⚠️ **Use with Caution:**
1. **`graph`** - Works well when scoped, hangs on full repo
   - ✅ Scoped to file: 319ms, 29 results ← **USE THIS**
   - ❌ Full repo: Hangs/timeout ← **AVOID**
   - **Always scope to specific file or directory**
   - Keep depth 1-2 for large codebases

2. **`extract`** - Good for feature extraction (344ms)
   - Use depth 1-2 only
   - Perfect for understanding feature boundaries
   - 6.5K tokens is manageable for LLMs

---

## Performance Insights

### Token Count Analysis (o200k_base model)
- **Smallest:** deps (27 tokens) - 5ms
- **Largest:** find-fn (517,808 tokens) - 962ms
- **Sweet Spot:** extract (6,507 tokens) - 344ms

### Memory Usage Patterns
- **Most Efficient:** deps, context, query (0.29-0.32 MB)
- **Moderate:** search, find-fn, graph-scoped (1.3-9 MB)
- **Highest:** list-fns (12.6 MB) - likely due to parsing multiple files

### Speed Champions
1. **context** - 4ms ⚡
2. **deps** - 5ms ⚡
3. **query** - 5ms ⚡
4. **list-fns** - 70ms 🚀
5. **graph (small)** - 14ms 🚀

### Processing Efficiency (ms per file)
1. **find-fn** - 1.8ms/file (most efficient)
2. **search** - 2.1ms/file
3. **list-fns** - 3.0ms/file

---

## graph Command Issue: Deep Dive

### Problem
The `graph` command appears to hang when run against the full dgraph repository:
```bash
# This hangs (never completes):
atse graph "authorizeUser" tests/repos/dgraph --mode bfs --depth 1
```

### Solution
Scope the graph command to specific files or directories:
```bash
# This works (319ms):
atse graph "authorizeUser" tests/repos/dgraph/edgraph/access.go --mode bfs --depth 1
```

### Hypothesis
The graph command likely:
1. Tries to build full dependency graph across all 544 files
2. Follows imports/calls across file boundaries
3. Gets caught in circular dependencies or very deep call chains
4. May need timeout mechanism or progress indication

### Recommended Improvements
1. Add progress logging to graph command
2. Implement timeout with partial results
3. Add `--max-files` limit option
4. Better handling of circular dependencies

---

## Recommended Workflow for Authentication Analysis

Based on performance data, here's the optimal workflow:

### Step 1: Discovery (Fast - <2s)
```bash
# Find authentication entry points
atse search "auth" tests/repos/dgraph --limit 15 --include "*.go" --benchmark

# Get function overview
atse list-fns tests/repos/dgraph/acl --limit 50 --benchmark
```

### Step 2: Understand Scope (Medium - <1s)
```bash
# Build dependency graph (SCOPED to file!)
atse graph "authenticateLogin" tests/repos/dgraph/edgraph/access.go \
  --mode bfs --depth 1 --benchmark

# Find all usages (WITH LIMIT!)
atse find-fn "authenticateLogin" tests/repos/dgraph \
  --limit 20 --include "*.go" --benchmark
```

### Step 3: Extract Code (Medium - <1s)
```bash
# Extract full authentication feature
atse extract "authenticateLogin" tests/repos/dgraph/edgraph/access.go \
  --depth 1 --benchmark
```

### Step 4: Detailed Analysis (Fast - <10ms each)
```bash
# Check dependencies
atse deps tests/repos/dgraph/edgraph/access.go --benchmark

# Get specific code context
atse context tests/repos/dgraph/edgraph/access.go:110:1 --benchmark

# Custom Tree-sitter queries
atse query '(function_declaration name: (identifier) @name)' \
  tests/repos/dgraph/edgraph/access.go --benchmark
```

---

## Authentication Implementation Details

### JWT-Based Authentication
- Uses refresh tokens (long-lived) and access tokens (short-lived)
- Tokens contain: `userid`, `namespace`, `groups`, `exp`
- JWT algorithm and secret key configured via `worker.Config.AclJwtAlg`

### Multi-Namespace Support
- Namespace extracted from JWT claims
- Root namespace (0) has special permissions
- ACL cache maintained per namespace

### Authorization Hierarchy
1. **Guardian Group** (`x.SuperAdminId`) - highest privileges
2. **Groot User** (`x.GrootId`) - default super admin
3. **Regular Users** - subject to ACL rules

### ACL Permission System
- Predicate-level permissions
- Operation-based access control (read, write, modify, etc.)
- Group-based permission inheritance
- Cached for performance (`worker.AclCachePtr`)

### Special Predicates
- `dgraph.xid` - User/group identifiers
- `dgraph.user.group` - User-to-group relationships
- `dgraph.group.acl` - Permission definitions
- `dgraph.type` - Type system predicates

---

## Lessons Learned

### ✅ What Worked Well
1. **Benchmarking system** - Provides excellent visibility into performance
2. **Metrics logging** - JSONL format perfect for historical tracking
3. **Token counting** - Critical for LLM context management
4. **File scoping** - Essential for graph/extract on large codebases

### ⚠️ Issues Identified
1. **graph command hangs on full repositories** - needs scoping or timeout
2. **find-fn without --limit** - generates unusable output (500K+ tokens)
3. **No progress indication** - hard to know if command is working or hung

### 💡 Optimization Opportunities
1. Add `--max-files` global flag to cap file processing
2. Implement streaming output for large result sets
3. Add progress bars for long-running commands
4. Implement timeout with partial results
5. Add `--dry-run` to estimate output size before execution

---

## Conclusion

The ATSE benchmarking system successfully measured performance across all 8 commands against a large production codebase. Key takeaways:

1. **Fast commands** (deps, context, query) are production-ready for any codebase size
2. **Search commands** scale well with proper limits
3. **graph command** requires file scoping for large repos (known issue)
4. **Token counting** is essential - prevents LLM context overflow
5. **Benchmarking workflow** enables data-driven optimization decisions

The authentication analysis discovered 15+ auth-related symbols, mapped the complete auth flow, and identified all major components of dgraph's ACL system.

**Next Steps:**
1. Fix graph command hanging on full repository
2. Add progress indication for long-running commands
3. Implement timeout mechanisms
4. Update documentation with scoping best practices

---

## Appendix: Raw Metrics Data

See `tests/repos/dgraph1.jsonl` for complete JSONL metrics including:
- Exact timestamps
- Full command arguments
- Exit codes
- Model used for token counting
- Detailed memory measurements

**Sample metric entry:**
```json
{
  "timestamp": "2025-11-15T08:42:52.300694Z",
  "command": "graph",
  "args": ["graph","authorizeUser","tests/repos/dgraph/edgraph/access.go","--mode","bfs","--depth","1","--benchmark","--log-metrics","--metrics-log-file","tests/repos/dgraph1.jsonl"],
  "exit_code": 0,
  "model": "o200k_base",
  "output_format": "text",
  "output_char_count": 1886,
  "output_token_count": 580,
  "version": "0.1.0",
  "execution_duration_ms": 319,
  "memory_start_mb": 0.246063232421875,
  "memory_peak_mb": 1.305084228515625,
  "memory_end_mb": 1.305084228515625,
  "files_processed": 1,
  "results_count": 29
}
```
