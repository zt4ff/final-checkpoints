package piscine

func ToUpper(s string) string {
	sRune := []rune(s)

	for i, x := range s {
		if x >= 'a' && x <= 'z' {
			sRune[i] = x - 32
		}
	}

	return string(sRune)
}
