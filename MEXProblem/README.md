#   Yet Another MEX Problem - Codeforces 2146E

**Problem:** E. Yet Another MEX Problem  
**Platform:** Codeforces  
**Contest:** Codeforces Round 1052 (Div. 2)  
**Problem ID:** 2146E  
**Link:** https://codeforces.com/contest/2146/problem/E  

---

## Problem Statement (Summary)

You are given an array `a` of length `n` with elements `0 ≤ a[i] ≤ n`.

For any subarray `b = a[l..r]`:

- Let `mex(b)` be the smallest non-negative integer that does **not** occur in `b`.
- Define the **weight** of `b` as the number of indices `i` in `[l, r]` such that `b[i] > mex(b)`.

Formally:
- `weight(b) = |{ i ∈ [l, r] : b[i] > mex(b) }|`.

For every position `i` (1 ≤ i ≤ n), you must compute:

\[
\text{ans}[i] = \max_{1 \le l \le i} w(l, i)
\]

That is, over **all subarrays ending at i**, find the maximum possible weight.

Multiple test cases are given. The sum of `n` over all test cases does not exceed `3 * 10^5`.

---

## Solution Idea

### 1. Rewriting the Weight via a Parameter `x`

For a subarray `b` and an integer `x` that **does not appear** in `b`, define:

- `g(b, x) = number of elements in b that are > x`.

If we choose `x = mex(b)`, then:

- `weight(b) = g(b, mex(b))`.

Furthermore, among all integers `x` that are absent from `b`, `x = mex(b)` maximizes `g(b, x)`.  
So for each subarray `b`, the best possible weight is exactly `g(b, mex(b))`.

---

### 2. Fixing the Right Endpoint

Fix a right endpoint `r`. We want:

\[
\max_{1 \le l \le r} w(l, r)
\]

Instead of enumerating subarrays `(l, r)`, we enumerate **integers** `x` and consider subarrays ending at `r` that **do not contain `x`**.

Let `last[x]` be the last position where value `x` appeared before or at `r` (or `0` if `x` never appeared).

- Any subarray `a[l..r]` that does not contain `x` must satisfy `l > last[x]`.
- Among such subarrays, the **longest** one is `a[last[x] + 1 .. r]`, and that one maximizes `g(a[l..r], x)`.

Define:

\[
value_x(r) = g(a[last[x] + 1 .. r], x)
\]

Then, for a fixed `r`, **every** subarray weight is ≤ `max_x value_x(r)`.  
Moreover, this upper bound is tight: for some `x = mex(b)` and its maximal subarray starting from `last[x]+1`, the value is achieved.

Therefore:

\[
\max_{1 \le l \le r} w(l, r) = \max_{x} value_x(r)
\]

So the problem becomes: while `r` moves from `1` to `n`, we need to maintain `value_x(r)` for all `x ∈ [0, n]` and query the global maximum.

---

### 3. Updating When We Extend to the Right

We process `a` from left to right. Suppose we are at position `r` and read `v = a[r]`.

**We maintain:** for each `x` in `[0, n]`

- a conceptual subarray `[last[x] + 1 .. r]`,
- and `value[x] = g(a[last[x] + 1 .. r], x)`.

Now we add one more element `v` at the right:

1. For all `x < v`:

   - The new element `v` is `> x`, and it lies inside the subarray `[last[x] + 1 .. r]`.
   - So `g` increases by `1`:

     \[
     value[x] \mathrel{+}= 1 \quad \text{for all } x \in [0, v-1].
     \]

2. For `x = v`:

   - We update `last[v] = r`.
   - The active subarray for `x = v` becomes `[r + 1 .. r]` (empty), so:

     \[
     value[v] = 0.
     \]

3. For all `x > v`:

   - The new element `v` is not greater than `x`, so `g` does not change.
   - Hence `value[x]` stays the same.

So each step involves:

- **Range add** `+1` on `x ∈ [0, v-1]`.
- **Point set** `value[v] = 0`.

The answer for position `r` is:

\[
\text{ans}[r] = \max_x value[x]
\]

---

### 4. Data Structure

We need a structure over the domain `x ∈ [0, n]` that supports:

1. Range add on `[L, R]`.
2. Point assignment at a single position `p`.
3. Query global maximum over `[0, n]`.

A classic **segment tree with lazy propagation** is ideal:

- Each node stores the maximum in its segment.
- Lazy node stores the pending “add” value.
- `point set` overrides that cell and clears its lazy value.
- `range add` propagates down the lazy increments.
- `max` is the root value.

For each element `a[r]`:

- `point set(a[r], 0)`
- if `a[r] > 0`, `range add(0, a[r] - 1, +1)`
- answer for this `r` = `tree.max()`.

---

### 5. Complexity

Let `N` be the sum of `n` over all test cases.

- Every position `r` causes:
  - one point assignment,
  - one range add,
  - one global max query.

Each operation costs `O(log n)`.  
Therefore:

- **Time complexity:** `O(N log N)`
- **Memory complexity:** `O(N)`

This fits comfortably in the limits (`N ≤ 3 * 10^5`).

---

## Implementations
The solution is implemented in three different languages.
File names are as requested:

پیاده‌سازی در سه زبان مختلف است.  
نام فایل‌ها طبق خواسته:

- 🐍 Python: `awnser.py`
- 🟦 Go: `awnser.go`
- 🧊 C++: `awnser.cpp`

All three files implement exactly the same algorithm (segment tree with lazy propagation).
---

## How to Run

### C++ (awnser.cpp)

```bash
g++ -std=c++17 -O2 -pipe -static -s awnser.cpp -o awnser_cpp
./awnser_cpp < input.txt > output.txt
```

### Go (awnser.go)

```bash
go run awnser.go < input.txt > output.txt
```

### Python (awnser.py)

```bash
python3 awnser.py < input.txt > output.txt
```
- `All programs read input from stdin and write output to stdout, fully compatible with the Codeforces format.`.

## Author
This code and documentation are prepared for personal and educational use and for integration into competitive programming setups
- Author: `soroosh morshedi`
- Website: https://sorooshmorshedi.ir

