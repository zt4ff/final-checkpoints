package piscine

func JumpOver(str string) string {
	result := ""
	for i, c := range str {
		if (i+1)%3 == 0 {
			result += string(c)
		}
	}

	return result + "\n"
}
