package piscine

func PrintIfNot(str string) string {
	count := 0

	for range str {
		count++
	}

	if count < 3 || count == 0 {
		return "G\n"
	}

	return "Invalid Input\n"
}
