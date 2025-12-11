package piscine

import (
	"fmt"
	"os"
)

func Options() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
		return
	}

	bits := uint32(0)
	help := false

	for _, arg := range args {
		if len(arg) == 0 || arg[0] != '-' {
			fmt.Println("Invalid Option")
			return
		}
		if arg == "-" {
			fmt.Println("Invalid Option")
			return
		}

		for i := 1; i < len(arg); i++ {
			c := arg[i]

			if c == 'h' {
				help = true
				break
			}

			if c < 'a' || c > 'z' {
				fmt.Println("Invalid Option")
				return
			}

			pos := c - 'a'
			bits |= 1 << pos
		}

		if help {
			fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
			return
		}
	}

	a := (bits >> 0) & 0xFF
	b := (bits >> 8) & 0xFF
	c := (bits >> 16) & 0xFF
	d := (bits >> 24) & 0xFF

	fmt.Printf("%08b %08b %08b %08b\n", d, c, b, a)
}
