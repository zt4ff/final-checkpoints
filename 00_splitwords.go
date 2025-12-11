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

func Split(s, sep string) []string {
	out := []string{}
	curr := ""
	n := len(sep)

	for i := 0; i < len(s); {
		if i+n <= len(s) && s[i:i+n] == sep {
			out = append(out, curr)
			curr = ""
			i += n
		} else {
			curr += string(s[i])
			i++
		}
	}
	out = append(out, curr)

	return out
}
