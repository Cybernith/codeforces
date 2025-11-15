# Spy Detected! — Codeforces 1512A

**Source:** https://codeforces.com/problemset/problem/1512/A

---

## Problem Statement

You are given an array `a` consisting of `n` (`n ≥ 3`) positive integers. It is known that in this array, **all the numbers except one are the same**.

For example, in the array `[4, 11, 4, 4]` all numbers except one are equal to `4`.

Your task is to find the **index** (1-based) of the element that does **not** equal the others.

### Input

The first line contains a single integer `t` (`1 ≤ t ≤ 100`) — the number of test cases.

Then `t` test cases follow.

For each test case:

- The first line contains a single integer `n` (`3 ≤ n ≤ 100`) — the length of the array `a`.
- The second line contains `n` integers `a1, a2, …, an` (`1 ≤ ai ≤ 100`).  
  It is guaranteed that all the numbers except one in the array `a` are the same.

### Output

For each test case, output a single integer — the **1-based index** of the element that is different from the others.

---

## Solution Overview

We know that in each array:

- Exactly **one element** is different.
- All the other `n - 1` elements are equal to some common value.

This structure allows a very simple and efficient detection.

### Key Observation

Look at the **first three elements** of the array: `a[0]`, `a[1]`, `a[2]`.

Because there is only one unique element, among these three:

- Either the unique value appears in these first three positions, or  
- The common value appears at least two times among them.

From this, we can determine the **common value** quickly:

1. If `a[0] == a[1]`, then the common value is `a[0]` (and the unique one is somewhere else).
2. Else if `a[0] == a[2]`, then the common value is `a[0]`, and `a[1]` is the unique one.
3. Else, `a[1] == a[2]`, so the common value is `a[1]`, and `a[0]` is the unique one.

Once we know the common value, we scan the array to find the element that is **not equal** to it and output its index (1-based).

### Algorithm

For each test case:

1. Read `n` and the array `a`.
2. Use the first three elements to detect the common value:
   - If `a[0] == a[1]`:
     - `common = a[0]`
   - Else if `a[0] == a[2]`:
     - The unique index is `2` (1-based), so we can immediately output `2`.
   - Else:
     - The unique index is `1` (1-based), so we can immediately output `1`.
3. If we already found the unique index in step 2, we are done for this test case.
4. Otherwise, we know the common value. Loop over the array from index `0` to `n - 1`:
   - Find the index `i` where `a[i] != common` and output `i + 1`.
   - There will be exactly one such index by problem guarantee.

### Complexity

For each test case:

- Determining the common value from the first three elements is `O(1)`.
- Scanning the array to find the unique element is `O(n)`.
- So overall complexity per test case is `O(n)`.

Given `t ≤ 100` and `n ≤ 100`, this is easily efficient.

---

## Implementation Notes

### Python (`awnser.py`)

- Uses standard input (`sys.stdin`) for fast reading.
- For each test case:
  - Reads `n` and the list `a`.
  - Handles the logic described above to find and print the unique element’s 1-based index.
- Designed for direct submission in a Codeforces environment.

### Go (`awnser.go`)

- Uses `bufio.NewReader` and `bufio.NewWriter` for efficient I/O.
- For each test case:
  - Reads `n` and the array `a` as integers.
  - Applies the same logic as the Python implementation.
  - Prints the 1-based index of the unique element.

---

## Notes

CyberNith — Soroosh Morshedi  
https://sorooshmorshedi.ir
