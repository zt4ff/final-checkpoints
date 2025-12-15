package piscine

func NRune(s string, n int) rune {
	if n <= 0 {
		return 0
	}
	str := []rune(s)
	if n > len(str) {
		return 0
	}
	return str[n-1]
}
