package piscine

func IsNumeric(s string) bool {
	for _, v := range s {
		if v >= '0' && v <= '9' {
		} else {
			return false
		}
	}
	return true
}
