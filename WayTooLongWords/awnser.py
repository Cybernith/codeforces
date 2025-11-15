import sys


def abbreviate(word: str) -> str:
    length = len(word)
    if length <= 10:
        return word
    return f"{word[0]}{length - 2}{word[-1]}"


def main() -> None:
    data = sys.stdin.read().strip().split()
    if not data:
        return

    n = int(data[0])
    words = data[1:1 + n]

    out_lines = [abbreviate(w) for w in words]
    sys.stdout.write("\n".join(out_lines))


if __name__ == "__main__":
    main()
