package piscine

func RetainFirstHalf(str string) string {
	if str == "" || len(str) == 1 {
		return ""
	}

	return str[0 : len(str)/2]
}
