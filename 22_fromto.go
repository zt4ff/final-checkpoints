package piscine

func itoa(i int) string {
	result := ""

	for i > 0 {
		d := i % 10
		result = string('0'+d) + result
		i /= 10
	}

	return result
}

func isBiggerThan100orLessThan0(i int) bool {
	if i < 0 || i > 99 {
		return true
	}
	return false
}

func formatNumber(i int) string {
	if i < 10 {
		return "0" + itoa(i)
	}
	return itoa(i)
}

func FromTo(from int, to int) string {
	if isBiggerThan100orLessThan0(from) || isBiggerThan100orLessThan0(to) {
		return "Invalid\n"
	}
	result := ""

	if from < to {
		for i := from; i <= to; i++ {
			if i != from {
				result += ", "
			}
			result += formatNumber(i)
		}
	} else {
		for i := from; i >= to; i-- {
			if i != from {
				result += ", "
			}
			result += formatNumber(i)
		}
	}

	return result + "\n"
}
