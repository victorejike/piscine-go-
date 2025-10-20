package main

import (
	"github.com/01-edu/z01"
)

func main() {
	digit := '0'
	for digit <= '9' {
		z01.PrintRune(digit)
		digit++
	}
	z01.PrintRune('\n')
}
