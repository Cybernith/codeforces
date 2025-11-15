# H. Cycle Sort — Codeforces 2161H

## Problem Summary

You are given two integer arrays `a[0..n-1]` and `b[0..m-1]`.  
Each integer from `1` to `n + m` appears **exactly once** in the concatenation of both arrays.

You perform `k` operations. For every integer `i` from `0` to `k - 1` (in this order):

- Let `ia = i mod n`
- Let `ib = i mod m`
- If `a[ia] > b[ib]`, you **swap** `a[ia]` and `b[ib]`  
- Otherwise, you do nothing

Your task is to output the final state of both arrays after all `k` operations are completed, for each test case.

Constraints (per test case):

- `1 ≤ n, m ≤ 2 * 10^5`
- `0 ≤ k ≤ 10^18`
- `a` and `b` together form a permutation of `1..n+m`
- Sum of `n + m` over all test cases ≤ `2 * 10^5`

---

## Simple Simulation Approach (Educational Version)

> ⚠️ Note: This implementation is **meant as a clean, readable, educational solution**.  
> It directly simulates the `k` operations and is suitable for small or moderate values of `k`.  
> For the full Codeforces constraints (`k` up to 10^18) an optimized solution with deeper math / data structures is required.

### Idea

The operation definition is very direct:

- We maintain two indices `ia` and `ib` which represent the current positions in arrays `a` and `b`.
- At step `i`:
  - `ia = i % n`
  - `ib = i % m`
  - If `a[ia] > b[ib]`, we swap them.

This can be implemented literally as a loop from `0` to `k - 1`.

Because `ia` and `ib` are just cycling through indices with modulo arithmetic, we don’t need any complex data structures to simulate the process — just access-by-index and a conditional swap.

### Python API-style Solver

In Python we wrap the logic in a small, PEP 8–compliant class so it’s easy to reuse and unit-test:

- `CycleSortSolver.solve_single_case(n, m, k, a, b)` → returns the final `(a, b)`
- `main()` reads from standard input in Codeforces format and prints the result.

### Go Implementation

In Go we follow the same logic:

- Read `t` test cases
- For each case, read `n, m, k` and the arrays
- Simulate the `k` operations in a simple loop
- Print the final arrays

Each language keeps the code structured, readable, and easy to adapt later if you want to plug in a more advanced, optimized core instead of pure simulation.

---

## Complexity (for this educational version)

For each test case:

- Time: **O(k)** operations  
- Memory: **O(n + m)**

This is perfectly fine if `k` is not huge in your tests (for example, `k` up to ~10^6 or so), and is great for understanding and debugging the process. For the full problem constraints one would need to exploit periodic structure and value movements much more aggressively.

---

## Author / Credits

Educational implementation skeleton & README structure by:

**Cybernith / Soroosh Morshedi**  
https://sorooshmorshedi.ir
