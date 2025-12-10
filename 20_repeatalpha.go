package piscine

func isLower(s rune) bool {
	return s >= 'a' && s <= 'z'
}

func isUpper(s rune) bool {
	return s >= 'A' && s <= 'Z'
}

func repeatChar(s rune, num int) string {
	ans := ""
	for i := 1; i <= num; i++ {
		ans += string(s)
	}
	return ans
}

func RepeatAlpha(s string) string {
	ans := ""
	for _, x := range s {
		if isUpper(x) {
			ans += repeatChar(x, int(x)-64)
		} else if isLower(x) {
			ans += repeatChar(x, int(x)-96)
		} else {
			ans += string(x)
		}
	}

	return ans
}
