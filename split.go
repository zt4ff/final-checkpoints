package piscine

func Split1(s, sep string) []string {
	var result []string
	sepLen := len(sep)
	sLen := len(s)

	if sepLen == 0 {
		return []string{s}
	}

	start := 0

	for i := 0; i <= sLen-sepLen; {
		if s[i:i+sepLen] == sep {
			result = append(result, s[start:i])
			i += sepLen
			start = i
		} else {
			i++
		}
	}

	result = append(result, s[start:sLen])

	return result
}
