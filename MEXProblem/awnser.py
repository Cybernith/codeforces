import sys

class SegmentTree:
    def __init__(self, size):
        self.size = size
        self.n = 1
        while self.n < size:
            self.n <<= 1
        self.value = [0] * (self.n * 4)
        self.lazy = [0] * (self.n * 4)

    def _push(self, index):
        delta = self.lazy[index]
        if delta != 0:
            left = index * 2
            right = left + 1
            self.value[left] += delta
            self.value[right] += delta
            self.lazy[left] += delta
            self.lazy[right] += delta
            self.lazy[index] = 0

    def _range_add(self, index, left, right, query_left, query_right, delta):
        if query_left > right or query_right < left:
            return
        if query_left <= left and right <= query_right:
            self.value[index] += delta
            self.lazy[index] += delta
            return
        self._push(index)
        middle = (left + right) // 2
        self._range_add(index * 2, left, middle, query_left, query_right, delta)
        self._range_add(index * 2 + 1, middle + 1, right, query_left, query_right, delta)
        if self.value[index * 2] >= self.value[index * 2 + 1]:
            self.value[index] = self.value[index * 2]
        else:
            self.value[index] = self.value[index * 2 + 1]

    def range_add(self, left, right, delta):
        if left > right:
            return
        self._range_add(1, 0, self.size - 1, left, right, delta)

    def _point_set(self, index, left, right, position, new_value):
        if left == right:
            self.value[index] = new_value
            self.lazy[index] = 0
            return
        self._push(index)
        middle = (left + right) // 2
        if position <= middle:
            self._point_set(index * 2, left, middle, position, new_value)
        else:
            self._point_set(index * 2 + 1, middle + 1, right, position, new_value)
        if self.value[index * 2] >= self.value[index * 2 + 1]:
            self.value[index] = self.value[index * 2]
        else:
            self.value[index] = self.value[index * 2 + 1]

    def point_set(self, position, new_value):
        self._point_set(1, 0, self.size - 1, position, new_value)

    def max_value(self):
        return self.value[1]


def solve():
    data = list(map(int, sys.stdin.read().split()))
    t = data[0]
    position = 1
    output_parts = []
    for _ in range(t):
        n = data[position]
        position += 1
        array = data[position:position + n]
        position += n
        coordinate_count = n + 1
        tree = SegmentTree(coordinate_count)
        answers = []
        for value in array:
            tree.point_set(value, 0)
            if value >= 1:
                tree.range_add(0, value - 1, 1)
            answers.append(str(tree.max_value()))
        output_parts.append(" ".join(answers))
    sys.stdout.write("\n".join(output_parts))


if __name__ == "__main__":
    solve()
