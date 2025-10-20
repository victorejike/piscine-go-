package piscine

func UltimateDivMod(a *int, b *int) {
	// Store the original values before modifying
	dividend := *a
	divisor := *b

	// Store division result in a
	*a = dividend / divisor

	// Store remainder in b
	*b = dividend % divisor
}
