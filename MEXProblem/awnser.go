package main

import (
	"bufio"
	"fmt"
	"os"
)

type SegmentTree struct {
	size int
	max  []int
	lazy []int
}

func NewSegmentTree(size int) *SegmentTree {
	n := 1
	for n < size {
		n <<= 1
	}
	return &SegmentTree{
		size: size,
		max:  make([]int, n*4),
		lazy: make([]int, n*4),
	}
}

func (st *SegmentTree) push(index int) {
	delta := st.lazy[index]
	if delta != 0 {
		left := index * 2
		right := left + 1
		st.max[left] += delta
		st.max[right] += delta
		st.lazy[left] += delta
		st.lazy[right] += delta
		st.lazy[index] = 0
	}
}

func (st *SegmentTree) rangeAdd(index, left, right, queryLeft, queryRight, delta int) {
	if queryLeft > right || queryRight < left {
		return
	}
	if queryLeft <= left && right <= queryRight {
		st.max[index] += delta
		st.lazy[index] += delta
		return
	}
	st.push(index)
	middle := (left + right) / 2
	st.rangeAdd(index*2, left, middle, queryLeft, queryRight, delta)
	st.rangeAdd(index*2+1, middle+1, right, queryLeft, queryRight, delta)
	if st.max[index*2] >= st.max[index*2+1] {
		st.max[index] = st.max[index*2]
	} else {
		st.max[index] = st.max[index*2+1]
	}
}

func (st *SegmentTree) RangeAdd(left, right, delta int) {
	if left > right {
		return
	}
	st.rangeAdd(1, 0, st.size-1, left, right, delta)
}

func (st *SegmentTree) pointSet(index, left, right, position, newValue int) {
	if left == right {
		st.max[index] = newValue
		st.lazy[index] = 0
		return
	}
	st.push(index)
	middle := (left + right) / 2
	if position <= middle {
		st.pointSet(index*2, left, middle, position, newValue)
	} else {
		st.pointSet(index*2+1, middle+1, right, position, newValue)
	}
	if st.max[index*2] >= st.max[index*2+1] {
		st.max[index] = st.max[index*2]
	} else {
		st.max[index] = st.max[index*2+1]
	}
}

func (st *SegmentTree) PointSet(position, newValue int) {
	st.pointSet(1, 0, st.size-1, position, newValue)
}

func (st *SegmentTree) MaxValue() int {
	return st.max[1]
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		var n int
		fmt.Fscan(in, &n)
		array := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &array[i])
		}
		coordinateCount := n + 1
		tree := NewSegmentTree(coordinateCount)
		for i := 0; i < n; i++ {
			value := array[i]
			tree.PointSet(value, 0)
			if value >= 1 {
				tree.RangeAdd(0, value-1, 1)
			}
			if i+1 == n {
				fmt.Fprint(out, tree.MaxValue())
			} else {
				fmt.Fprint(out, tree.MaxValue(), " ")
			}
		}
		if t > 1 {
			fmt.Fprintln(out)
		}
	}
}
