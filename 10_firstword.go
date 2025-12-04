package piscine

func trimLeadingSpace(s string) string {
	if s == "" {
		return s
	}
	if s[0] != ' ' {
		return s
	}

	for s[0] == ' ' {
		s = s[1:]
	}

	return s
}

func FirstWord(s string) string {
	s = trimLeadingSpace(s)

	result := ""
	lastIndex := 0

	for i, c := range s {
		if c == ' ' {
			lastIndex = i + 1
			break
		}
	}

	result += s[0:lastIndex]

	return result + "\n"
}
