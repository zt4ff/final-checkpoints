package main

import (
	"fmt"
	"os"
)

func parsePositiveInt(s string) (int, bool) {
	num := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return 0, false
		}
		num = num*10 + int(c-'0')
	}
	if num <= 0 {
		return 0, false
	}
	return num, true
}

func main() {
	args := os.Args[1:]
	if len(args) < 2 || args[0] != "-c" {
		fmt.Println("Usage: go run . -c <number> file1 [file2 ...]")
		os.Exit(1)
	}

	count, ok := parsePositiveInt(args[1])
	if !ok {
		fmt.Println("Invalid count")
		os.Exit(1)
	}

	files := args[2:]
	hadError := false

	for i, f := range files {
		file, err := os.Open(f)
		if err != nil {
			fmt.Printf("open %s: no such file or directory\n", f)
			hadError = true
			continue
		}

		stat, _ := file.Stat()
		size := stat.Size()
		start := int64(0)
		if int64(count) < size {
			start = size - int64(count)
		}

		buf := make([]byte, size-start)
		_, _ = file.ReadAt(buf, start)

		if len(files) > 1 {
			if i > 0 {
				fmt.Println()
			}
			fmt.Printf("==> %s <==\n", f)
		}

		fmt.Print(string(buf))
		file.Close()
	}

	if hadError {
		os.Exit(1)
	}
}
