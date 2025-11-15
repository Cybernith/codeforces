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

	var t int
	if _, err := fmt.Fscan(in, &t); err != nil {
		return
	}

	for ; t > 0; t-- {
		var n int
		fmt.Fscan(in, &n)

		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &arr[i])
		}

		if arr[0] == arr[1] {
			common := arr[0]
			for i := 0; i < n; i++ {
				if arr[i] != common {
					fmt.Fprintln(out, i+1)
					break
				}
			}
		} else {
			if arr[0] == arr[2] {
				fmt.Fprintln(out, 2)
			} else {
				fmt.Fprintln(out, 1)
			}
		}
	}
}
