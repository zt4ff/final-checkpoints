package piscine

func FirstWord(s string) string {
	var pos int

	for i, c := range s {
		if c == ' ' {
			pos = i
			break
		}
	}

	return s[:pos] + "\n"
}
