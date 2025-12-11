package piscine

func isAlphabet(c byte) bool {
	if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' {
		return true
	}

	return false
}

func itoa1(s int) string {
	result := ""

	for s > 0 {
		d := s % 10
		result = string('0'+d) + result
		s /= 10
	}

	return result
}

func ZipString(s string) string {
	ans := ""

	for i := 0; i < len(s); i++ {
		if isAlphabet(s[i]) {
			count := 1
			j := i + 1
			for j < len(s) && s[i] == s[j] {
				j++
				count++
			}
			ans += itoa1(count) + string(s[i])
			i = j - 1
		} else {
			ans += string(s[i])
		}
	}

	return ans
}
