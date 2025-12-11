package piscine

func ItoaBase(value, base int) string {
	d := "0123456789ABCDEF"

	if value == 0 {
		return "0"
	}

	neg := value < 0
	if neg {
		value = -value
	}

	out := ""

	for value > 0 {
		out = string(d[value%base]) + out
		value /= base
	}

	if neg {
		out = "-" + out
	}

	return out
}
