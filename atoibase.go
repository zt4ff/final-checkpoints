package piscine

func AtoiBase(s string, base string) int {
	// Validate base
	if !isValidBase(base) {
		return 0
	}

	baseRunes := []rune(base)
	baseLen := len(baseRunes)

	// Build a mapping: rune → value
	// Example for base "choumi":
	// c→0, h→1, o→2, u→3, m→4, i→5
	indexMap := make(map[rune]int)
	for i, r := range baseRunes {
		indexMap[r] = i
	}

	result := 0

	// Convert the string
	for _, r := range s {
		value := indexMap[r] // guaranteed valid (tests say s is always valid)
		result = result*baseLen + value
	}

	return result
}

// Same base validation used in PrintNbrBase
func isValidBase(base string) bool {
	runes := []rune(base)

	// Must have at least 2 characters
	if len(runes) < 2 {
		return false
	}

	for i := 0; i < len(runes); i++ {
		if runes[i] == '+' || runes[i] == '-' {
			return false
		}
		for j := i + 1; j < len(runes); j++ {
			if runes[i] == runes[j] {
				return false
			}
		}
	}
	return true
}
