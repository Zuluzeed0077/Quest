package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	var board [8]int

	var abs func(int) int
	abs = func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}

	var printSolution func()
	printSolution = func() {
		for i := 0; i < 8; i++ {
			z01.PrintRune(rune(board[i] + '0'))
		}
		z01.PrintRune('\n')
	}

	var isSafe func(int, int) bool
	isSafe = func(col, row int) bool {
		for c := 0; c < col; c++ {
			if board[c] == row || abs(board[c]-row) == abs(c-col) {
				return false
			}
		}
		return true
	}

	var solve func(int)
	solve = func(col int) {
		if col == 8 {
			printSolution()
			return
		}
		for row := 1; row <= 8; row++ {
			if isSafe(col, row) {
				board[col] = row
				solve(col + 1)
			}
		}
	}

	solve(0)
}
