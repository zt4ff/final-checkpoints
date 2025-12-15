package piscine

func Capitalize(s string) string {
	r := []rune(s)
	newWord := true

	for i, c := range r {

		// Check if c is alphanumeric
		isAlphaNum := (c >= 'a' && c <= 'z') ||
			(c >= 'A' && c <= 'Z') ||
			(c >= '0' && c <= '9')

		if isAlphaNum {
			if newWord {
				// Capitalize if letter
				if c >= 'a' && c <= 'z' {
					r[i] = c - ('a' - 'A')
				}
				newWord = false
			} else {
				// Lowercase if letter
				if c >= 'A' && c <= 'Z' {
					r[i] = c + ('a' - 'A')
				}
			}
		} else {
			newWord = true
		}
	}

	return string(r)
}
