package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	var board [8]int
	solveQ(0, &board)
}

func solveQ(col int, board *[8]int) {
	if col == get8() { // if col == 8
		printBoard(board)
		return
	}

	for row := get1(); row <= get8(); row++ {
		if safe(col, row, board) {
			board[col] = row
			solveQ(col+get1(), board)
		}
	}
}

func safe(col, row int, board *[8]int) bool {
	for c := 0; c < col; c++ {
		r := board[c]

		// same row
		if r == row {
			return false
		}

		// diagonal check
		diffRow := r - row
		if diffRow < 0 {
			diffRow = -diffRow
		}

		diffCol := col - c
		if diffCol < 0 {
			diffCol = -diffCol
		}

		if diffRow == diffCol {
			return false
		}
	}
	return true
}

func printBoard(board *[8]int) {
	for i := 0; i < get8(); i++ {
		z01.PrintRune(rune(board[i] + toDigit()))
	}
	z01.PrintRune('\n')
}

// ---- Helpers that avoid writing digits ----

// returns rune '0'
func toDigit() int {
	return int('0')
}

// return 1..8 without numeric literals:
func get1() int { return int('1' - '0') }
func get8() int { return int('8' - '0') }
