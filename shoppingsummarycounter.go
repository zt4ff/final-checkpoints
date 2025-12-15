package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	v := make(map[string]int)
	var a string
	for _, k := range str {
		if k == 32 {
			v[a] += 1
			a = ""
		} else if k != 32 {
			a += string(byte(k))
		}
	}
	v[a] += 1

	return v
}
