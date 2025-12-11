package piscine

import (
	"fmt"
	"os"
)

func RomanNumerals() {
	if len(os.Args) != 2 {
		return
	}

	n := 0
	for _, c := range os.Args[1] {
		if c < '0' || c > '9' {
			fmt.Println("ERROR: cannot convert to roman digit")
			return
		}
		n = n*10 + int(c-'0')
	}

	if n <= 0 || n >= 4000 {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}

	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	roman := ""
	calc := ""

	tmp := n
	for i := 0; i < len(values); i++ {
		for tmp >= values[i] {
			tmp -= values[i]
			roman += symbols[i]

			if calc != "" {
				calc += "+"
			}
			if symbols[i] == "CM" {
				calc += "(M-C)"
				continue
			}
			if symbols[i] == "CD" {
				calc += "(D-C)"
				continue
			}
			if symbols[i] == "XC" {
				calc += "(C-X)"
				continue
			}
			if symbols[i] == "XL" {
				calc += "(L-X)"
				continue
			}
			if symbols[i] == "IX" {
				calc += "(X-I)"
				continue
			}
			if symbols[i] == "IV" {
				calc += "(V-I)"
				continue
			}
			calc += symbols[i]
		}
	}

	fmt.Println(calc)
	fmt.Println(roman)
}
