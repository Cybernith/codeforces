# Shortest Path Queries (Dijkstra in Go) - soroosh morshedi

This repository contains a Go implementation of **shortest path queries on a directed weighted graph** using **Dijkstra’s algorithm** with simple caching per source node.

The main entrypoint is:

- `answer.go` – reads a graph and multiple queries from standard input, then prints the shortest distance and one shortest path for each query.

---

## 🧩 Problem Description

You are given a **directed weighted graph**:

- `n` nodes, numbered from `1` to `n`
- `m` edges
- then `q` queries

Each edge is:

```text
u v w
```

meaning there is an edge from node `u` to node `v` with weight `w` (positive integer).

Each query is:

```text
s t
```

For each query, the program must:

1. Compute the **shortest distance** from `s` to `t`.
2. If no path exists, print `-1`.
3. If a path exists:
   - First line: the minimum distance `d`
   - Second line: an integer `k` followed by `k` nodes – one valid shortest path from `s` to `t`.

---

## 📥 Input Format

From `stdin`:

```text
n m q
u1 v1 w1
u2 v2 w2
...
um vm wm
s1 t1
s2 t2
...
sq tq
```

- `1 ≤ u, v ≤ n`
- Edge weights `w` are positive.
- Nodes are 1-based in the input.

---

## 📤 Output Format

For each query `(s, t)`:

- If there is **no path** from `s` to `t`:

```text
-1
```

- If there **is** a path:

```text
d
k a1 a2 ... ak
```

Where:

- `d` – shortest distance from `s` to `t`
- `k` – number of nodes in the path
- `a1, ..., ak` – nodes along one shortest path
- `a1 = s` and `ak = t`

---

## ⚙️ Implementation Details

File: `answer.go`

Key components:

- `type Edge` – represents a directed edge with a weight.
- `type Graph` – adjacency list representation of the graph.
- `dijkstra(g *Graph, src int)` – runs Dijkstra from a single source:
  - Returns:
    - `dist[]` – shortest distance from `src` to every node
    - `parent[]` – used to reconstruct paths
- `reconstructPath(parent []int, src, dst int)` – rebuilds the path from `src` to `dst` (if it exists).
- `FastScanner` – a buffered input reader for competitive-style fast I/O.
- `main()`:
  - Reads `n, m, q`
  - Builds the graph
  - Reads all queries
  - For each **distinct source `s`**, runs Dijkstra once and caches `(dist, parent)` in maps:
    - `distCache[src]`
    - `parentCache[src]`
  - Answers each query using the cached result.

This approach is educational and relatively efficient when there are multiple queries sharing the same source.

---

## ⏱️ Complexity

For each distinct source `s` that appears in the queries:

- Dijkstra’s algorithm runs in:

> **O(m log n)**

using a binary heap (`container/heap`) priority queue.

Let `S` be the number of **distinct** sources among all queries:

- Total complexity is roughly:

> **O(S · m log n + q · L)**

where `L` is the length of a printed path.

---

## ▶️ How to Build & Run

Make sure you have **Go** installed.

### Run directly:

```bash
go run answer.go < input.txt
```

### Build and run:

```bash
go build -o solution answer.go
./solution < input.txt
```

---

## 📝 Example (Conceptual)

Example input:

```text
4 5 2
1 2 3
2 3 4
1 3 10
3 4 2
2 4 8
1 4
3 1
```

- Query 1: `1 -> 4`
  - Shortest distance: `3 + 4 + 2 = 9`
  - One valid shortest path: `1 2 3 4`
- Query 2: `3 -> 1`
  - No path, so output: `-1`

Example output:

```text
9
4 1 2 3 4
-1
```

(Exact formatting may depend on number of queries and graph structure.)

---

## 📄 License

This code can be used under the **MIT License** (or your preferred license if you add a `LICENSE` file).

You are free to:

- Reuse the implementation in your own competitive programming templates
- Modify and extend it for additional features (e.g., undirected graphs, negative edges with Bellman–Ford, etc.)

---

Happy hacking with Go and Dijkstra! 🚀

### by soroosh morshedi - https://sorooshmorshedi.ir
