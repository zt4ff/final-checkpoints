package piscine

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func print3(s string) {
	for _, x := range s {
		z01.PrintRune(x)
	}
	z01.PrintRune('\n')
}

func isPrime1(n int) bool {
	// 0 1 are not prime
	if n < 2 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func atoi(n string) int {
	result := 0

	for _, x := range n {
		result = result*10 + int(x-'0')
	}

	return result
}

func AddPrimeSum() {
	z01.PrintRune('1')
	args := os.Args[1:]

	count := 0

	if len(args) != 1 || atoi(args[0]) < 0 {
		print2(itoa(count))
	}

	x := atoi(args[0])

	for i := 2; i <= x; i++ {
		if isPrime1(i) {
			count += i
		}
	}

	fmt.Println(count)
}
