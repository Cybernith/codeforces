package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	line, err := in.ReadString('\n')
	if err != nil && len(line) == 0 {
		return
	}

	var lastLetter rune
	found := false

	for i := len(line) - 1; i >= 0; i-- {
		ch := rune(line[i])
		if unicode.IsLetter(ch) {
			lastLetter = unicode.ToUpper(ch)
			found = true
			break
		}
	}

	if !found {
		fmt.Fprintln(out, "NO")
		return
	}

	if isVowel(lastLetter) {
		fmt.Fprintln(out, "YES")
	} else {
		fmt.Fprintln(out, "NO")
	}
}

func isVowel(ch rune) bool {
	switch ch {
	case 'A', 'E', 'I', 'O', 'U', 'Y':
		return true
	default:
		return false
	}
}
