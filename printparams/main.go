package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	params := os.Args[1:]

	for _, x := range params {
		for _, y := range x {
			z01.PrintRune(y)
		}
		z01.PrintRune('\n')
	}
}
