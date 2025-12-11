package piscine

func contains(s string, c rune) bool {
	for _, x := range s {
		if x == c {
			return true
		}
	}
	return false
}

func WeAreUnique(str1, str2 string) int {

	if str1 == "" || str2 == "" {
		return -1
	}

	count := 0
	for _, x := range str1 {
		if !contains(str2, x) {
			count++
		}
	}

	for _, x := range str2 {
		if !contains(str1, x) {
			count++
		}
	}

	return count
}
