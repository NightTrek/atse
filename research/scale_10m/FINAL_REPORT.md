# Final Report: Scaling ATSE to 10M+ Files

This report synthesizes the research conducted on scaling the Agent Tree Search Engine (ATSE) to handle "absolutely enormous" repositories (10M+ files) without relying on a centralized database daemon.

## 1. Executive Summary

Scaling to 10M+ files requires moving beyond the in-memory "Lazy Hybrid" architecture of v0.2.0. We propose a **"Filesystem-as-Database"** architecture that uses content-addressable storage, compressed sparse row (CSR) graphs, and Bloom filters to achieve persistent, incremental, and fast code analysis.

**Key Recommendation**: Adopt a **hybrid storage model** where granular file data is stored in immutable "loose objects" (Cap'n Proto) and the global dependency graph is stored in a memory-mapped CSR format.

## 2. Research Documents

Detailed findings are organized in the following documents:

*   **[01_literature_review.md](01_literature_review.md)**: Academic researchers and industry case studies (Google, Meta, Sourcegraph). Key takeaways: Explicit dependency graphs are critical; Bloom filters prevent wasted I/O.
*   **[02_industry_practice.md](02_industry_practice.md)**: Deep dive into how tech giants manage monorepos. Key takeaways: Virtual filesystems, graph-based build systems (Buck2), and trigram indexing.
*   **[03_architecture_comparison.md](03_architecture_comparison.md)**: Trade-off analysis of architectural patterns. Recommendation: "Filesystem as Database" over centralized DBs or distributed sharding.
*   **[04_storage_research.md](04_storage_research.md)**: Technical evaluation of storage formats. Recommendation: Cap'n Proto for file fragments, CSR for the graph topology, LMDB for symbol lookups.
*   **[05_proposed_solutions.md](05_proposed_solutions.md)**: The concrete roadmap for ATSE v0.3.0. Definitions of the `.atse` directory structure and incremental update workflow.

## 3. The Proposed Architecture (ATSE v0.3.0)

### The Data Store (`.atse/`)
Instead of a single `sqlite.db` file, we use a structured directory:
*   `/objects/`: Immutable binary fragments of parsed files (AST summary + imports). Keyed by content hash.
*   `/index/`: Mutable mapping of `FilePath -> ContentHash`.
*   `/graph/`: Read-optimized Compressed Sparse Row (CSR) adjacency list of the dependency graph.
*   `/filters/`: Directory-level Bloom filters for fast symbol existence checks.

### The Workflow
1.  **Incremental Scan**: Ripgrep scans for changed files.
2.  **Parallel Parse**: Worker pool parses changed files into "Loose Objects".
3.  **Graph Patching**: The global CSR graph is patched (or rebuilt if changes are extensive) with new edges.
4.  **Zero-Copy Query**: Queries `mmap` the CSR graph and Symbol Index for instant traversals.

## 4. Ensuring Coverage & "Blast Radius" Accuracy

To guarantee we find *all* dependencies (the core user requirement):
1.  **Explicit Imports**: Use the build graph / import statements as the ground truth backbone.
2.  **Global Symbol Index**: Resolve cross-module calls using a persistent symbol table.
3.  **Verification**: When a dependency is found, perform a localized check (Tree-sitter) to confirm it's a true positive, pruning the graph of false positives.

## 5. Testing Strategy

We cannot check in a 10M file repo. We must generate one.
*   **Synthetic Repo Generator**: A tool to create 10M files with a realistic, power-law distributed dependency graph.
*   **Benchmark Suite**: Measures "Cold Start" (indexing time) vs. "Warm Query" (impact analysis time) vs. "Incremental Update" (re-index time).

## 6. Next Steps

1.  **Prototype the Object Store**: Implement the `internal/store` package to write Cap'n Proto objects.
2.  **Build the Generator**: Create the synthetic 10M file repo generator.
3.  **Measure Baselines**: Run current ATSE against the synthetic repo to establish the baseline to beat.
