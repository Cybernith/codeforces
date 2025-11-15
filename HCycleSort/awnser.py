from __future__ import annotations

import sys
from typing import List, Tuple


class CycleSortSolver:

    @staticmethod
    def process_case(
        length_a: int,
        length_b: int,
        operations_count: int,
        array_a: List[int],
        array_b: List[int],
    ) -> Tuple[List[int], List[int]]:
        if operations_count == 0:
            return array_a, array_b

        index_a = 0
        index_b = 0

        for _ in range(operations_count):
            if array_a[index_a] > array_b[index_b]:
                array_a[index_a], array_b[index_b] = array_b[index_b], array_a[index_a]

            index_a += 1
            if index_a == length_a:
                index_a = 0

            index_b += 1
            if index_b == length_b:
                index_b = 0

        return array_a, array_b

    def solve(self, data: str | None = None) -> str:
        if data is None:
            data = sys.stdin.read()
        tokens = list(map(int, data.split()))
        pointer = 0

        test_cases = tokens[pointer]
        pointer += 1

        output_lines: List[str] = []

        for _ in range(test_cases):
            length_a = tokens[pointer]
            length_b = tokens[pointer + 1]
            operations_count = tokens[pointer + 2]
            pointer += 3

            array_a = tokens[pointer:pointer + length_a]
            pointer += length_a
            array_b = tokens[pointer:pointer + length_b]
            pointer += length_b

            final_a, final_b = self.process_case(
                length_a,
                length_b,
                operations_count,
                array_a,
                array_b,
            )

            output_lines.append(" ".join(map(str, final_a)))
            output_lines.append(" ".join(map(str, final_b)))

        return "\n".join(output_lines)


def main() -> None:
    solver = CycleSortSolver()
    result = solver.solve()
    sys.stdout.write(result)


if __name__ == "__main__":
    main()
