package piscine

func RevConcatAlternate(slice1, slice2 []int) []int {
	reverseA := []int{}
	reverseB := []int{}
	for i := len(slice1) - 1; i >= 0; i-- {
		reverseA = append(reverseA, slice1[i])
	}
	for i := len(slice2) - 1; i >= 0; i-- {
		reverseB = append(reverseB, slice2[i])
	}

	ans := []int{}
	i, j := 0, 0

	if len(reverseA) > len(reverseB) {
		for ; i < len(reverseA)-len(reverseB); i++ {
			ans = append(ans, reverseA[i])
		}
	} else if len(reverseB) > len(reverseA) {
		for ; j < len(reverseB)-len(reverseA); j++ {
			ans = append(ans, reverseB[j])
		}
	}

	for i < len(reverseA) && j < len(reverseB) {
		ans = append(ans, reverseA[i])
		ans = append(ans, reverseB[j])
		i++
		j++
	}

	return ans
}
