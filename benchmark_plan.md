# ATSE Benchmark & Regression Test Plan (Dgraph)

This plan defines 10 test cases to validate the performance and correctness of the ATSE optimizations against the Dgraph codebase.

## Test Cases

| ID | Type | Command | Goal | Expected Result |
|---|---|---|---|---|
| **1** | **Discovery (Limit)** | `atse search "Test" . --limit 20` | Verify default limits work on broad queries. | Truncated to 20 results, warning shown. |
| **2** | **Discovery (Precision)** | `atse search "processMutation" .` | Find specific entry points. | Accurate findings, sub-2s time. |
| **3** | **Traversal (Deep)** | `atse graph "processMutation" . --depth 3` | Stress test deep graph traversal with new index. | Sub-3s execution (vs 75s+ baseline). |
| **4** | **Traversal (Wide)** | `atse graph "Server" . --mode bfs --depth 2` | Test BFS on a high-connectivity node. | Fast traversal, no O(N^2) slowdown. |
| **5** | **Extraction** | `atse extract "processMutation" . --depth 1` | Verify source extraction speed. | Instant extraction after index build. |
| **6** | **Filtering (Prod)** | `atse find-fn "Error" . --production-only --limit 20` | Verify noise reduction logic. | ZERO results from `*_test.go` files. |
| **7** | **Filtering (Regression)**| `atse find-fn "Error" . --limit 20` | Confirm default behavior finds tests (comparison). | Results include `*_test.go` files. |
| **8** | **File Analysis** | `atse list-fns dgraph/query/query.go` | Baseline check for single-file parsing. | Instant (<100ms), list all functions. |
| **9** | **Dependencies** | `atse deps dgraph/edgraph/server.go` | Verify import resolution speed. | Instant (<50ms), list imports. |
| **10** | **Edge Case** | `atse search "NonExistentSymbol123" .` | Verify graceful failure. | 0 results, no crash, clean output. |

## Execution Plan

We will execute these commands against `/Users/danielsteigman/code/randomSource/dgraph` using the `--benchmark` flag to capture precise metrics.
