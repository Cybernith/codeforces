# Theatre Square — Codeforces 1A

**Source:** https://codeforces.com/problemset/problem/1/A

---

## Problem Statement

Theatre Square in the capital city of Berland has a rectangular shape with size `n × m` meters. On the occasion of the city's anniversary, a decision was taken to pave the Square with **square granite flagstones**. Each flagstone is of size `a × a` meters.

You must determine the **least number of flagstones** needed to pave the Square. It is allowed to cover a surface **larger** than the Theatre Square, but the Square has to be fully covered. It is **not** allowed to break flagstones. The sides of the flagstones should be parallel to the sides of the Square.

### Input

The input contains three positive integers in the first line:

```text
n m a
```

with constraints:

- `1 ≤ n, m, a ≤ 10^9`

### Output

Print a single integer — the **minimum number of flagstones** needed to cover the Theatre Square.

---

## Solution Overview

We need to cover an `n × m` rectangle using `a × a` square tiles, without breaking tiles, and we may go beyond the rectangle's border if needed.

The key observation is that along each dimension we need to know **how many tiles** are required:

- Along the `n` side (height): we need `ceil(n / a)` tiles.
- Along the `m` side (width): we need `ceil(m / a)` tiles.

Then, because the tiles form a grid:

```text
total_tiles = ceil(n / a) * ceil(m / a)
```

### Integer Ceiling Without Floating Point

We should avoid floating-point arithmetic because `n`, `m`, and `a` can be as large as `10^9`, and their division might introduce precision issues in floating point.

A standard integer trick to compute `ceil(x / y)` (for positive integers `x` and `y`) is:

```text
ceil(x / y) = (x + y - 1) / y   (integer division)
```

So we compute:

```text
tiles_n = (n + a - 1) // a
tiles_m = (m + a - 1) // a
answer  = tiles_n * tiles_m
```

This works safely within 64-bit integer range since the maximum result is at most:

```text
ceil(10^9 / 1) * ceil(10^9 / 1) = 10^18
```

which fits in a signed 64-bit integer.

### Complexity

- Time complexity: `O(1)` — only a few arithmetic operations.
- Space complexity: `O(1)` — constant extra memory.

---

## Implementation Notes

### Python (`awnser.py`)

- Reads `n`, `m`, `a` from standard input.
- Computes `tiles_n = (n + a - 1) // a` and `tiles_m = (m + a - 1) // a`.
- Prints `tiles_n * tiles_m` as the final answer.
- Uses Python’s built-in arbitrary-precision integers, so overflow is not a concern.

### Go (`awnser.go`)

- Uses `int64` to safely hold values up to `10^18`.

- Reads `n`, `m`, `a` as `int64` using buffered input.

- Computes the same integer ceiling logic:

  ```text
  tilesN := (n + a - 1) / a
  tilesM := (m + a - 1) / a
  answer := tilesN * tilesM
  ```

- Writes the result using buffered output.

---

## Notes

CyberNith — Soroosh Morshedi  
https://sorooshmorshedi.ir
