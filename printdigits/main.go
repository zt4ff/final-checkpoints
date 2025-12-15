package main

import "github.com/01-edu/z01"

func main() {
	for compteur := '0'; compteur <= '9'; compteur++ {
		z01.PrintRune(compteur)
	}
	z01.PrintRune('\n')
}
