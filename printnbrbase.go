package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	// Validate base
	if !isValidBase(base) {
		printStr("NV")
		return
	}

	baseRunes := []rune(base)
	baseLen := len(baseRunes)

	// Handle negative numbers
	if nbr < 0 {
		z01.PrintRune('-')
		// careful with int min: convert by multiplying instead of negating
		if nbr == -nbr { // overflow check
			// This is the minimal int, handle manually
			// Convert using nbr/baseLen recursively while nbr is negative
			printNbrRec(nbr, baseRunes, baseLen)
			return
		}
		nbr = -nbr
	}

	printNbrRec(nbr, baseRunes, baseLen)
}

// func isValidBase(base string) bool {
// 	runes := []rune(base)
// 	if len(runes) < 2 {
// 		return false
// 	}
// 	for i := 0; i < len(runes); i++ {
// 		if runes[i] == '+' || runes[i] == '-' {
// 			return false
// 		}
// 		for j := i + 1; j < len(runes); j++ {
// 			if runes[i] == runes[j] {
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }

func printNbrRec(n int, base []rune, baseLen int) {
	if n <= -baseLen || n >= baseLen {
		printNbrRec(n/baseLen, base, baseLen)
	}
	digit := n % baseLen
	if digit < 0 {
		digit = -digit
	}
	z01.PrintRune(base[digit])
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
