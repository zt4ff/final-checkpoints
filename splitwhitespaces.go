package piscine

func SplitWhiteSpaces(s string) []string {
	var words []string
	length := len(s)
	start := -1

	for i := 0; i < length; i++ {
		if s[i] != ' ' && s[i] != '\t' && s[i] != '\n' {
			if start == -1 {
				start = i
			}
		} else {
			if start != -1 {
				words = append(words, s[start:i])
				start = -1
			}
		}
	}

	if start != -1 {
		words = append(words, s[start:length])
	}

	return words
}
