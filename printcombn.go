package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}

	var comb [9]int

	var printComb func(pos int)
	printComb = func(pos int) {
		if pos == n {
			for i := 0; i < n; i++ {
				z01.PrintRune(rune('0' + comb[i]))
			}
			if comb[0] != 10-n {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			return
		}

		start := 0
		if pos > 0 {
			start = comb[pos-1] + 1
		}

		for d := start; d <= 9; d++ {
			comb[pos] = d
			printComb(pos + 1)
		}
	}

	printComb(0)
	z01.PrintRune('\n')
}
