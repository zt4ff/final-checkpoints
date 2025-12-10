package piscine

import "github.com/01-edu/z01"

// how to solve
// loop through the byte two by two and print the hexcode of them using '>>4' and '&15'
//

func PrintMemory(arr [10]byte) {
	hex := "0123456789abcdef"

	for i := 0; i < 10; i += 2 {
		z01.PrintRune(rune(hex[arr[i]>>4]))
		z01.PrintRune(rune(hex[arr[i]&15]))

		z01.PrintRune(' ')

		z01.PrintRune(rune(hex[arr[i+1]>>4]))
		z01.PrintRune(rune(hex[arr[i+1]&15]))

		if i == 2 || i == 6 {
			z01.PrintRune('\n')
		} else {
			z01.PrintRune(' ')
		}
	}

	z01.PrintRune('\n')

	// ascii graphic character
	for _, x := range arr {
		if x > 31 {
			z01.PrintRune(rune(x))
		} else {
			z01.PrintRune('.')
		}
	}

}
