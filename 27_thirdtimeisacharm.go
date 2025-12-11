package piscine

func ThirdTimeIsACharm(str string) string {
	n := len(str)

	if n < 3 {
		return "\n"
	}

	result := ""

	for i, x := range str {
		if (i+1)%3 == 0 {
			result += string(x)
		}
	}

	return result + "\n"
}
