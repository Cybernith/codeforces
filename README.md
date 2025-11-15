# Codeforces Solutions Collection

This repository contains a growing collection of **Codeforces problem solutions**, each organized in its own folder.

Every problem folder is self–contained and includes:

- `awnser.py` — Python solution (PEP 8–compliant, with clear function and variable names)
- `awnser.go` — Go solution (idiomatic Go, using buffered I/O)
- `awnser.cpp` — c++ implementation
- `README.md` — problem statement summary + explanation of the solution + notes

The structure is designed so you can quickly:

- read the problem summary,
- inspect or run the Python / Go solution,
- and extend the repo with more problems over time.

---

## Repository Structure

A typical layout looks like:

```text
.
├── <problem-folder-1>/
│   ├── awnser.py
│   ├── awnser.go
│   └── README.md
├── <problem-folder-2>/
│   ├── awnser.py
│   ├── awnser.go
│   └── README.md
├── <problem-folder-3>/
│   ├── awnser.py
│   ├── awnser.go
│   └── README.md
├── LICENSE
└── README.md   (this file)
```

Each `<problem-folder>` usually corresponds to one Codeforces task, for example:

- `way_too_long_words/` — Codeforces 71A
- `watermelon/` — Codeforces 4A
- `hulk/` — Codeforces 705A
- `team/` — Codeforces 231A
- `doors_and_keys/` — Codeforces 1644A
- `spy_detected/` — Codeforces 1512A
- `payment_without_change/` — Codeforces 1256A
- `mammoths_genome/` — Codeforces 747B
- `disturbed_people/` — Codeforces 1077B
- `sleuth/` — Codeforces  (Sleuth problem)
- `theatre_square/` — Codeforces 1A
- `cycle_sort/` or similar — Codeforces 2161H (Cycle Sort educational implementation)

> Note: Folder names are flexible; they just need to stay consistent inside the repo.

Each problem folder is fully documented locally inside its own `README.md`.

---

## Running the Solutions

### Python

From a problem folder (for example `watermelon/`):

```bash
cd watermelon
python awnser.py < input.txt
```

or, for interactive use:

```bash
python awnser.py
```

Most solutions read from standard input exactly as required by Codeforces.

### Go

From the same folder:

```bash
cd watermelon
go run awnser.go < input.txt
```

Or build a binary:

```bash
go build -o solution awnser.go
./solution < input.txt
```

---

## Adding New Problems

To add a new Codeforces problem to this repo, follow the same pattern:

1. Create a new folder, e.g. `new_problem_name/`.
2. Add:
   - `awnser.py` — Python implementation
   - `awnser.go` — Go implementation
   - `awnser.cpp` — c++ implementation
   - `README.md` — problem summary, explanation, and notes
3. Optionally update this top-level `README.md` and list the new folder in the structure section.

This layout makes it easy to keep Python and Go implementations in sync, and to use the repo as a personal Codeforces notebook.

---

## License

All code in this repository is available under the **MIT License**.

See the [`LICENSE`](./LICENSE) file for full details.

---

## Author / Credits

Maintained by **Cybernith / Soroosh Morshedi**  
https://sorooshmorshedi.ir
