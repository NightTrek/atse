# Storage Research: Persistence for 10M+ Files

This document explores storage strategies for persisting ATSE's dependency graph and partial indexes, enabling incremental updates and fast queries without a running database daemon.

## 1. The Storage Problem

At 10M+ files, the dependency graph is too large to re-compute from scratch on every run (the "Lazy Hybrid" approach hits its limit). We need a way to **store** the graph on disk so that:
1.  **Read speed is near-instant** (memory-mapped).
2.  **Updates are incremental** (only change affected parts).
3.  **No daemon is required** (simple file I/O).

### Data to Store
*   **Symbol Table**: Map `SymbolName` -> `[FileLocation]`
*   **Dependency Graph**: File A -> imports -> File B -> calls -> Function C
*   **Metadata**: File hashes, timestamps (for invalidation)

---

## 2. Candidate Technologies

### A. LMDB (Lightning Memory-Mapped Database)
*   **Mechanism**: Uses `mmap` to map a B+Tree directly into memory. OS manages paging.
*   **Pros**:
    *   **Zero-copy**: Data is read directly from the OS cache.
    *   **ACID**: Safe against crashes.
    *   **Key-Value**: Flexible schema.
*   **Cons**:
    *   **Write amplification**: Random writes can be slow.
    *   **Sparse files**: Can consume large disk space if not careful.
*   **Fit for ATSE**: High. Excellent for the "Symbol Table" (fast lookups).

### B. Compressed Sparse Row (CSR)
*   **Mechanism**: Two arrays. `OFFSETS` array points to ranges in `EDGES` array.
*   **Pros**:
    *   **Ultra-compact**: Minimal memory overhead (just integers).
    *   **Cache-friendly**: Sequential memory access.
    *   **Ideal for Read-Only**: Perfect for shipping a pre-computed graph.
*   **Cons**:
    *   **Immutable structure**: Cannot easily insert edge A->B without shifting the whole array.
*   **Fit for ATSE**: High for the "Graph Topology", *if* we treat the graph as an append-only or rebuild-on-change structure.

### C. Cap'n Proto / FlatBuffers
*   **Mechanism**: Serialization formats that are same in-memory as on-wire/disk. Zero parsing cost.
*   **Pros**:
    *   **Typed**: Strong schema guarantees.
    *   **Zero-parsing**: `mmap` and access struct fields directly.
*   **Cons**:
    *   Not a database; no indexing (need to build B-Trees manually).
*   **Fit for ATSE**: Good for storing the "Per-File Fragment" (the partial index of a single file).

---

## 3. Proposed "Hybrid Storage" Architecture

We propose a hybrid storage model that combines the strengths of the above:

### Layer 1: The "Loose Object" Store (Per-File)
*   **Format**: Cap'n Proto or dense binary.
*   **Location**: `.atse/objects/xx/yyyy...` (content-addressable).
*   **Content**: The AST summary and local dependency graph for **one source file**.
*   **Update Strategy**: When `file.go` changes, re-parse it, hash it, write new object. Old object garbage collected later.

### Layer 2: The "Global Index" (Graph Topology)
*   **Format**: LMDB or a custom mutable CSR.
*   **Content**:
    *   **Symbol Index**: `Symbol -> List<ObjectHash>`
    *   **Import Graph**: `FileHash -> List<ImportedFileHash>`
*   **Update Strategy**:
    *   Read changed "Loose Objects".
    *   Update entries in LMDB (transactional).

### Layer 3: The "Bloom Filter" Gatekeeper
*   **Format**: Single binary file (bit array).
*   **Content**: Bloom filter of all defined symbols in the repo.
*   **Usage**: Fast negative lookups ("Is `MyFunction` defined in this repo?").
*   **Update**: Rebuilt periodically or updated incrementally if filter supports it.

---

## 4. Storage Schema Design (Draft)

### Object File Schema (Cap'n Proto)
```capnp
struct FileIndex {
  path @0 :Text;
  contentHash @1 :Data;
  symbols @2 :List(Symbol);
  imports @3 :List(Import);
  
  struct Symbol {
    name @0 :Text;
    kind @1 :SymbolKind; # Function, Class, etc.
    range @2 :Range;
  }
  
  struct Import {
    path @0 :Text; # Or module name
    range @1 :Range;
  }
}
```

### Global Index Schema (LMDB Key-Value)
*   **Key**: `S:<SymbolName>` -> **Value**: `[FileHash1, FileHash2]`
*   **Key**: `F:<Path>` -> **Value**: `FileHash` (Current state)
*   **Key**: `D:<FileHash>` -> **Value**: `[DepHash1, DepHash2]` (Cached edges)

---

## 5. Benchmarking & Validation Plan

To validate this architecture, we need to benchmark:
1.  **Write Speed**: Time to write 10k small "Loose Objects" vs 1 big LMDB write.
2.  **Read Speed**: Time to `mmap` and traverse a 10M node CSR vs LMDB lookups.
3.  **Size**: Disk footprint of CSR vs LMDB vs Raw JSON.

### Expected Outcomes
*   **Loose Objects** will be fastest for incremental builds (granularity).
*   **CSR** will be fastest for global graph analysis (PageRank, Reachability).
*   **LMDB** will be the best compromise for a queryable index.
