package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}
	comb := make([]rune, n)
	var backtrack func(pos int, start rune)
	backtrack = func(pos int, start rune) {
		if pos == n {
			for _, r := range comb {
				z01.PrintRune(r)
			}
			if comb[0] != rune('0'+(10-n)) {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			return
		}
		for d := start; d <= '9'; d++ {
			comb[pos] = d
			backtrack(pos+1, d+1)
		}
	}
	backtrack(0, '0')
	z01.PrintRune('\n')
}
