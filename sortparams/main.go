package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isVowel(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u',
		'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}

func printString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func main() {
	if len(os.Args) < 2 {
		z01.PrintRune('\n')
		return
	}
	args := os.Args[1:]
	type Pos struct {
		argIndex  int
		runeIndex int
	}
	positions := []Pos{}
	allVowels := []rune{}
	for ai, word := range args {
		for ri, r := range word {
			if isVowel(r) {
				positions = append(positions, Pos{ai, ri})
				allVowels = append(allVowels, r)
			}
		}
	}
	if len(allVowels) == 0 {
		for i, w := range args {
			printString(w)
			if i < len(args)-1 {
				z01.PrintRune(' ')
			}
		}
		return
	}
	for i, j := 0, len(allVowels)-1; i < j; i, j = i+1, j-1 {
		allVowels[i], allVowels[j] = allVowels[j], allVowels[i]
	}
	wordsRunes := make([][]rune, len(args))
	for i, w := range args {
		wordsRunes[i] = []rune(w)
	}
	for i, pos := range positions {
		wordsRunes[pos.argIndex][pos.runeIndex] = allVowels[i]
	}
	for i, wr := range wordsRunes {
		printString(string(wr))
		if i < len(wordsRunes)-1 {
			z01.PrintRune(' ')
		}
	}
}
