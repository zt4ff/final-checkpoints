package piscine

func checkIfNumberOrDot(x rune) bool {
	if (x >= '0' && x <= '9') || x == '.' {
		return true
	}
	return false
}

func NotDecimal(dec string) string {
	if len(dec) == 0 {
		return "\n"
	}

	// if the first two strings are |0.| split them out
	if dec[0:2] == "0." {
		dec = dec[2:]
	}

	var isNeg bool

	if dec[0] == '-' {
		isNeg = true
	}

	ans := ""

	for _, x := range dec {
		if checkIfNumberOrDot(x) {

			if isNeg {
				ans += string(x)
			} else if x != '.' {
				ans += string(x)
			}
		} else {
			return dec
		}
	}

	return ans + "\n"
}
