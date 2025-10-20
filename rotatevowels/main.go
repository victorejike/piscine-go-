package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		z01.PrintRune('\n')
		return
	}

	// Collect all vowels from all arguments
	var allVowels []rune
	vowelPositions := make([]struct {
		argIndex int
		strIndex int
	}, 0)

	// First pass: collect all vowels and their positions
	for argIndex, arg := range args {
		for strIndex, ch := range arg {
			if isVowel(ch) {
				allVowels = append(allVowels, ch)
				vowelPositions = append(vowelPositions, struct {
					argIndex int
					strIndex int
				}{argIndex, strIndex})
			}
		}
	}

	// If no vowels found, print arguments as is
	if len(allVowels) == 0 {
		for i, arg := range args {
			for _, ch := range arg {
				z01.PrintRune(ch)
			}
			if i < len(args)-1 {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
		return
	}

	// Rotate vowels (mirror positions)
	for i := 0; i < len(allVowels)/2; i++ {
		j := len(allVowels) - 1 - i
		allVowels[i], allVowels[j] = allVowels[j], allVowels[i]
	}

	// Second pass: reconstruct arguments with rotated vowels
	result := make([]string, len(args))
	for i, arg := range args {
		runes := []rune(arg)
		result[i] = string(runes)
	}

	// Replace vowels with rotated ones
	for i, pos := range vowelPositions {
		runes := []rune(result[pos.argIndex])
		runes[pos.strIndex] = allVowels[i]
		result[pos.argIndex] = string(runes)
	}

	// Print the result
	for i, str := range result {
		for _, ch := range str {
			z01.PrintRune(ch)
		}
		if i < len(result)-1 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}

func isVowel(ch rune) bool {
	switch ch {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}
