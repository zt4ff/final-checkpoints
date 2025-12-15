package main

import "github.com/01-edu/z01"

func main() {
	// Loop through all lowercase letters from 'a' to 'z'
	// := is Go's short declaration operator
	// also in go, characters are actually numbers so 'a' = 97, 'b' = 98, ...
	for char := 'z'; char >= 'a'; char-- {
		z01.PrintRune(char) // Print each character
	}
	z01.PrintRune('\n') // Print newline at the end
}
