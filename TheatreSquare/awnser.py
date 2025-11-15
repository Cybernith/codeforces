import sys


def main() -> None:
    data = sys.stdin.read().strip().split()
    if len(data) < 3:
        return

    n = int(data[0])
    m = int(data[1])
    a = int(data[2])

    tiles_n = (n + a - 1) // a
    tiles_m = (m + a - 1) // a

    print(tiles_n * tiles_m)


if __name__ == "__main__":
    main()
