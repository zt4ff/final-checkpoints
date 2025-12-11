package piscine

func SplitWords(s string) []string {
	words := []string{}
	word := ""

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if word != "" {
				words = append(words, word)
				word = ""
			}
		} else {
			word += string(s[i])
		}
	}

	// Add the last word if any
	if word != "" {
		words = append(words, word)
	}

	return words
}
