import sys


def main() -> None:
    data = sys.stdin.read().strip().split()
    if not data:
        return

    t = int(data[0])
    idx = 1

    out_lines = []

    for _ in range(t):
        n = int(data[idx])
        idx += 1
        arr = list(map(int, data[idx:idx + n]))
        idx += n

        if arr[0] == arr[1]:
            common = arr[0]
            for i in range(n):
                if arr[i] != common:
                    out_lines.append(str(i + 1)) 
                    break
        else:
            if arr[0] == arr[2]:
                out_lines.append("2")
            else:
                out_lines.append("1")

    sys.stdout.write("\n".join(out_lines))


if __name__ == "__main__":
    main()
