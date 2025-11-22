# Proposed Solutions: ATSE v0.3.0 for 10M+ Files

This document synthesizes the research into a concrete roadmap for the next major version of ATSE, with a focus on scalability, database-free persistence, and complete dependency coverage.

## 1. The "Database-Free" Architecture

To scale to 10M+ files without a running database daemon, we will adopt a hybrid "Content-Addressable Filesystem" approach.

### Components

*   **`.atse/objects/`**: A directory storing immutable, content-addressed binary fragments (Cap'n Proto).
    *   Values: `sha256(file_content) -> { symbols: [...], imports: [...] }`
    *   This layer is purely functional: `code -> graph_fragment`. It is cacheable and shareable.

*   **`.atse/index/`**: A mutable index mapping current file paths to object hashes.
    *   `file_path -> current_content_hash`
    *   This layer handles file changes. It's a lightweight LMDB or binary map.

*   **`.atse/graph/`**: A read-optimized Compressed Sparse Row (CSR) graph topology.
    *   `topology.csr`: Stores the current dependency graph structure.
    *   Rebuilt incrementally or fully when "too much" has changed.

### Workflow

1.  **Scan**: Ripgrep finds all source files. Check modtimes against `.atse/index`.
2.  **Parse**: For modified files, re-parse with Tree-sitter.
3.  **Hash**: Compute content hash. Write new fragment to `.atse/objects` if missing.
4.  **Update**: Update `.atse/index` with new hashes.
5.  **Link**: Rebuild or patch `.atse/graph/topology.csr` with new edges.
6.  **Query**: Load CSR into memory (mmap). Run graph algorithms.

---

## 2. Ensuring Complete Dependency Coverage

The core problem is "How do I know I found *all* dependencies?" especially with dynamic features or implicit interfaces.

### Solution: Multi-Layered Reachability Analysis

We propose a layered approach to guarantee coverage:

**Layer 1: Explicit Imports (100% Accuracy)**
*   Source code `import` statements are ground truth.
*   ATSE will recursively resolve imports to build a "Module Graph".
*   This narrows the search space from 10M files to the relevant subset of modules.

**Layer 2: Symbol References (High Accuracy)**
*   Use the **Global Symbol Table** (built from `.atse/objects`) to resolve cross-module calls.
*   `function Foo()` called by `module B` -> Edge A->B.

**Layer 3: Heuristic & Textual Search (Fallback)**
*   For dynamic languages (Python, JS), use **Bloom Filters** to prune search.
*   If looking for `my_dynamic_method`, check Bloom filters of all reachable modules.
*   If filter says "maybe", use ripgrep to verify.

### "Blast Radius" Calculation
When a file changes:
1.  Identify changed Entry Points (functions/classes).
2.  Perform **Reverse Reachability** on the CSR graph.
3.  Collect all `[Affected Files]`.
4.  **Verification**: Run localized Tree-sitter queries on `[Affected Files]` to confirm they actually use the changed symbol (prune false positives).

---

## 3. Implementation Roadmap

### Phase 1: The Object Store (v0.3.0)
*   Implement `internal/store`: A content-addressable storage engine.
*   Define Cap'n Proto schema for `FileFragment`.
*   Modify `extract` command to write to this store.

### Phase 2: The CSR Graph Engine (v0.3.1)
*   Implement `internal/graph/csr`: In-memory CSR graph builder.
*   Implement `internal/graph/algo`: BFS/DFS and Reachability algorithms on CSR.
*   Benchmark against current pointer-based graph.

### Phase 3: The Incremental Engine (v0.4.0)
*   Implement file watching / modtime checking.
*   Wire up the "Scan -> Parse -> Hash -> Link" pipeline.
*   Goal: <100ms response for queries on a 10M file repo (hot state).

---

## 4. Testing Strategy for Scale

### Synthetic Repo Generator
We need a tool to generate "fake" 10M file repos with realistic topology.
*   **Parameters**: Depth, Breadth (fan-out), Inter-module dependencies.
*   **Output**: 10M `.go` or `.py` files with valid imports and function calls.

### Benchmark Suite
*   **Cold Start**: Time to index 10M files from scratch.
*   **Warm Query**: Time to find "impact of changing file X".
*   **Incremental Update**: Time to update index after modifying 1 file.
*   **Disk Usage**: Size of `.atse` directory.
