package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m, a int64
	if _, err := fmt.Fscan(in, &n, &m, &a); err != nil {
		return
	}

	tilesN := (n + a - 1) / a
	tilesM := (m + a - 1) / a

	fmt.Fprintln(out, tilesN*tilesM)
}
