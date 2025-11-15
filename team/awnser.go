package main

import (
	"bufio"
	"fmt"
	"os"
)

func isProblemImplemented(petyaConfidence, vasyaConfidence, tonyaConfidence int) bool {
	sumOfConfidences := petyaConfidence + vasyaConfidence + tonyaConfidence
	return sumOfConfidences >= 2
}

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	outputWriter := bufio.NewWriter(os.Stdout)
	defer outputWriter.Flush()

	var problemsCount int
	if _, err := fmt.Fscan(inputReader, &problemsCount); err != nil {
		return
	}

	implementedProblemsCount := 0

	for problemIndex := 0; problemIndex < problemsCount; problemIndex++ {
		var petyaConfidence, vasyaConfidence, tonyaConfidence int
		fmt.Fscan(inputReader, &petyaConfidence, &vasyaConfidence, &tonyaConfidence)

		if isProblemImplemented(petyaConfidence, vasyaConfidence, tonyaConfidence) {
			implementedProblemsCount++
		}
	}

	fmt.Fprintln(outputWriter, implementedProblemsCount)
}
