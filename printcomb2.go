package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for a := '0'; a <= '9'; a++ {
		for b := '0'; b <= '9'; b++ {
			j := b + 1
			for i := a; i <= '9'; i++ {
				for ; j <= '9'; j++ {
					z01.PrintRune(a)
					z01.PrintRune(b)
					z01.PrintRune(' ')
					z01.PrintRune(i)
					z01.PrintRune(j)
					if a < '9' || b < '8' || i < '9' || j < '9' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}

				}
				j = '0'
			}
		}
	}
	z01.PrintRune('\n')
}
