# Algorithms & Competitive Programming Playground

This repository is a personal playground for algorithms, data structures, and competitive programming solutions across multiple languages (currently **Python**, **Go**, and **C++**).

It’s designed as a reference and learning space rather than a polished library: each file focuses on one concept, problem, or technique, with clear names and minimal boilerplate so it’s easy to read and reuse.

---

## 🚀 Goals

- Practice implementing classic algorithms and data structures
- Collect solutions to online judge problems (e.g. Codeforces-style tasks)
- Compare idiomatic implementations across different languages
- Build a small, practical reference for future projects and interviews

---

## 📂 Repository Structure

A suggested structure (you can adapt it to your actual folders):

```text
.
├── python/
│   ├── graphs/
│   ├── dp/
│   ├── searching/
│   └── misc/
├── cpp/
│   ├── graphs/
│   ├── dp/
│   ├── data_structures/
│   └── misc/
├── go/
│   ├── graphs/
│   ├── dp/
│   └── misc/
├── tests/
│   └── sample_inputs/
└── README.md
```

Typical patterns:

- `graphs/` – BFS, DFS, Dijkstra, Floyd–Warshall, topological sort, etc.
- `dp/` – classic DP problems (knapsack, LIS, grid paths, etc.)
- `data_structures/` – segment tree, Fenwick tree, disjoint-set (DSU), etc.
- `searching/` – binary search, ternary search, custom search patterns
- `misc/` – smaller utilities, math helpers, or problem-specific code

> Note: The exact folders and filenames may differ; this README is meant as a clean, English entry point for the repo.

---

## 🧠 Included Concepts (Examples)

Some of the topics covered (or planned) in this repo:

- Graph algorithms:
  - Shortest paths (Dijkstra, BFS on unweighted graphs)
  - Graph traversal (DFS/BFS)
  - Topological sorting
- Dynamic programming:
  - 1D & 2D DP states
  - Classic problems from contests
- Searching & sorting:
  - Binary search patterns
  - Custom comparators and stable sorting
- Data structures:
  - Union-Find (DSU)
  - Segment tree / Fenwick tree (BIT)
- Problem-specific solutions:
  - Contest-style problems with input/output format
  - Small utilities for parsing and testing

---

## 🔧 How to Run

Each language directory is self-contained.

### Python

```bash
cd python
python some_script.py < input.txt
```

### Go

```bash
cd go
go run answer.go < input.txt
```

Or for compiled binaries:

```bash
go build -o solution answer.go
./solution < input.txt
```

### C++

```bash
cd cpp
g++ -std=c++17 -O2 main.cpp -o main
./main < input.txt
```

Adjust filenames as needed according to your actual structure.

---

## 🧪 Testing

You can keep example inputs in a `tests/` or `sample_inputs/` directory and run:

```bash
cat tests/sample1.in | python python/graphs/dijkstra.py
cat tests/sample1.in | go run go/graphs/answer.go
cat tests/sample1.in | ./cpp/main
```

This keeps a simple workflow that is easy to reproduce for competitive programming sessions.

---

## 🧾 License

This repository is licensed under the **MIT License**.

You are free to:

- Use the code in your own projects
- Modify and adapt snippets
- Learn from the examples and structure

Just keep the license file if you reuse substantial parts of the repository.

---

## 🤝 Contributing / Personal Use

This repo is mainly for personal use and learning, but you can:

- Fork it and adapt the structure to your own style
- Add more algorithms, optimizations, and comments
- Use it as a template for your own “algorithms playground”

If you ever open this publicly, feel free to add issues or pull requests to track ideas, missing algorithms, or improvements.

---

Happy coding and have fun exploring algorithms! 🧩
