package piscine

func sortWordArr(a []int) {
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}

func Abort(a, b, c, d, e int) int {
	arr := []int{a, b, c, d, e}

	sortWordArr(arr)

	return arr[2]
}
