package piscine

func PigLatin(s string) string {
	if len(s) == 0 {
		return ""
	}

	v := "aeiouAEIOU"
	first := s[0]

	for i := 0; i < len(v); i++ {
		if first == v[i] {
			return s + "ay"
		}
	}

	return s[1:] + string(first) + "ay"
}
