package piscine

import (
	"os"

	"github.com/01-edu/z01"
)

func print1(s string) {
	for _, x := range s {
		z01.PrintRune(x)
	}
	z01.PrintRune('\n')
}

func Union() {
	args := os.Args[1:]

	if len(args) != 2 {
		z01.PrintRune('\n')
	}

	ans := ""

	dict := make(map[rune](int))

	for _, x := range args {
		for _, y := range x {
			if dict[y] != 1 {
				ans += string(y)
				dict[y]++
			}
		}
	}

	print1(ans)
}
