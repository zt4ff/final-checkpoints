package piscine

func LastRune(s string) rune {
	str := []rune(s)
	strLen := len(str)
	return str[strLen-1]
}
