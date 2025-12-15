package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	value := atoibase(nbr, baseFrom)

	return fromBase(value, baseTo)
}

func atoibase(nbr string, base string) int {
	baseLen := len(base)
	result := 0

	for _, r := range nbr {
		result = result*baseLen + indexInBase(r, base)
	}

	return result
}

func indexInBase(r rune, base string) int {
	for i, br := range base {
		if br == r {
			return i
		}
	}
	return 0
}

func fromBase(n int, base string) string {
	if n == 0 {
		return string(base[0])
	}

	baseLen := len(base)
	result := ""

	for n > 0 {
		remainder := n % baseLen
		result = string(base[remainder]) + result
		n /= baseLen
	}

	return result
}
