package piscine

func ToLower(s string) string {
	sRune := []rune(s)

	for i, r := range sRune {
		if r >= 'A' && r <= 'Z' {
			sRune[i] = r + 32
		}
	}

	return string(sRune)
}
