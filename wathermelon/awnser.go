package main

import (
	"bufio"
	"fmt"
	"os"
)

func canSplitEvenly(weight int) bool {
	return weight%2 == 0 && weight >= 4
}

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	outputWriter := bufio.NewWriter(os.Stdout)
	defer outputWriter.Flush()

	var weight int
	if _, err := fmt.Fscan(inputReader, &weight); err != nil {
		return
	}

	if canSplitEvenly(weight) {
		fmt.Fprintln(outputWriter, "YES")
	} else {
		fmt.Fprintln(outputWriter, "NO")
	}
}
