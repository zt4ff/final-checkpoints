package piscine

import (
	"os"

	"github.com/01-edu/z01"
)

func print2(s string) {
	for _, x := range s {
		z01.PrintRune(x)
	}

	z01.PrintRune('\n')
}

func WdMatch() {
	args := os.Args[1:]

	if len(args) != 2 {
		return
	}

	currentIndex := 0

	for j, x := range args[0] {
		if currentIndex >= len(args[1]) {
			return
		}
		for i, y := range args[1][currentIndex:] {
			if x == y {
				if j == len(args[0])-1 {
					print2(args[0])
					return
				}
				currentIndex = i
			}

		}
	}

}
