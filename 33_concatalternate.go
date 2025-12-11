package piscine

func ConcatAlternate(slice1, slice2 []int) []int {
	l1 := len(slice1)
	l2 := len(slice2)

	if l1 == 0 {
		return slice2
	}

	if l2 == 0 {
		return slice1
	}

	var ans []int

	if l2 > l1 {
		for i := 0; i < l1; i++ {
			ans = append(ans, slice2[i])
			ans = append(ans, slice1[i])
		}
		ans = append(ans, slice2[l1:]...)
	} else {
		for i := 0; i < l2; i++ {
			ans = append(ans, slice1[i])
			ans = append(ans, slice2[i])
		}

		ans = append(ans, slice1[l2:]...)
	}

	return ans
}
