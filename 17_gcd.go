package piscine

func Gcd(a, b uint) uint {
	if a == 0 || b == 0 {
		return 0

	}

	if a < b {
		for i := a; i >= 1; i-- {
			if b%i == 0 && a%i == 0 {
				return i
			}
		}
	} else {
		for i := b; i >= 1; i-- {
			if a%i == 0 && b%i == 0 {
				return i
			}
		}
	}

	return 0
}
