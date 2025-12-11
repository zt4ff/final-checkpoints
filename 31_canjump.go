package piscine

func CanJump(arr []uint) bool {
	if len(arr) == 0 {
		return false
	}
	if len(arr) == 1 {
		return true
	}

	pos := 0
	for pos < len(arr)-1 {
		step := arr[pos]
		if step == 0 {
			return false
		}

		next := int(pos + int(step))
		if next >= len(arr) {
			return false
		}

		pos = next
	}

	return true
}
