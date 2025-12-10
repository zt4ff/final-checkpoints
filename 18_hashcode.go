package piscine

func HashCode(dec string) string {
	ans := ""
	l := len(dec)
	for _, c := range dec {
		code := (int(c) + l) % 127
		for code > 127 {
			code %= 127
		}
		if code > 31 {
			code += 33
		}
		ans += string(((int(c) + l) % 127))
	}
	return ans
}
