package main

import (
	"bufio"
	"fmt"
	"os"
)

func decodeGenome(n int, s string) string {
	if n%4 != 0 {
		return "==="
	}

	target := n / 4

	countA, countC, countG, countT, countQ := 0, 0, 0, 0, 0
	for _, ch := range s {
		switch ch {
		case 'A':
			countA++
		case 'C':
			countC++
		case 'G':
			countG++
		case 'T':
			countT++
		case '?':
			countQ++
		}
	}

	if countA > target || countC > target || countG > target || countT > target {
		return "==="
	}

	needA := target - countA
	needC := target - countC
	needG := target - countG
	needT := target - countT

	if needA+needC+needG+needT != countQ {
		return "==="
	}

	result := make([]byte, n)
	for i := 0; i < n; i++ {
		ch := s[i]
		if ch != '?' {
			result[i] = ch
			continue
		}
		if needA > 0 {
			result[i] = 'A'
			needA--
		} else if needC > 0 {
			result[i] = 'C'
			needC--
		} else if needG > 0 {
			result[i] = 'G'
			needG--
		} else if needT > 0 {
			result[i] = 'T'
			needT--
		} else {
			result[i] = 'A'
		}
	}

	return string(result)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	if _, err := fmt.Fscan(in, &n); err != nil {
		fmt.Fprintln(out, "===")
		return
	}

	var s string
	if _, err := fmt.Fscan(in, &s); err != nil {
		fmt.Fprintln(out, "===")
		return
	}

	if len(s) != n {
		n = len(s)
	}

	result := decodeGenome(n, s)
	fmt.Fprintln(out, result)
}
