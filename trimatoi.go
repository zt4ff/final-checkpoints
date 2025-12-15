package piscine

func TrimAtoi(s string) int {
	sign := 1
	number := 0
	foundDigit := false

	for _, r := range s {

		// Detect '-' only if sign not yet set by a number
		if r == '-' && !foundDigit {
			sign = -1
		}

		// If it's a digit, build the number
		if r >= '0' && r <= '9' {
			foundDigit = true
			number = number*10 + int(r-'0')
		}
	}

	if !foundDigit {
		return 0
	}

	return number * sign
}
