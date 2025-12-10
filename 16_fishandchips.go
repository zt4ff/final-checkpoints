package piscine

func FishAndChips(n int) string {
	isDivisibleBy2 := n%2 == 0
	isDivisibleBy3 := n%3 == 0

	if n < 0 {
		return "error: number is negative"
	}

	if isDivisibleBy2 && isDivisibleBy3 {
		return "fish and chips"
	}

	if isDivisibleBy2 {
		return "fish"
	}

	if isDivisibleBy3 {
		return "chips"
	}

	return "error: non divisible"
}
