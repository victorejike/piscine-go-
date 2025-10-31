package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for a := 99; a >= 0; a-- {
		for b := a - 1; b >= 0; b-- {
			// Print the first two-digit number
			z01.PrintRune(rune(a/10 + '0'))
			z01.PrintRune(rune(a%10 + '0'))
			z01.PrintRune(' ')
			// Print the second two-digit number
			z01.PrintRune(rune(b/10 + '0'))
			z01.PrintRune(rune(b%10 + '0'))

			// If not the last combination, print comma and space
			if !(a == 1 && b == 0) {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
}
