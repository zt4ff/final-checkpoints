package piscine

func isNull(a any) bool {
	if a == "" || a == false || a == 0 {
		return true
	}

	return false
}

func Compact(ptr *[]string) int {
	newArr := []string{}
	count := 0
	for _, c := range *ptr {
		if !isNull(c) {
			newArr = append(newArr, c)
			count++
		}
	}
	*ptr = newArr
	return count
}
