package main

import (
	"bufio"
	"fmt"
	"os"
)

type CycleSortSolver struct{}

func (s *CycleSortSolver) processCase(
	n, m int,
	k int64,
	a, b []int,
) ([]int, []int) {
	if k == 0 {
		return a, b
	}

	ia, ib := 0, 0
	for i := int64(0); i < k; i++ {
		if a[ia] > b[ib] {
			a[ia], b[ib] = b[ib], a[ia]
		}
		ia++
		if ia == n {
			ia = 0
		}
		ib++
		if ib == m {
			ib = 0
		}
	}
	return a, b
}

func (s *CycleSortSolver) solve() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	if _, err := fmt.Fscan(in, &t); err != nil {
		return
	}

	for ; t > 0; t-- {
		var n, m int
		var k int64
		fmt.Fscan(in, &n, &m, &k)

		a := make([]int, n)
		b := make([]int, m)

		for i := 0; i < n; i++ {
			fmt.Fscan(in, &a[i])
		}
		for i := 0; i < m; i++ {
			fmt.Fscan(in, &b[i])
		}

		a, b = s.processCase(n, m, k, a, b)

		for i := 0; i < n; i++ {
			if i > 0 {
				fmt.Fprint(out, " ")
			}
			fmt.Fprint(out, a[i])
		}
		fmt.Fprintln(out)

		for i := 0; i < m; i++ {
			if i > 0 {
				fmt.Fprint(out, " ")
			}
			fmt.Fprint(out, b[i])
		}
		fmt.Fprintln(out)
	}
}

func main() {
	solver := &CycleSortSolver{}
	solver.solve()
}
