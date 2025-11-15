# Way Too Long Words - Codeforces 71A

**Source:** https://codeforces.com/problemset/problem/71/A

---

## Problem Statement

Sometimes some words like `localization` or `internationalization` are so long that writing them many times in one text is quite tiresome.

We consider a word **too long** if its length is **strictly more than 10** characters. All too long words should be replaced with a special abbreviation.

The abbreviation is built as follows:

- Take the **first letter** of the word,
- Take the **last letter** of the word,
- Between them, write the **number of letters between** the first and last letters (in decimal, with no leading zeroes).

For example:

- `localization` → `l10n`
- `internationalization` → `i18n`

Words that are **not too long** (length ≤ 10) should **not** be changed.

### Input

- The first line contains an integer `n` (`1 ≤ n ≤ 100`) — the number of words.
- Each of the following `n` lines contains exactly one word.
- Each word:
  - consists of lowercase Latin letters, and
  - has a length from `1` to `100` characters.

### Output

Print `n` lines.  
The *i*-th line should contain the transformed version of the *i*-th word:

- If it is longer than 10 characters, its abbreviation,
- Otherwise, the word itself unchanged.

---

## Solution Overview

For each word, we decide whether it is "too long" based on its length:

- If `len(word) <= 10`, we simply print the word as-is.

- If `len(word) > 10`, we replace it with:

  ```text
  word[0] + str(len(word) - 2) + word[-1]
  ```

  where:

  - `word[0]` is the first character,
  - `word[-1]` is the last character,
  - `len(word) - 2` is the number of characters between them.

⚠️ Note: The problem statement says "the number of letters between the first and the last letters", which is `len(word) - 2`, not `len(word) - 1`.

### Example

Input:

```text
4
word
localization
internationalization
cat
```

Processing:

- `word` → length 4 → not too long → stays `word`
- `localization` → length 12 → `l` + `10` + `n` → `l10n`
- `internationalization` → length 20 → `i` + `18` + `n` → `i18n`
- `cat` → length 3 → not too long → stays `cat`

Output:

```text
word
l10n
i18n
cat
```

### Complexity

Let `L` be the length of a word and `n` be the number of words.

- For each word, we compute its length and possibly build a short string: `O(L)`.
- Summed over all words, the total complexity is `O(total input size)`.
- With the given constraints (`n ≤ 100`, `L ≤ 100`), this is trivial in terms of performance.

Space usage is `O(1)` extra besides storing a single word at a time (if implemented in a streaming fashion).

---

## Implementation Notes

### Python (`awnser.py`)

- Reads the integer `n`.
- Then reads `n` words, one per line.
- For each word, checks its length and prints either:
  - the original word, or
  - the abbreviation `first + count + last`.

### Go (`awnser.go`)

- Uses a buffered reader to read input efficiently.
- Reads `n`, then reads `n` words.
- For each word, applies the same logic as the Python implementation and prints the result line by line using a buffered writer.

---

## Notes

CyberNith — Soroosh Morshedi  
https://sorooshmorshedi.ir
