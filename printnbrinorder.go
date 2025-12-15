package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}

	if n == 0 {
		z01.PrintRune('0')
		return
	}

	var counts [10]int

	for n > 0 {
		d := n % 10
		counts[d]++
		n /= 10
	}

	for digit := 0; digit <= 9; digit++ {
		for counts[digit] > 0 {
			z01.PrintRune(rune('0' + digit))
			counts[digit]--
		}
	}
}
