# ATSE Benchmark Results (Dgraph Codebase)

These results were collected by running ATSE against the full Dgraph codebase (547 Go files).

| ID | Test Case | Command | Result | Performance | Verdict |
|---|---|---|---|---|---|
| **1** | **Discovery (Limit)** | `search "Test"` | Found 2255, Truncated to 20 | **1.19s** | ✅ PASS (Warning shown) |
| **2** | **Discovery (Precision)** | `search "Process"` | Found 61, Truncated to 5 | **1.15s** | ✅ PASS |
| **3** | **Traversal (Deep)** | `graph "Process"` | Built graph (depth 3) | **2.53s** | ✅ PASS (vs 75s+ baseline) |
| **4** | **Traversal (Wide)** | `graph "authenticateLogin"` | Built graph (depth 2) | **2.68s** | ✅ PASS |
| **5** | **Extraction** | `extract "Process"` | Extracted 4 components | **2.53s** | ✅ PASS (vs 20s+ baseline) |
| **6** | **Filtering (Prod)** | `find-fn "Process" --prod` | Processed **343** files (tests excluded) | **0.65s** | ✅ PASS (vs 547 files) |
| **7** | **Filtering (Reg)** | `find-fn "Process"` | Processed **547** files | **1.02s** | ✅ PASS (Baseline comparison) |
| **8** | **File Analysis** | `list-fns query.go` | Listed 33 functions | **0.017s** | ✅ PASS |
| **9** | **Dependencies** | `deps server.go` | Listed imports | **0.008s** | ✅ PASS |
| **10** | **Edge Case** | `search "NonExistent"` | 0 results | **1.16s** | ✅ PASS (Graceful) |

## Key Improvements Verified

1.  **Token Safety**:
    - `find-fn` results truncated to 20 (from 10,000+ potential).
    - Warnings clearly displayed (`⚠️ Warning: Results truncated...`).

2.  **Performance**:
    - **Graph Build**: 28x faster (2.6s vs 75s).
    - **Extraction**: 8x faster (2.5s vs 20s).
    - **Filtering**: 37% I/O reduction by skipping test files (343 vs 547 files).

3.  **Usability**:
    - Output categorized into `[Production Code]`.
    - "Instant" feedback for file-level operations (<20ms).
