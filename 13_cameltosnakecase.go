package piscine

func isCap1(c rune) bool {
	if c >= 'A' && c <= 'Z' {
		return true
	}

	return false
}

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}

	ans := ""

	for i, c := range s {
		if isCap1(c) {
			if i+1 < len(s) && isCap1(rune(s[i+1])) {
				return s
			}
			if i != 0 {
				ans += "_"
			}
		}
		ans += string(c)
	}

	return ans
}
