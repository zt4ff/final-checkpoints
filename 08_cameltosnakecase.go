package piscine

func isCapitalLetter(s rune) bool {
	if s >= 'A' && s <= 'Z' {
		return true
	}

	return false
}

func CamelToSnakeCase(s string) string {
	result := ""

	for i, c := range s {
		if isCapitalLetter(c) {
			if i != 0 && isCapitalLetter(rune(s[i-1])) {
				return s
			}
			if i != 0 {
				result += "_"
			}
			result += string(c)
		} else {
			result += string(c)
		}
	}

	return result
}
