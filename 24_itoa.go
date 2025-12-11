package piscine

func Itoa(n int) string {
	isNeg := false

	if n < 0 {
		isNeg = true
		n *= -1
	}

	result := ""

	for n > 0 {
		d := n % 10
		result = string(d+'0') + result
		n = n / 10
	}

	if isNeg {
		return "-" + result
	}

	return result
}
