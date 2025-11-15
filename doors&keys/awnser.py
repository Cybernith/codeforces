import sys
from typing import Set


def can_knight_reach_princess(hallway_map: str) -> bool:
    collected_keys: Set[str] = set()

    for cell in hallway_map:
        if cell.islower():
            collected_keys.add(cell)
        else:
            required_key = cell.lower()
            if required_key not in collected_keys:
                return False

    return True


def main() -> None:
    input_data = sys.stdin.read().strip().split()
    if not input_data:
        return

    test_case_count = int(input_data[0])
    hallway_strings = input_data[1:1 + test_case_count]

    outputs = []
    for hallway_map in hallway_strings:
        if can_knight_reach_princess(hallway_map):
            outputs.append("YES")
        else:
            outputs.append("NO")

    sys.stdout.write("\n".join(outputs))


if __name__ == "__main__":
    main()
