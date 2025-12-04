package piscine

func DigitLen(n, base int) int {
	if base >= 2 && base <= 36 {

		if n < 0 {
			n *= -1
		}

		count := 0

		for n > 0 {
			count++
			n = n / base
		}

		return count
	}
	return -1
}
