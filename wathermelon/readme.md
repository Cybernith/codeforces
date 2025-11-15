# Watermelon - Codeforces 4A

**Source:** https://codeforces.com/problemset/problem/4/A

---

## Problem Statement

One hot summer day Pete and his friend Billy decided to buy a watermelon. They chose the biggest and the ripest one, in their opinion. After that the watermelon was weighed, and the scales showed `w` kilos.

They rushed home, dying of thirst, and decided to divide the berry. However, they faced a hard problem.

Pete and Billy are great fans of **even numbers**, so they want to divide the watermelon in such a way that **each of the two parts weighs an even number of kilos**. It is *not* necessary that the parts are equal, but:

- Each part must be of **positive** weight.
- Each part must be an **even** integer number of kilos.

Your task is to determine if this is possible.

### Input

The first (and only) line of input contains a single integer `w` — the weight of the watermelon in kilos.

Constraints:

- `1 ≤ w ≤ 100`

### Output

Print `YES` if the boys can divide the watermelon into two parts, each of them weighing an even number of kilos.  
Print `NO` otherwise.

---

## Solution Overview

We need to decide if we can split `w` into two **positive even integers**:

```text
w = x + y
```

where `x > 0`, `y > 0`, and both `x` and `y` are even numbers.

Key observations:

1. The sum of two even numbers is always even.  
   Therefore, if `w` is **odd**, it is **impossible** to represent `w` as a sum of two even numbers → answer is `NO`.

2. If `w` is even, we still need each part to be **positive** and **even**:

   - The smallest positive even number is `2`.

   - So the smallest valid split would be:

     ```text
     w = 2 + 2 = 4
     ```

   - That means the **minimum** `w` for which a valid split exists is `4`.

Combining the conditions:

- If `w` is even **and** `w ≥ 4`, then the answer is `YES`.
- Otherwise, the answer is `NO`.

### Examples

1. `w = 2`

   - Even, but cannot be split into two positive even parts → `NO`.

2. `w = 4`

   - `4 = 2 + 2` → both parts are even and positive → `YES`.

3. `w = 5`

   - Odd → impossible to write as sum of two even numbers → `NO`.

4. `w = 8`

   - `8 = 2 + 6` or `8 = 4 + 4` → valid splits → `YES`.

### Algorithm

Given `w`:

1. If `w` is even **and** `w ≥ 4`, print `YES`.
2. Otherwise, print `NO`.

This is a direct `O(1)` check with no loops required.

### Complexity

- Time complexity: `O(1)` — a couple of comparisons and a modulus operation.
- Space complexity: `O(1)` — constant extra memory.

---

## Implementation Notes

### Python (`awnser.py`)

- Reads an integer `w` from standard input.
- Checks the condition `(w % 2 == 0 and w >= 4)`.
- Prints `YES` if the condition holds, otherwise prints `NO`.

The Python implementation uses clear naming (`weight`) and a small helper function for readability.

### Go (`awnser.go`)

- Reads an integer `w` using buffered input.
- Performs the same check `(w % 2 == 0 && w >= 4)`.
- Prints `YES` or `NO` using buffered output.

Both implementations are suitable for direct submission to Codeforces.

---

## Notes

CyberNith - Soroosh Morshedi  
https://sorooshmorshedi.ir
