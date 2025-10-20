package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	upperCase := false
	startIndex := 0

	// Check for --upper flag
	if len(args) > 0 && args[0] == "--upper" {
		upperCase = true
		startIndex = 1
	}

	// Process each argument
	for i := startIndex; i < len(args); i++ {
		num := 0
		valid := true

		// Convert string to integer and validate
		for _, ch := range args[i] {
			if ch < '0' || ch > '9' {
				valid = false
				break
			}
			num = num*10 + int(ch-'0')
		}

		// Check if number is within alphabet range (1-26)
		if valid && num >= 1 && num <= 26 {
			if upperCase {
				z01.PrintRune(rune('A' + num - 1))
			} else {
				z01.PrintRune(rune('a' + num - 1))
			}
		} else {
			z01.PrintRune(' ')
		}
	}

	// Print newline at the end
	if len(args) > startIndex {
		z01.PrintRune('\n')
	}
}
