package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Get the program name from os.Args[0]
	programName := os.Args[0]

	// Find the last occurrence of '/' to get just the executable name
	// (in case the program was run with a path like ./program or /path/to/program)
	name := programName
	for i := len(programName) - 1; i >= 0; i-- {
		if programName[i] == '/' {
			name = programName[i+1:]
			break
		}
	}

	// Print each character of the program name
	for _, char := range name {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
