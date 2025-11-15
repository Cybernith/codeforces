# Mammoth's Genome Decoding — Codeforces 747B

**Source:** https://codeforces.com/problemset/problem/747/B

---

## Problem Statement

The process of mammoth's genome decoding in Berland comes to its end!

One of the few remaining tasks is to restore unrecognized nucleotides in a found chain `s`. Each nucleotide is coded with a capital letter of English alphabet: `'A'`, `'C'`, `'G'` or `'T'`. Unrecognized nucleotides are coded by a question mark `'?'`. Thus, `s` is a string consisting of letters `'A'`, `'C'`, `'G'`, `'T'` and characters `'?'`.

It is known that the number of nucleotides of each of the four types in the decoded genome of mammoth in Berland should be equal.

Your task is to decode the genome and replace each unrecognized nucleotide with one of the four types so that the number of nucleotides of each of the four types becomes equal.

### Input

The first line contains the integer `n` (`4 ≤ n ≤ 255`) — the length of the genome.

The second line contains the string `s` of length `n` — the coded genome. It consists of characters `'A'`, `'C'`, `'G'`, `'T'` and `'?'`.

### Output

If it is possible to decode the genome, print it. If there are multiple answers, print any of them. If it is not possible, print three equals signs in a row: `===` (without quotes).

---

## Solution Overview

We need to replace each `'?'` in the genome string with one of `A`, `C`, `G`, `T` such that, in the final string, **each of the four letters appears exactly `n / 4` times**.

This gives us some immediate constraints:

1. The length `n` **must** be divisible by 4.  

   - If `n % 4 != 0`, there is no way to have equal counts of the four nucleotides, so the answer is `===`.

2. Let:

   - `cntA` = current number of `'A'`
   - `cntC` = current number of `'C'`
   - `cntG` = current number of `'G'`
   - `cntT` = current number of `'T'`
   - `cntQ` = current number of `'?'`

   Let `target = n / 4`.  
   Then **each** of the counts must eventually become `target`.

3. If any existing count already exceeds the target, i.e.

   - `cntA > target` or `cntC > target` or `cntG > target` or `cntT > target`,  
     then we cannot reduce any of them (we only replace `'?'`, we never delete letters), so it is **impossible** → print `===`.

4. Otherwise, we compute the **deficit** for each nucleotide:

   ```text
   needA = target - cntA
   needC = target - cntC
   needG = target - cntG
   needT = target - cntT
   ```

   The sum of these deficits must be exactly `cntQ`. If not, then there is no way to assign `'?'` characters to exactly fill all deficits, and the answer is again `===`.

If all checks pass, we can build a valid genome by scanning the string from left to right and replacing `'?'` positions as needed.

### Construction Algorithm

1. Read `n` and the string `s`.

2. If `n % 4 != 0`, print `===` and stop.

3. Count `cntA`, `cntC`, `cntG`, `cntT` and `cntQ` in `s`.

4. Set `target = n / 4`. If any of `cntA`, `cntC`, `cntG`, `cntT` is greater than `target`, print `===` and stop.

5. Compute the deficit for each nucleotide:

   ```text
   needA = target - cntA
   needC = target - cntC
   needG = target - cntG
   needT = target - cntT
   ```

6. Check that `needA + needC + needG + needT == cntQ`. If not, print `===` and stop.

7. Now iterate over `s` again and build the final answer:

   - If the current character is not `'?'`, copy it as-is.
   - If it is `'?'`, then:
     - If `needA > 0`, put `'A'` and decrement `needA`.
     - Else if `needC > 0`, put `'C'` and decrement `needC`.
     - Else if `needG > 0`, put `'G'` and decrement `needG`.
     - Else if `needT > 0`, put `'T'` and decrement `needT`.

By the end of this process, all deficits will be filled and the resulting string will have exactly `target` occurrences of each of `A`, `C`, `G`, `T`.

### Correctness

- We never modify non-`'?'` positions, so the existing nucleotides are preserved.
- The preliminary checks ensure that none of the nucleotides exceed the allowed count.
- The deficits `needA`, `needC`, `needG`, `needT` are exactly matched to the number of `'?'` characters.  
  Therefore, assigning the `'?'` characters to fill these deficits guarantees that each nucleotide ends up with count `target`.
- We only use simple counting and a single pass for replacement, so any ordering of assignments respecting the deficits yields a valid answer.

### Complexity

- Counting characters is `O(n)`.
- Replacing `'?'` characters is also `O(n)`.
- Overall time complexity: **`O(n)`**.  
- Extra space complexity: **`O(1)`** (besides the output string).

---

## Implementation Notes

### Python (`awnser.py`)

- Reads `n` and the string `s` from standard input.
- Performs the divisibility and count checks described above.
- If decoding is possible, constructs and prints a valid decoded genome.
- If decoding is impossible, prints `===`.

The Python implementation is stored in:

```text
awnser.py
```

and is designed for direct use in a Codeforces-style environment (stdin/stdout only).

### Go (`awnser.go`)

- Mirrors the same logic and checks in Go.
- Uses buffered I/O for efficient reading/writing.
- Produces the decoded genome or `===` using the same algorithm.

The Go implementation is stored in:

```text
awnser.go
```

and is compatible with online judges that support Go.

---

## Notes

(Reserved for additional comments, links, or branding to be filled by Cybernith - sorooshmorshedi)
---
https://sorooshmorshedi.ir
