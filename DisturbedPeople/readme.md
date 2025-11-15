# Disturbed People

https://codeforces.com/problemset/problem/1077/B

There is a house with nn flats situated on the main street of Berlatov. Vova is watching this house every night. The house can be represented as an array of nn integer numbers a1,a2,…,ana1,a2,…,an, where ai=1ai=1 if in the ii-th flat the light is on and ai=0ai=0 otherwise.

Vova thinks that people in the ii-th flats are disturbed and cannot sleep if and only if 1<i<n1<i<n and ai−1=ai+1=1ai−1=ai+1=1 and ai=0ai=0.

Vova is concerned by the following question: what is the minimum number **k** such that if people from exactly kk pairwise distinct flats will turn off the lights then nobody will be disturbed? Your task is to find this number **k**.

### Input

The first line of the input contains one integer nn (3≤n≤1003≤n≤100) — the number of flats in the house.

The second line of the input contains nn integers a1,a2,…,ana1,a2,…,an (ai∈{0,1}ai∈{0,1}), where aiai is the state of light in the ii-th flat.

### Output


Print only one integer — the minimum number **k** such that if people from exactly **k** pairwise distinct flats will turn off the light then nobody will be disturbed.


__________________________________________________________________________

# Disturbed People (Codeforces 1077B)

This folder contains a solution for the **"Disturbed People"** problem from Codeforces (problem 1077B).  
In the overall repository structure, each folder corresponds to one Codeforces task and its answer.

## Problem Overview

We have a building with `n` flats in a row. For each flat `i`:

- `a[i] = 1` if the light is ON
- `a[i] = 0` if the light is OFF

A person in flat `i` (with `1 < i < n`) is considered **disturbed** if:

- `a[i - 1] == 1`
- `a[i] == 0`
- `a[i + 1] == 1`

We are allowed to turn OFF the light in some flats (change `1` to `0`).  
We want the **minimum number** of such operations so that **no one is disturbed**.

### Key Idea (Greedy)

Scan the array from left to right (from index `1` to `n - 2`, 0-based).  
Whenever we see a pattern `[1, 0, 1]` centered at `i`, we:

- Turn OFF the light at position `i + 1` (the right neighbor).
- Increase the operation counter by 1.

This greedy approach is optimal because:

- Each disturbance is local to a triple `(i - 1, i, i + 1)`.
- By always modifying the right neighbor, we prevent overlapping disturbances
  from being counted more than once.

---

## Files

- `disturbed_people.py` — Python solution (OOP, PEP 8, ready for competitive or reuse)
- `main.go` — Go solution (idiomatic Go, competitive-style I/O)
- `LICENSE` — MIT License
- `README.md` — This documentation

---

## Python Solution (`disturbed_people.py`)

### Design

The main class is:

- `DisturbedPeopleSolver`
  - Attributes:
    - `flats: List[int]` — the current state of lights in the building.
  - Methods:
    - `minimum_switch_offs() -> int` — returns the minimal number of operations required.

The implementation works on a copy of the `flats` list to avoid in-place side effects.

### Usage (Command Line)

Input format:

```text
n
a1 a2 ... an
```

Run:

```bash
echo -e "10\n1 1 0 1 1 0 1 0 1 0" | python disturbed_people.py
```

Output:

```text
2
```

### Usage (As a Module)

```python
from disturbed_people import DisturbedPeopleSolver

flats = [1, 1, 0, 1, 1, 0, 1, 0, 1, 0]
solver = DisturbedPeopleSolver(flats=flats)
print(solver.minimum_switch_offs())  # 2
```

---

## Go Solution (`main.go`)

The Go version mirrors the Python logic, using:

- `DisturbedPeopleSolver` struct with field:
  - `Flats []int`
- Method:
  - `MinimumSwitchOffs() int`

### Usage (Command Line)

```bash
echo -e "10\n1 1 0 1 1 0 1 0 1 0" | go run main.go
```

Output:

```text
2
```

The Go solution uses `bufio` for efficient input/output,
which is standard practice in competitive programming.

---

## Complexity

For both Python and Go solutions:

- Time complexity: **O(n)** — a single pass through the array.
- Space complexity: **O(n)** — due to copying the input state (could be reduced to O(1) if in-place modification is allowed).

---

## License

This folder (and its contents) is licensed under the **MIT License**.  
See the `LICENSE` file for the full license text.

** by soroosh morshedi **
