# Team - Codeforces 231A

**Source:** https://codeforces.com/problemset/problem/231/A

---

## Problem Statement

Three best friends - Petya, Vasya and Tonya - decided to form a team and take part in programming contests. Participants are usually given several problems in a contest.

Long before the start, the friends agreed on the following rule:

> They will implement a solution for a problem **if and only if at least two of them are sure about the solution**.  
> Otherwise, they will skip that problem.

The contest offers `n` problems. For each problem, we know which of the three friends is sure about the solution.

Your task is to determine **how many problems** the friends will implement.

### Input

- The first line contains a single integer `n` (`1 ≤ n ≤ 1000`) — the number of problems.
- Each of the next `n` lines contains three integers, each either `0` or `1`:
  - The first integer represents Petya's opinion.
  - The second integer represents Vasya's opinion.
  - The third integer represents Tonya's opinion.

If a number is `1`, then the corresponding friend is sure about the solution to that problem; otherwise, the number is `0`.

### Output

Print a single integer — the **number of problems** for which the friends will write a solution.

---

## Solution Overview

For each problem, we are given three binary values indicating whether each friend is confident about solving that problem. The friends agreed to solve a problem **only if at least two of them are sure**.

Thus, for each line of three integers:

- Let them be `p`, `v`, `t` (for Petya, Vasya, Tonya).
- Compute the sum: `p + v + t`.
- If `p + v + t ≥ 2`, it means at least two friends are sure → they will implement the solution for this problem.

We simply need to:

1. Initialize a counter `implemented_problems` to 0.
2. For each of the `n` lines:
   - Read the three integers.
   - If their sum is at least 2, increment `implemented_problems` by 1.
3. After processing all problems, print `implemented_problems`.

### Example

Input:

```text
5
1 1 0
1 1 1
0 0 0
1 0 1
1 0 0
```

Processing:

1. `1 1 0` → sum = 2 → solve this problem ✅
2. `1 1 1` → sum = 3 → solve this problem ✅
3. `0 0 0` → sum = 0 → skip ❌
4. `1 0 1` → sum = 2 → solve this problem ✅
5. `1 0 0` → sum = 1 → skip ❌

Total problems solved: `3`.

Output:

```text
3
```

### Complexity

Let `n` be the number of problems.

- For each problem, we read three integers and compute their sum — `O(1)` work per problem.
- Overall time complexity: **O(n)**.
- Memory usage: **O(1)** extra space (only a few counters and temporary variables).

---

## Implementation Notes

### Python (`awnser.py`)

- Reads the integer `n` from standard input.
- For each of the next `n` lines:
  - Parses three integers: Petya's, Vasya's and Tonya's opinions.
  - Sums them and checks whether the result is at least 2.
- Maintains a running count of how many problems will be implemented.
- Prints the final count.

The implementation uses:

- Clear, descriptive variable names.
- A helper function to encapsulate the decision logic per problem.

### Go (`awnser.go`)

- Uses buffered input and output (`bufio.NewReader`, `bufio.NewWriter`) for efficiency.
- Reads `n`, then loops over `n` problem descriptions.
- For each problem:
  - Reads the three integers.
  - Sums them and checks whether the sum is at least 2.
  - Increments `implementedProblemsCount` if so.
- Prints the final `implementedProblemsCount` value.

---

## Notes

CyberNith - Soroosh Morshedi  
https://sorooshmorshedi.ir
