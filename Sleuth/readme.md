# Sleuth

Vasya plays the sleuth with his friends. The rules of the game are as follows: those who play for the first time, that is Vasya is the sleuth, he should investigate a "crime" and find out what is happening. He can ask any questions whatsoever that can be answered with "Yes" or "No". All the rest agree beforehand to answer the questions like that: if the question’s last letter is a vowel, they answer "Yes" and if the last letter is a consonant, they answer "No". Of course, the sleuth knows nothing about it and his task is to understand that.

Unfortunately, Vasya is not very smart. After 5 hours of endless stupid questions everybody except Vasya got bored. That’s why Vasya’s friends ask you to write a program that would give answers instead of them.

The English alphabet vowels are: `A, E, I, O, U, Y`.

The English alphabet consonants are: `B, C, D, F, G, H, J, K, L, M, N, P, Q, R, S, T, V, W, X, Z`.

## Input

The single line contains a question represented by a non-empty line consisting of large and small Latin letters, spaces and a question mark. The line length does not exceed 100. It is guaranteed that the question mark occurs exactly once in the line — as the last symbol and that the line contains at least one letter.

## Output

Print answer for the question in a single line: `YES` if the answer is "Yes", `NO` if the answer is "No".

Remember that in the reply to the question the last letter, not the last character counts. That is, the spaces and the question mark do not count as letters.

---

## Solution Overview

We are given one line that ends with a question mark `?`. The line contains:

- uppercase and lowercase Latin letters,
- spaces,
- exactly one `?` at the end,
- and at least one letter somewhere in the line.

All of Vasya's friends answer **Yes** if the **last letter** of the question is a vowel, and **No** if the last letter is a consonant. We must output:

- `YES` if the last letter (ignoring spaces and the final `?`) is in `{A, E, I, O, U, Y}` (case-insensitive),
- otherwise `NO`.

### Algorithm

1. Read the entire line as a string `s`.
2. Starting from the end of the string, move left until we find a character that is a letter (`A`–`Z` or `a`–`z`).
3. Convert that letter to uppercase.
4. If this uppercase letter is one of `A, E, I, O, U, Y`, print `YES`.  
   Otherwise, print `NO`.

Because the line length is at most 100, this simple linear scan from the end is more than efficient enough.

### Complexity

- Time complexity: `O(L)`, where `L` is the length of the string (at most 100).
- Space complexity: `O(1)` extra memory (we only keep a few variables).

---

## Implementation Notes

### Python (`awnser.py`)

- Reads the full line using `sys.stdin.readline()`.
- Scans from right to left to find the last alphabetic character.
- Checks if it is a vowel and prints `YES` or `NO` accordingly.

### Go (`awnser.go`)

- Reads the full line using a buffered reader.
- Iterates from the end of the string to find the last alphabetic character.
- Converts the character to uppercase and checks if it is a vowel.
- Prints `YES` or `NO` accordingly.

---

## Notes

CyberNith — Soroosh Morshedi  
https://sorooshmorshedi.ir
