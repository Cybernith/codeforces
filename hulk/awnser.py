import sys


def build_hulk_feeling(layers_count: int) -> str:
    parts: list[str] = []

    for layer_index in range(1, layers_count + 1):
        if layer_index % 2 == 1:
            parts.append("I hate")
        else:
            parts.append("I love")

    return " that ".join(parts) + " it"


def main() -> None:
    data = sys.stdin.read().strip().split()
    if not data:
        return

    layers_count = int(data[0])
    result = build_hulk_feeling(layers_count)
    print(result)


if __name__ == "__main__":
    main()
