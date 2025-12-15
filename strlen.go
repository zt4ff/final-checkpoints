package piscine

func StrLen(s string) int {
	aString := []rune(s)
	length := 0
	for i := range aString {
		i++
		length++
	}
	return length
}
