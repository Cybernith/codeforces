# Doors and Keys - Codeforces 1644A

**Source:** https://codeforces.com/problemset/problem/1644/A

---

## Problem Statement

The knight is standing in front of a long and narrow hallway. A princess is waiting at the end of it.

In the hallway there are three doors: a **red** door, a **green** door and a **blue** door. The doors are placed one after another in some order. To proceed to the next door, the knight must first open the previous one.

Each door can be opened only with a key of the corresponding color. So three keys — red, green and blue — are also placed somewhere in the hallway. To open a door, the knight must first pick up the key of its color.

The knight has a map of the hallway. It is represented as a string of **six characters** where:

- `'R'`, `'G'`, `'B'` denote **red, green and blue doors**,
- `'r'`, `'g'`, `'b'` denote **red, green and blue keys**.

Each of these six characters appears in the string **exactly once**.  
The knight starts at the **left** end of the string and walks to the right.

### Task

Given this map, determine if the knight can open **all three doors** and reach the princess at the end of the hallway.

### Input

- The first line contains a single integer `t` (`1 ≤ t ≤ 720`) — the number of test cases.
- Each of the following `t` lines contains a string of length `6` consisting of the characters `R`, `G`, `B`, `r`, `g`, `b`, each appearing exactly once.

### Output

For each test case, print:

- `YES` if the knight can open all doors and reach the end,
- `NO` otherwise.

---

## Solution Overview

The knight moves from **left to right** along the string. When he sees:

- a **key** (lowercase letter: `'r'`, `'g'`, `'b'`), he picks it up and keeps it;
- a **door** (uppercase letter: `'R'`, `'G'`, `'B'`), he can pass through it **only if** he already picked up the corresponding key.

The question becomes:

> While scanning the string from left to right, is it always true that when we encounter a door `'R'`, `'G'`, or `'B'`, we have already seen the corresponding key `'r'`, `'g'`, or `'b'`?

If this condition holds for **every door** in the string, then the answer is `YES`; otherwise, `NO`.

### Algorithm

For each hallway string `s` of length 6:

1. Create an empty set `keys_collected`.
2. Iterate through `s` from left to right:
   - If the current character `ch` is lowercase (`'r'`, `'g'`, `'b'`):
     - Add `ch` to `keys_collected`.
   - Else (the character is uppercase: `'R'`, `'G'`, `'B'`):
     - Convert `ch` to lowercase: `door_key = ch.lower()`.
     - If `door_key` is **not** in `keys_collected`, then the knight has reached a door without its key → return `NO` for this test case.
3. If we finish scanning the string without failing at any door, return `YES`.

This directly simulates the knight's movement and checks the condition for each door.

### Example

Consider the hallway string:

```text
rgbBRG
```

Step-by-step:

- `r`: pick up red key → `{r}`
- `g`: pick up green key → `{r, g}`
- `b`: pick up blue key → `{r, g, b}`
- `B`: blue door, we have `'b'` → pass
- `R`: red door, we have `'r'` → pass
- `G`: green door, we have `'g'` → pass

All doors can be opened → `YES`.

As another example:

```text
RgbrBG
```

- `R`: red door, but we **do not** have `'r'` yet → immediately `NO`.

### Complexity

Each hallway string has fixed length 6, so:

- Time complexity per test is `O(6)` which is effectively `O(1)`.
- Space complexity is `O(1)` — we store at most three keys in a small set.

Overall, for up to `t ≤ 720` test cases, the program easily runs within any reasonable limits.

---

## Implementation Notes

### Python (`awnser.py`)

- Reads `t` from standard input, then processes `t` hallway strings.
- Uses a `set` to store collected keys.
- For each character:
  - `str.islower()` is used to detect keys,
  - `str.lower()` is used to map doors to their corresponding key characters.
- Prints `YES` or `NO` for each test case.

### Go (`awnser.go`)

- Uses `bufio.NewReader` and `bufio.NewWriter` for efficient I/O.
- For each test case:
  - Reads the hallway string.
  - Tracks collected keys using a small boolean map/array.
  - Checks whether each encountered door has its key collected.
- Prints `YES` or `NO` accordingly.

---

## Notes

CyberNith - Soroosh Morshedi  
https://sorooshmorshedi.ir
