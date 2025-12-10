package piscine

func CountChar(str string, c rune) int {
	count := 0
	for _, i := range str {
		if i == c {
			count++
		}
	}
	return count
}
