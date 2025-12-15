package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	var n food
	switch order {
	case "burger":
		n.preptime = 15
	case "chips":
		n.preptime = 10
	case "nuggets":
		n.preptime = 12
	default:
		return 404
	}
	return n.preptime
}
