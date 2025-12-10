package piscine

func LastWord(s string) string {
	// Remove trailing spaces
	end := len(s) - 1

	if end == 0 {
		return "\n"
	}

	for s[end] == ' ' {
		end--
	}
	start := 0
	for s[start] == ' ' {
		start++
	}

	s = s[start : end+1]

	i := len(s) - 1

	for i >= 0 && s[i] != ' ' {
		i--
	}

	return s[i+1:] + "\n"
}
