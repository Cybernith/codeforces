package main

import (
	"bufio"
	"fmt"
	"os"
)

func canKnightReachPrincess(hallwayMap string) bool {
	collectedKeys := make(map[rune]bool, 3)

	for _, cell := range hallwayMap {
		if cell >= 'a' && cell <= 'z' {
			collectedKeys[cell] = true
		} else {
			requiredKey := cell + ('a' - 'A') 
			if !collectedKeys[requiredKey] {
				return false
			}
		}
	}

	return true
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testCaseCount int
	if _, err := fmt.Fscan(in, &testCaseCount); err != nil {
		return
	}

	for i := 0; i < testCaseCount; i++ {
		var hallwayMap string
		fmt.Fscan(in, &hallwayMap)

		if canKnightReachPrincess(hallwayMap) {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
