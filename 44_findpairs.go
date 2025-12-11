package piscine

import (
	"fmt"
	"os"
)

func SplitWord(s string, sep string) []string {

	var out []string

	n := len(sep)
	word := ""

	for i := 0; i < len(s); {
		if i+n < len(s) && s[i:i+n] == sep {
			out = append(out, word)
			i += n
			word = ""
		} else {
			word += string(s[i])
			i++
		}
	}

	// final word
	out = append(out, word)

	return out
}

func Revwstr() {
	args := os.Args[1:]

	word := SplitWord(args[0], " ")

	for i := len(word) - 1; i >= 0; i-- {
		if i != len(word)-1 {
			fmt.Print(" ")
		}
		fmt.Print(word[i])
	}

	fmt.Println()
}
