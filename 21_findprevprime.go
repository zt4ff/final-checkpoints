package piscine

func isPrime(n int) bool {
	// 0 1 are not primes
	if n < 2 {
		return false
	}

	for a := 2; a*a <= n; a++ {
		if n%a == 0 {
			return false
		}
	}

	return true
}

func FindPrevPrime(nb int) int {
	for i := nb; i >= 2; i-- {
		if isPrime(i) {
			return i
		}
	}

	return 0
}
