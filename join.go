package piscine

func Join(strs []string, sep string) string {
	ans := ""
	for i, x := range strs {
		ans += x
		if i < len(strs)-1 {
			ans += sep
		}
	}

	return ans
}
