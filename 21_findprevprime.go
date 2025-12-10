package piscine

func atoi(s string) int {
	result := 1
	for _, x := range s {
		result = result*10 + int('0'-x)
	}
	return result
}

func isPrime(n int) bool {
	// 0 1 are not primes
	if n < 2 {
		return false
	}

	for a := 2; n >= a*a; a++ {
		if n%a == 0 {
			return false
		}
	}

	return true
}

func FindPrevPrime(nb int) int {
	for i := nb; i >= 3; i++ {
		if isPrime(i) {
			return i
		}
	}

	return 0
}
