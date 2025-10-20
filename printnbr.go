package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	// Handle negative numbers
	if n < 0 {
		z01.PrintRune('-')
		// Handle MinInt specifically (Go's smallest int)
		if n == -9223372036854775808 {
			// For 64-bit environments, this value can't be negated directly.
			// We manually print it as its absolute value.
			str := "9223372036854775808"
			for _, r := range str {
				z01.PrintRune(r)
			}
			return
		}
		n = -n
	}

	// Collect digits in reverse order
	var digits []rune
	for n > 0 {
		digit := n % 10
		digits = append(digits, rune(digit)+'0')
		n = n / 10
	}

	// Print digits in correct order
	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(digits[i])
	}
}
