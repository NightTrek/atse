# Literature Review: Code Search at Scale

This document summarizes key findings from academic and industry literature relevant to scaling ATSE to 10M+ files, with a focus on dependency analysis, change impact, and database-free indexing.

## 1. Industry Case Studies

### Google's Monorepo (Piper)
*   **Scale**: Billions of lines of code, 40k+ commits/day.
*   **Architecture**: Centralized, cloud-based filesystem.
*   **Build System**: Bazel (Blaze) - uses BUILD files to define explicit dependency graph.
*   **Search**: Kythe (based on Grok) - extracts semantic graph nodes.
*   **Key Lesson**: Explicit dependency declarations (BUILD files) are critical for scale. Inferred dependencies are too slow/brittle at this size.

### Meta (Facebook) Sapling & Buck2
*   **Scale**: Similar to Google.
*   **VCS**: Sapling - optimized for massive file trees, partial checkouts.
*   **Build System**: Buck2 - emphasizes remote caching and parallelism.
*   **Key Lesson**: "Graph-based" thinking permeates everything. The build graph IS the dependency graph.

### Sourcegraph (Zoekt)
*   **Approach**: Trigram-based regex search (Zoekt) + LSIF/SCIP for precise code navigation.
*   **Storage**: Sharded on-disk indexes.
*   **Key Lesson**: Trigrams are incredibly fast for "candidate finding" (better than ripgrep for massive scale? worth investigating).

---

## 2. Academic Research

### Compressed Sparse Row (CSR) for Graphs
*   **Concept**: standard format for sparse matrices/graphs.
*   **Relevance**: Highly efficient for storing the static call graph.
*   **Performance**: Traversals are cache-friendly linear memory scans.
*   **Constraint**: Hard to mutate. Best for read-only snapshots.

### Bloom Filters in Distributed Systems
*   **Application**: widely used to avoid expensive I/O.
*   **Benefit**: "Definitely No" answers prevent wasted disk reads.
*   **Relevance to ATSE**: Use Directory-level Bloom filters to skip parsing huge subtrees that don't contain relevant symbols.

### Incremental Static Analysis
*   **Paper**: "Salsa: Incremental Recompilation" (Rust Analyzer team).
*   **Concept**: Query-based architecture. `query(x)` depends on `query(y)`. Memoize results.
*   **Relevance**: Fundamental to ATSE's "Partial Index". We must formalize the invalidation logic.

---

## 3. Key Takeaways for ATSE

1.  **Abandon "Full Scans"**: At 10M files, you cannot stat() every file. We need a persistent index that knows what changed.
2.  **Hybrid Indexing**:
    *   **Trigrams/Bloom Filters** for fast candidate filtering.
    *   **CSR Graph** for the "core" dependency network (imports/calls).
    *   **Loose Objects** for detailed per-file AST data.
3.  **Explicit vs. Implicit**: We must lean on explicit signals (imports, build files) to prune the search space before attempting expensive implicit resolution (symbol matching).
