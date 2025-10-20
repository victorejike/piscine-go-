package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for a1 := '0'; a1 <= '9'; a1++ {
		for a2 := '0'; a2 <= '9'; a2++ {
			for b1 := a1; b1 <= '9'; b1++ {
				var b2Start rune
				if b1 == a1 {
					b2Start = a2 + 1
				} else {
					b2Start = '0'
				}
				for b2 := b2Start; b2 <= '9'; b2++ {
					z01.PrintRune(a1)
					z01.PrintRune(a2)
					z01.PrintRune(' ')
					z01.PrintRune(b1)
					z01.PrintRune(b2)

					// Avoid comma and space after the last combination
					if !(a1 == '9' && a2 == '8' && b1 == '9' && b2 == '9') {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
