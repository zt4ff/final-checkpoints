package piscine

import (
	"fmt"
	"os"
)

// a - 97
// A - 65

func formatWord(word string) string {
	if word == "" {
		return ""
	}

	runes := []rune(word)

	for i := 0; i < len(runes)-1; i++ {
		if runes[i] >= 'A' && runes[i] <= 'Z' {
			runes[i] += 32
		}
	}

	last := len(runes) - 1
	if runes[last] >= 'a' && runes[last] <= 'z' {
		runes[last] -= 32
	}

	return string(runes)
}

func ReverseStrCap() {
	if len(os.Args) < 2 {
		return
	}

	for _, arg := range os.Args[1:] {
		word := ""
		result := ""

		for _, ch := range arg {
			if ch == ' ' {
				result += formatWord(word) + " "
				word = ""
			} else {
				word += string(ch)
			}
		}

		result += formatWord(word)
		fmt.Println(result)
	}
}
