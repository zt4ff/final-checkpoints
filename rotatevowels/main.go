package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		z01.PrintRune('\n')
		return
	}
	// Collect all vowels from all arguments
	var vowels []rune
	for _, arg := range args {
		for _, char := range arg {
			if isVowel(char) {
				vowels = append(vowels, char)
			}
		}
	}

	// If no vowels found, print arguments as they are
	if len(vowels) == 0 {
		for i, arg := range args {
			for _, char := range arg {
				z01.PrintRune(char)
			}
			if i < len(args)-1 {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
		return
	}

	// Rotate vowels by reversing their order
	reverseVowels(vowels)

	// Replace vowels in arguments with rotated vowels
	vowelIndex := 0
	for i, arg := range args {
		for _, char := range arg {
			if isVowel(char) {
				z01.PrintRune(vowels[vowelIndex])
				vowelIndex++
			} else {
				z01.PrintRune(char)
			}
		}
		if i < len(args)-1 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}

func isVowel(char rune) bool {
	switch char {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}

func reverseVowels(vowels []rune) {
	for i, j := 0, len(vowels)-1; i < j; i, j = i+1, j-1 {
		vowels[i], vowels[j] = vowels[j], vowels[i]
	}
}
