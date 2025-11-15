from __future__ import annotations

from dataclasses import dataclass
from typing import List


@dataclass
class DisturbedPeopleSolver:
    flats: List[int]

    def minimum_switch_offs(self) -> int:
        state = self.flats[:]
        n = len(state)
        operations = 0

        for i in range(1, n - 1):
            if state[i - 1] == 1 and state[i] == 0 and state[i + 1] == 1:
                state[i + 1] = 0
                operations += 1

        return operations


def main() -> None:
    import sys

    data = sys.stdin.read().split()
    if not data:
        return

    n = int(data[0])
    flats = list(map(int, data[1:1 + n]))

    solver = DisturbedPeopleSolver(flats=flats)
    result = solver.minimum_switch_offs()
    print(result)


if __name__ == "__main__":
    main()
