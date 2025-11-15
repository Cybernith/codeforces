import sys
from typing import Final


VOWELS: Final[set[str]] = {"A", "E", "I", "O", "U", "Y"}


def main() -> None:
    line = sys.stdin.readline()
    if not line:
        return

    last_letter = None
    for ch in reversed(line):
        if ch.isalpha():
            last_letter = ch.upper()
            break

    if last_letter is None:
        print("NO")
        return

    if last_letter in VOWELS:
        print("YES")
    else:
        print("NO")


if __name__ == "__main__":
    main()
