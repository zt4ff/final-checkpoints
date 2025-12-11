package piscine

func SaveAndMiss(arg string, num int) string {
	if num <= 0 {
		return arg
	}

	ans := ""
	save := true

	counter := 0

	for _, x := range arg {
		if save {
			ans += string(x)
		}

		counter++

		if counter == num {
			if save {
				save = false
			} else {
				save = true
			}
			counter = 0
		}
	}

	return ans
}
