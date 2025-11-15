# Hulk - Codeforces 705A

**Source:** https://codeforces.com/problemset/problem/705/A

---

## Problem Statement

Dr. Bruce Banner hates his enemies (like others don't). When he turns into the incredible Hulk, he can barely talk, so he needs help expressing his feelings.

Hulk's feelings have `n` layers:

- The 1st layer is **hate**,
- The 2nd layer is **love**,
- The 3rd layer is **hate**,
- The 4th layer is **love**, and so on — alternating between hate and love.

We must output a sentence that describes these feelings in sequence.

Formally:

- If `n = 1`, the sentence is:

  ```text
  I hate it
  ```

- If `n = 2`, the sentence is:

  ```text
  I hate that I love it
  ```

- If `n = 3`, the sentence is:

  ```text
  I hate that I love that I hate it
  ```

- And so on for larger `n`.

### Input

The input consists of a single integer `n` (`1 ≤ n ≤ 100`) — the number of layers of love and hate.

### Output

Print Hulk's feeling in one line, following the pattern described above.

---

## Solution Overview

We need to construct a sentence with `n` segments, where each segment is either:

- `"I hate"` for odd positions (1st, 3rd, 5th, ...),
- `"I love"` for even positions (2nd, 4th, 6th, ...).

Between consecutive segments, we insert the word `" that "`.

At the very end of the sentence, we append `" it"`.

So for general `n`, the sentence has the structure:

```text
I hate [that I love] [that I hate] ... it
```

where `[ ... ]` represents repeating `"that I love"` or `"that I hate"` segments appropriately.

### Algorithm

1. Read integer `n`.
2. Initialize an empty list `parts`.
3. For `i` from `1` to `n`:
   - If `i` is odd, append `"I hate"` to `parts`.
   - If `i` is even, append `"I love"` to `parts`.
4. Join all elements in `parts` using `" that "` as the separator.
5. Append `" it"` to the end of this joined string.
6. Print the final string.

This directly follows the problem statement and guarantees the correct format.

### Example

For `n = 3`:

- `i = 1` → `"I hate"`
- `i = 2` → `"I love"`
- `i = 3` → `"I hate"`

`parts = ["I hate", "I love", "I hate"]`

Join with `" that "`:

```text
"I hate that I love that I hate"
```

Append `" it"`:

```text
"I hate that I love that I hate it"
```

which matches the example.

### Complexity

- Time complexity: `O(n)` — we build a list of `n` segments and join them.
- Space complexity: `O(n)` — we store `n` short strings in a list.

Given `n ≤ 100`, this is trivially efficient.

---

## Implementation Notes

### Python (`awnser.py`)

- Reads `n` using standard input.
- Uses a loop to build the list of segments (`"I hate"` or `"I love"`).
- Joins them with `" that "` and appends `" it"`.
- Prints the resulting sentence.

### Go (`awnser.go`)

- Reads `n` using a buffered reader.
- Builds the sentence in a slice of strings.
- Joins them manually with `" that "` and appends `" it"`.
- Prints the result using a buffered writer.

---

## Notes

CyberNith — Soroosh Morshedi  
https://sorooshmorshedi.ir
