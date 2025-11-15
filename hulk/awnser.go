package main

import (
	"bufio"
	"fmt"
	"os"
)

func buildHulkFeeling(layersCount int) string {
	parts := make([]string, 0, layersCount)

	for layerIndex := 1; layerIndex <= layersCount; layerIndex++ {
		if layerIndex%2 == 1 {
			parts = append(parts, "I hate")
		} else {
			parts = append(parts, "I love")
		}
	}

	result := ""
	for i, segment := range parts {
		if i > 0 {
			result += " that "
		}
		result += segment
	}
	result += " it"

	return result
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var layersCount int
	if _, err := fmt.Fscan(in, &layersCount); err != nil {
		return
	}

	fmt.Fprintln(out, buildHulkFeeling(layersCount))
}
