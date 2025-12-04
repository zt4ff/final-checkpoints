package piscine

func CountChar(str string, c rune) int {
	count := 0
	for _, i := range str {
		if c == i {
			count++
		}
	}

	return count
}
