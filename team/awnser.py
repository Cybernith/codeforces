import sys
from typing import Iterable


def is_problem_implemented(confidences: Iterable[int]) -> bool:
    return sum(confidences) >= 2


def main() -> None:
    input_data = sys.stdin.read().strip().split()
    if not input_data:
        return

    problems_count = int(input_data[0])
    index = 1

    implemented_problems_count = 0

    for _ in range(problems_count):
        petya_confidence = int(input_data[index])
        vasya_confidence = int(input_data[index + 1])
        tonya_confidence = int(input_data[index + 2])
        index += 3

        if is_problem_implemented(
            (petya_confidence, vasya_confidence, tonya_confidence)
        ):
            implemented_problems_count += 1

    print(implemented_problems_count)


if __name__ == "__main__":
    main()
