package piscine

import (
	"fmt"
	"os"
	"strconv"
)

func Fprime() {
	if len(os.Args) != 2 {
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 1 {
		return
	}

	d := 2
	first := true

	// factorize while divisor squared <= number
	for d*d <= n {
		for n%d == 0 {
			if !first {
				fmt.Print("*")
			}
			fmt.Print(d)
			first = false
			n /= d // reduce number
		}
		d++ // move to next possible factor
	}

	// if n is still >1, it is prime itself
	if n > 1 {
		if !first {
			fmt.Print("*")
		}
		fmt.Print(n)
	}

	fmt.Println()
}
