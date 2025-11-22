# Industry Practice: How Tech Giants Scale Code Search

To understand how to scale ATSE to 10M+ files, we analyzed publicly available information from companies operating at this scale (Google, Meta, Sourcegraph).

## 1. The "Monorepo" Mindset

At 10M files, you are no longer working with a "project"; you are working with a "universe."
*   **No Single Build**: You never build the whole repo.
*   **Partial Visibility**: Developers only check out a tiny fraction (Sparse Checkout / VFS).
*   **Graph is Truth**: The build graph (dependencies) determines what exists. Files not in the build graph effectively don't exist.

### Google (Piper + CitC + Kythe)
*   **Virtual Filesystem**: "Clients in the Cloud" (CitC). Developers see all files, but they are loaded on demand FUSE-style.
*   **Search (Kythe)**:
    *   Extracts semantic nodes (definitions, references) into a graph format.
    *   Stores graph in BigTable (Massive NoSQL).
    *   Query service runs on thousands of cores.
*   **Relevance to ATSE**: We can mimic the "Partial Visibility" by using regex (ripgrep) to find potential entry points, then expanding the graph only for reachable nodes.

### Meta (Sapling + EdenFS + Buck2)
*   **Virtual Filesystem**: EdenFS. Similar to Google, checks out files on demand.
*   **Dependency Tracking**: Buck2 query language (`buck2 query 'rdeps(//foo:bar)'`).
*   **Graph-Based Build**: The build system *is* the dependency graph.
*   **Relevance**: ATSE should integrate with build systems (Bazel/Buck/Make) where possible to get "free" dependency graphs, rather than re-parsing everything.

---

## 2. Key Scalability Techniques

### A. Content-Addressable Storage (CAS)
*   Everything is hashed (`sha256`).
*   If a file's content hash hasn't changed, its metadata/AST is reused.
*   **Application**: ATSE should store partial indexes keyed by file hash.
    *   `File A (Hash X)` -> `Index A (Hash Y)`
    *   If user edits `File A`, its hash changes -> compute `Index A'`.
    *   `File B` (unchanged) -> reuse `Index B`.

### B. Bloom Filters for "Negative" Confirmation
*   **Problem**: Checking if `MyFunction` is used in 100,000 files.
*   **Optimization**: Store a small Bloom filter for every directory.
*   **Query**: "Does `MyFunction` exist in `src/legacy`?" -> Bloom Filter says **NO**. -> Skip 50,000 files instantly.

### C. Trigram Indexing (Zoekt)
*   **Method**: Split code into 3-character chunks (`for`, `or `, `r i`, ` in`).
*   **Storage**: Inverted index mapping trigrams to file offsets.
*   **Speed**: Orders of magnitude faster than regex for literal string matching.
*   **Relevance**: Potential alternative to ripgrep for the initial "Discovery" phase in ATSE.

---

## 3. The "Local-First" Constraints

ATSE is a CLI tool, not a cloud service. We cannot spin up 1000 nodes.
*   **Constraint**: Memory is limited (16GB - 64GB typically).
*   **Constraint**: Disk I/O is the bottleneck.
*   **Constraint**: No persistent daemon (initially).

### Adaptation for ATSE
1.  **Persistent On-Disk Graph**: We cannot keep the whole graph in RAM. We must use a memory-mapped on-disk format (CSR or LMDB).
2.  **Lazy Loading**: Only map the parts of the graph relevant to the current query.
3.  **Incremental updates**: The graph must update incrementally. Rebuilding 10M node graph takes too long.
