package piscine

func Index(s string, toFind string) int {
	lenToFind := len(toFind)
	for i := 0; i < len(s); i++ {
		if s[i] == toFind[0] {
			if len(s[i:]) >= lenToFind && Compare(s[i:i+lenToFind], toFind) == 0 {
				return i
			}
		}
	}
	return -1
}
