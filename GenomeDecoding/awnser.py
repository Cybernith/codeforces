from __future__ import annotations

import sys


def decode_genome(length: int, genome: str) -> str:
    if length % 4 != 0:
        return "==="

    target = length // 4

    count_a = genome.count("A")
    count_c = genome.count("C")
    count_g = genome.count("G")
    count_t = genome.count("T")
    count_q = genome.count("?")

    if (
        count_a > target
        or count_c > target
        or count_g > target
        or count_t > target
    ):
        return "==="

    need_a = target - count_a
    need_c = target - count_c
    need_g = target - count_g
    need_t = target - count_t

    if need_a + need_c + need_g + need_t != count_q:
        return "==="

    result_chars: list[str] = []
    for ch in genome:
        if ch != "?":
            result_chars.append(ch)
            continue

        if need_a > 0:
            result_chars.append("A")
            need_a -= 1
        elif need_c > 0:
            result_chars.append("C")
            need_c -= 1
        elif need_g > 0:
            result_chars.append("G")
            need_g -= 1
        elif need_t > 0:
            result_chars.append("T")
            need_t -= 1
        else:
            result_chars.append("A")

    return "".join(result_chars)


def main() -> None:
    data = sys.stdin.read().strip().split()
    if len(data) < 2:
        print("===")
        return

    length = int(data[0])
    genome = data[1]

    if len(genome) != length:
        length = len(genome)

    decoded = decode_genome(length, genome)
    print(decoded)


if __name__ == "__main__":
    main()
