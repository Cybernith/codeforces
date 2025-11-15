package main

import (
	"bufio"
	"fmt"
	"os"
)

func abbreviate(word string) string {
	if len(word) <= 10 {
		return word
	}
	middleCount := len(word) - 2
	return fmt.Sprintf("%c%d%c", word[0], middleCount, word[len(word)-1])
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	if _, err := fmt.Fscan(in, &n); err != nil {
		return
	}

	for i := 0; i < n; i++ {
		var w string
		fmt.Fscan(in, &w)
		fmt.Fprintln(out, abbreviate(w))
	}
}
