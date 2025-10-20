package piscine

import "github.com/01-edu/z01"

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}
	seen := make(map[rune]bool)
	for _, c := range base {
		if c == '+' || c == '-' || seen[c] {
			return false
		}
		seen[c] = true
	}
	return true
}

func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	baseLen := len(base)
	if baseLen < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	// Handle zero
	if nbr == 0 {
		z01.PrintRune(rune(base[0]))
		return
	}

	// Handle negative numbers safely using uint
	isNegative := false
	var unbr uint

	if nbr < 0 {
		isNegative = true
		unbr = uint(-int64(nbr)) // convert to int64 first to avoid overflow
	} else {
		unbr = uint(nbr)
	}

	if isNegative {
		z01.PrintRune('-')
	}

	// Convert to base
	var digits []rune
	for unbr > 0 {
		remainder := unbr % uint(baseLen)
		digits = append(digits, rune(base[remainder]))
		unbr /= uint(baseLen)
	}

	// Print in reverse order
	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(digits[i])
	}
}
