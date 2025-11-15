# Payment Without Change - Codeforces 1256A

**Source:** https://codeforces.com/problemset/problem/1256/A

---

## Problem Statement

You have `a` coins of value `n` and `b` coins of value `1`. You always pay in exact change, so you want to know if there exist such `x` and `y` that if you take `x` (`0 ≤ x ≤ a`) coins of value `n` and `y` (`0 ≤ y ≤ b`) coins of value `1`, then the total value of taken coins will be `S`.

You have to answer `q` independent test cases.

### Input

The first line of the input contains one integer `q` (`1 ≤ q ≤ 10^4`) — the number of test cases. Then `q` test cases follow.

The only line of each test case contains four integers `a`, `b`, `n` and `S` (`1 ≤ a, b, n, S ≤ 10^9`) — the number of coins of value `n`, the number of coins of value `1`, the value `n` and the required total value `S`.

### Output

For the `i`-th test case print the answer for it — `YES` if there exist such `x` and `y` that if you take `x` coins of value `n` and `y` coins of value `1`, then the total value of taken coins will be `S`, and `NO` otherwise.

You may print every letter in any case you want (so, for example, the strings `yEs`, `yes`, `Yes` and `YES` will all be recognized as a positive answer).

---

## Solution Overview

We need to decide, for each test case, whether it is possible to represent `S` in the form

```text
S = x * n + y
```

subject to the constraints:

- `0 ≤ x ≤ a` (we have at most `a` coins of value `n`)
- `0 ≤ y ≤ b` (we have at most `b` coins of value `1`)
- `y` is an integer

Because the coins of value `1` can cover **any** remainder up to `b`, the natural greedy idea is:

1. Use as many coins of value `n` as possible, but no more than allowed:

   ```text
   x_max = min(a, S // n)
   ```

2. After taking `x_max` coins, the remaining value we need is:

   ```text
   remaining = S - x_max * n
   ```

3. If `remaining ≤ b`, then we can pay the rest with coins of value `1`, so the answer is `YES`.  
   Otherwise, it is impossible with this choice, and it turns out no smaller `x` will help either.

### Why Greedy Works

Suppose we take fewer coins of value `n` than `x_max`. That means we replace some `n`-value coins with `1`-value coins, which **increases** the required number of `1`-value coins by a multiple of `n`:

- If we decrease `x` by 1, the remaining value increases by `n`.
- So if `remaining` with `x_max` coins is already greater than `b`, it only gets larger (and thus worse) when using fewer `n`-value coins.

Therefore, checking just `x = min(a, S // n)` is sufficient.

### Algorithm

For each test case:

1. Read `a`, `b`, `n`, `S`.

2. Compute:

   ```text
   x_max = min(a, S // n)
   remaining = S - x_max * n
   ```

3. If `remaining <= b`, print `YES`, otherwise print `NO`.

### Complexity

For each test case, the computation is `O(1)`:

- A few integer operations and comparisons.

Given up to `10^4` test cases, the total complexity is `O(q)`, which is easily fast enough.

---

## Implementation Notes

### Python (`awnser.py`)

- Uses fast input via `sys.stdin`.
- Processes all `q` test cases in a loop.
- For each test, applies the greedy logic described above and prints `YES` or `NO`.

The file is named:

```text
awnser.py
```

and is suitable for direct submission to Codeforces (Python implementation).

### Go (`awnser.go`)

- Uses `bufio.NewReader` and `bufio.NewWriter` for efficient input/output.
- Reads `q`, then processes each test case in a loop.
- For each test, computes `x_max` and `remaining`, and prints `YES` or `NO` accordingly.

The file is named:

```text
awnser.go
```

and is suitable for direct submission to Codeforces (Go implementation).

---

## Notes

CyberNith — Soroosh Morshedi  
https://sorooshmorshedi.ir

