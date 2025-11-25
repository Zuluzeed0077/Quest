package piscine

func Swap(a *int, b *int) {
	z := *a
	*a = *b
	*b = z
}
