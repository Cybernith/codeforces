# Candies

Polycarpus has got *n* candies and *m* friends (*n* ≥ *m*). He wants to make a New Year present with candies to each friend.  Polycarpus is planning to present all candies and he wants to do this in the **fairest** (that is, most equal) manner. He wants to choose such *a**i*, where *a**i* is the number of candies in the *i*-th friend's present, that the maximum *a**i* differs from the least *a**i* as little as possible.

For example, if *n* is divisible by *m*, then he is going to present the same number of candies to all his friends, that is, the maximum *a**i* won't differ from the minimum one.

### Input

The single line of the input contains a pair of space-separated positive integers *n*, *m* (1 ≤ *n*, *m* ≤ 100;*n* ≥ *m*) — the number of candies and the number of Polycarpus's friends.

### Output


Print the required sequence *a*1, *a*2, ..., *a**m*, where *a**i* is the number of candies in the *i*-th friend's present. All numbers *a**i* must be positive integers, total up to *n*, the maximum one should differ from the minimum one by the smallest possible value.



---------------------------------------


# Candies (Codeforces-style Task)

This folder contains a solution for the **"Candies"** problem (a typical Codeforces-style task).  
Every folder in the repository corresponds to one problem and its solution.

## Problem Overview

Polycarpus has:

- `n` candies
- `m` friends (`n >= m`)

He wants to give **all** candies away and distribute them as **fairly** as possible.

We must choose integers:

\( a_1, a_2, \dots, a_m \)

such that:

- All \( a_i > 0 \)
- \( a_1 + a_2 + \dots + a_m = n \)
- The difference between the **maximum** and **minimum** value in the sequence is as small as possible.

### Optimal Strategy

Let:

- `base = n // m`
- `remainder = n % m`

Then a fair distribution is:

- Give `base + 1` candies to the first `remainder` friends
- Give `base` candies to each of the remaining `m - remainder` friends

This minimizes `max(a_i) - min(a_i)` while keeping all `a_i` positive and summing to `n`.

---

## Files

- `candies.py` — Python solution (OOP, PEP 8, ready for competitive use or as a module)
- `main.go` — Go solution (idiomatic Go, suitable for competitive programming)
- `LICENSE` — MIT License
- `README.md` — This documentation

---

## Python Solution (`candies.py`)

### Design

The main object is:

- `CandiesDistributor`
  - Attributes:
    - `total_candies: int`
    - `friends_count: int`
  - Methods:
    - `validate()` — checks basic constraints
    - `distribute() -> List[int]` — returns the distribution

### Usage (Command Line)

The script reads `n` and `m` from standard input and prints the sequence of candies as space-separated integers.

Example:

```bash
echo "7 4" | python candies.py
```

Output:

```text
2 2 2 1
```

### Usage (As a Module)

```python
from candies import CandiesDistributor

distributor = CandiesDistributor(total_candies=7, friends_count=4)
print(distributor.distribute())  # [2, 2, 2, 1]
```

---

## Go Solution (`main.go`)

The Go version defines a similar structure:

- `CandiesDistributor` with fields:
  - `TotalCandies int`
  - `FriendsCount int`
- Method:
  - `Distribute() []int`

### Usage (Command Line)

```bash
echo "7 4" | go run main.go
```

Output:

```text
2 2 2 1
```

The Go solution is structured to be compatible with typical competitive programming environments (fast I/O via `bufio`).

---

## License

This folder (and its contents) is licensed under the **MIT License**.  
See the `LICENSE` file for full terms.
