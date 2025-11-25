package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	digits := "0123456789"

	if n == 0 {
		z01.PrintRune('0')
		return
	}

	if n < 0 {
		z01.PrintRune('-')
		if n <= -10 {
			PrintNbr(-(n / 10))
		}
		z01.PrintRune(rune(digits[-(n % 10)]))
		return
	}

	if n >= 10 {
		PrintNbr(n / 10)
	}

	z01.PrintRune(rune(digits[n%10]))
}
