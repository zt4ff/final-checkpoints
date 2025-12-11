package piscine

func Slice(a []string, nbrs ...int) []string {
	start := nbrs[0]
	end := len(a)

	if len(nbrs) == 2 {
		end = nbrs[1]
	}

	for start < 0 {
		start = len(a) + start
	}

	for end < 0 {
		end = len(a) + end
	}

	if start > end {
		return []string(nil)
	}

	return a[start:end]
}
