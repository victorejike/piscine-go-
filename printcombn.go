package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}

	// Initialize digits 0..n-1
	digits := make([]int, n)
	for i := 0; i < n; i++ {
		digits[i] = i
	}

	for {
		// Print current combination
		for _, d := range digits {
			z01.PrintRune(rune(d + '0'))
		}

		// Detect if this is the last combination (starts at 10-n)
		last := true
		for i := 0; i < n; i++ {
			if digits[i] != 10-n+i {
				last = false
				break
			}
		}
		if last {
			break
		}

		// Print separator
		z01.PrintRune(',')
		z01.PrintRune(' ')

		// Increment digits
		for i := n - 1; i >= 0; i-- {
			if digits[i] < 9-(n-1-i) {
				digits[i]++
				for j := i + 1; j < n; j++ {
					digits[j] = digits[j-1] + 1
				}
				break
			}
		}
	}

	z01.PrintRune('\n')
}
