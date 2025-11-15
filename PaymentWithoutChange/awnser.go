package main

import (
	"bufio"
	"fmt"
	"os"
)

func canPayExact(a, b, n, s int64) bool {
	maxNCoins := a
	if s/n < maxNCoins {
		maxNCoins = s / n
	}

	remaining := s - maxNCoins*n
	return remaining <= b
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var q int
	if _, err := fmt.Fscan(in, &q); err != nil {
		return
	}

	for i := 0; i < q; i++ {
		var a, b, n, s int64
		fmt.Fscan(in, &a, &b, &n, &s)

		if canPayExact(a, b, n, s) {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
