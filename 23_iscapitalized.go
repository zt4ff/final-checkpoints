package piscine

// loop through string

func isCapOrNum(s rune) bool {
	if s >= 'A' && s <= 'Z' || s >= '0' && s <= '9' {
		return true
	}
	return false
}

func IsCapitalized(s string) bool {

	if len(s) == 0 {
		return false
	}

	seenSpace := false

	for _, c := range s {
		if seenSpace && !isCapOrNum(c) {
			return false
		} else {
			seenSpace = false
		}

		if c == ' ' {
			seenSpace = true
		}
	}

	return true
}
