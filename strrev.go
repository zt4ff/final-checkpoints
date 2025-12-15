package piscine

func StrRev(s string) string {
	aString := []rune(s)
	var reverseSort string
	for i := len(aString) - 1; i >= 0; i-- {
		reverseSort += string(aString[i])
	}
	return reverseSort
}
