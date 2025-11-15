import sys


def can_pay_exact(a: int, b: int, n: int, s: int) -> bool:
    max_n_coins = min(a, s // n)
    remaining = s - max_n_coins * n
    return remaining <= b


def main() -> None:
    data = sys.stdin.read().strip().split()
    if not data:
        return

    q = int(data[0])
    idx = 1

    out_lines = []
    for _ in range(q):
        a = int(data[idx])
        b = int(data[idx + 1])
        n = int(data[idx + 2])
        s = int(data[idx + 3])
        idx += 4

        if can_pay_exact(a, b, n, s):
            out_lines.append("YES")
        else:
            out_lines.append("NO")

    sys.stdout.write("\n".join(out_lines))


if __name__ == "__main__":
    main()
