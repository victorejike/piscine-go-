package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 || args[0] == "--help" || args[0] == "-h" {
		printHelp()
		return
	}

	var insertStr string
	var order bool
	var mainStr string

	// Parse flags
	for i := 0; i < len(args); i++ {
		arg := args[i]

		if len(arg) > 9 && arg[:9] == "--insert=" {
			insertStr = arg[9:]
		} else if len(arg) > 3 && arg[:3] == "-i=" {
			insertStr = arg[3:]
		} else if arg == "--order" || arg == "-o" {
			order = true
		} else if arg != "--insert" && arg != "-i" {
			// This is the main string (not a flag)
			mainStr = arg
		}
	}

	// If no main string found, use empty string
	if mainStr == "" {
		for i := 0; i < len(args); i++ {
			arg := args[i]
			if arg != "--insert" && arg != "-i" && arg != "--order" && arg != "-o" &&
				!(len(arg) > 9 && arg[:9] == "--insert=") && !(len(arg) > 3 && arg[:3] == "-i=") {
				mainStr = arg
				break
			}
		}
	}

	// Apply insert if specified
	if insertStr != "" {
		mainStr += insertStr
	}

	// Apply order if specified
	if order {
		mainStr = sortString(mainStr)
	}

	// Print the result
	for _, ch := range mainStr {
		z01.PrintRune(ch)
	}
	if mainStr != "" {
		z01.PrintRune('\n')
	}
}

func printHelp() {
	helpText := []string{
		"--insert",
		"  -i",
		"\t This flag inserts the string into the string passed as argument.",
		"--order",
		"  -o",
		"\t This flag will behave like a boolean, if it is called it will order the argument.",
	}

	for _, line := range helpText {
		for _, ch := range line {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
	}
}

func sortString(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		for j := i + 1; j < len(runes); j++ {
			if runes[i] > runes[j] {
				runes[i], runes[j] = runes[j], runes[i]
			}
		}
	}
	return string(runes)
}
