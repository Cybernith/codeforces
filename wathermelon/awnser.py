import sys


def can_split_evenly(weight: int) -> bool:
    return weight % 2 == 0 and weight >= 4


def main() -> None:
    data = sys.stdin.read().strip().split()
    if not data:
        return

    weight = int(data[0])

    if can_split_evenly(weight):
        print("YES")
    else:
        print("NO")


if __name__ == "__main__":
    main()
