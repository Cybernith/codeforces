# Disturbed People — Codeforces 1077B

**Source:** https://codeforces.com/problemset/problem/1077/B

---

## Problem Statement

There is a house with `n` flats situated on the main street of Berlatov. Vova is watching this house every night. The house can be represented as an array of `n` integer numbers `a1, a2, …, an`, where `ai = 1` if in the `i`-th flat the light is on and `ai = 0` otherwise.

Vova thinks that people in the `i`-th flats are **disturbed** and cannot sleep if and only if:

- `1 < i < n`, and  
- `ai−1 = ai+1 = 1`, and  
- `ai = 0`.

Vova is concerned by the following question: what is the minimum number `k` such that if people from exactly `k` pairwise distinct flats will turn off the lights then **nobody will be disturbed**? Your task is to find this number `k`.

### Input

The first line of the input contains one integer `n` (`3 ≤ n ≤ 100`) — the number of flats in the house.

The second line of the input contains `n` integers `a1, a2, …, an` (`ai ∈ {0, 1}`), where `ai` is the state of light in the `i`-th flat.

### Output

Print only one integer — the **minimum number `k`** such that if people from exactly `k` pairwise distinct flats turn off the light, then nobody will be disturbed.

---

## Solution Overview

We are given a binary array representing the lights in each flat. A person in flat `i` is **disturbed** if and only if the triple around `i` looks like:

```text
1 0 1
```

centered at `i` (i.e. `a[i-1] = 1`, `a[i] = 0`, `a[i+1] = 1`, with `1 < i < n`).

We are allowed to turn **some lights from 1 to 0**, and we want to **minimize** the number of such operations so that **no triple `1 0 1` remains**.

### Greedy Strategy

A simple and optimal greedy algorithm is:

1. Iterate `i` from `1` to `n - 2` (0-based indices, so this corresponds to `2` to `n-1` in 1-based).

2. If we see the pattern:

   ```text
   a[i - 1] = 1
   a[i]     = 0
   a[i + 1] = 1
   ```

   then we **turn off** the light in flat `i + 1`:

   ```text
   a[i + 1] = 0
   ```

3. Increase the counter `k` by 1 each time we do this.

4. Continue scanning to the right.

This works because:

- Any disturbed person is caused by a local `1 0 1` triple.
- Changing `a[i + 1]` from `1` to `0` both fixes the disturbance at `i` and avoids creating a new `1 0 1` starting at the next position.
- By scanning from left to right and always fixing the rightmost `1` in the pattern, we ensure we never count or fix the same disturbance twice.

### Complexity

- **Time:** `O(n)` — we visit each position at most once.  
- **Space:** `O(1)` extra if we are allowed to modify the array in-place (or `O(n)` if we work on a copy).

---

## Implementation Notes

### Python (`awnser.py`)

- Reads `n` and the array `a` from standard input.
- Implements the greedy scan from left to right.
- Counts how many times it breaks the pattern `1 0 1` by setting the rightmost `1` to `0`.
- Prints the final count `k` as the answer.

The Python file is named:

```text
awnser.py
```

and is designed to be used in a Codeforces-style environment (simple stdin/stdout, no external dependencies).

### Go (`awnser.go`)

- Mirrors the same greedy logic in Go.
- Uses `bufio.Reader`/`Writer` for efficient input and output.
- Also reads `n` and the array of `0/1`, applies the greedy algorithm, and prints `k`.

The Go file is named:

```text
awnser.go
```

and is suitable for submission to online judges that support Go.

---

## Notes

(Reserved for additional comments, links, or branding to be filled by Cybernith - soroosh morshedi)
https://sorooshmorshedi.ir

