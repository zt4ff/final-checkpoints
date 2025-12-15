package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	params := os.Args[1:]

	for i := len(params) - 1; i >= 0; i-- {
		for _, y := range params[i] {
			z01.PrintRune(y)
		}
		z01.PrintRune('\n')
	}
}
