from __future__ import annotations

from dataclasses import dataclass
from typing import List


@dataclass
class CandiesDistributor:
    total_candies: int
    friends_count: int

    def validate(self) -> None:
        if self.total_candies <= 0:
            raise ValueError("total_candies must be a positive integer.")
        if self.friends_count <= 0:
            raise ValueError("friends_count must be a positive integer.")
        if self.total_candies < self.friends_count:
            raise ValueError("total_candies must be greater than or equal to friends_count.")

    def distribute(self) -> List[int]:
        self.validate()

        base = self.total_candies // self.friends_count
        remainder = self.total_candies % self.friends_count

        distribution: List[int] = [base + 1] * remainder + [base] * (self.friends_count - remainder)
        return distribution


def main() -> None:
    import sys

    data = sys.stdin.read().split()
    if len(data) < 2:
        raise SystemExit("Expected two integers: n m")

    n, m = map(int, data[:2])
    distributor = CandiesDistributor(total_candies=n, friends_count=m)
    distribution = distributor.distribute()
    print(" ".join(str(value) for value in distribution))


if __name__ == "__main__":
    main()
