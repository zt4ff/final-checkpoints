package piscine

import (
	"fmt"
	"os"
)

func contains2(s string, c rune) bool {
	for _, x := range s {
		if x == c {
			return true
		}
	}
	return false
}

func Inter() {
	if len(os.Args) != 3 {
		return
	}

	dup := make(map[rune]int)

	s1 := (os.Args[1])
	s2 := (os.Args[2])

	ans := ""

	for _, x := range s1 {
		if contains2(s2, x) {
			if dup[x] <= 0 {
				ans += string(x)
			}
			dup[x]++
		}

	}

	fmt.Println(ans)

}
