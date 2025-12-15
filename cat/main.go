package main

import (
	"bufio"
	"io"
	"os"

	"github.com/01-edu/z01"
)

func printRune(r byte) {
	z01.PrintRune(rune(r))
}

func printFile(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		b, err := r.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		printRune(b)
	}
	return nil
}

func printStdin() {
	r := bufio.NewReader(os.Stdin)
	for {
		b, err := r.ReadByte()
		if err == io.EOF {
			break
		}
		printRune(b)
	}
}

func main() {
	if len(os.Args) == 1 {
		printStdin()
		return
	}

	for _, arg := range os.Args[1:] {
		err := printFile(arg)
		if err != nil {
			for _, c := range []rune("ERROR: " + err.Error() + "\n") {
				z01.PrintRune(c)
			}
			os.Exit(1)
		}
	}
}
