package piscine

func StringToIntSlice(str string) []int {
	res := []int(nil)

	for _, ch := range str {
		res = append(res, int(ch))
	}

	return res
}
