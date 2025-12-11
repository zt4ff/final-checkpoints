package piscine

func FifthAndSkip(str string) string {
	if len(str) < 5 {
		return "\n"
	}

	ans := ""
	count := 0
	for _, x := range str {
		if count < 5 {
			if x != ' ' {
				ans += string(x)
				count++
			}
		} else {
			ans += " "
			count = 0
		}
	}

	return ans + "\n"
}
