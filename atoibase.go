package piscine

import "github.com/01-edu/z01"

func AtoiBase(s string, base string) int {
	// Validate base length
	if len(base) < 2 {
		return 0
	}

	charsSeen := make(map[rune]bool)
	for _, ch := range base {
		if ch == '+' || ch == '-' {
			return 0
		}
		if charsSeen[ch] {
			return 0
		}
		charsSeen[ch] = true
	}

	charToVal := make(map[rune]int)
	for i, ch := range base {
		charToVal[ch] = i
	}

	baseLen := len(base)
	result := 0

	for _, ch := range s {
		val, exists := charToVal[ch]
		if !exists {
			return 0
		}
		result = result*baseLen + val
	}

	return result
}

// A helper function to print an int using only z01.PrintRune
func PrintInt(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	var digits []rune
	for n > 0 {
		digits = append(digits, rune('0'+(n%10)))
		n /= 10
	}

	// digits are reversed, print in correct order
	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(digits[i])
	}
}

// Example main (if needed for your tests)
func main() {
	PrintInt(AtoiBase("125", "0123456789"))
	z01.PrintRune('\n')
	PrintInt(AtoiBase("1111101", "01"))
	z01.PrintRune('\n')
	PrintInt(AtoiBase("7D", "0123456789ABCDEF"))
	z01.PrintRune('\n')
	PrintInt(AtoiBase("uoi", "choumi"))
	z01.PrintRune('\n')
	PrintInt(AtoiBase("bbbbbab", "-ab"))
	z01.PrintRune('\n')
}
