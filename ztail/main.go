package main

import (
	"fmt"
	"os"
)

func atoi(s string) int {
	result := 0
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return -1
		}
		result = result*10 + int(ch-'0')
	}
	return result
}

func main() {
	args := os.Args[1:]

	// Check minimum arguments
	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s -c <number> <file1> [file2...]\n", os.Args[0])
		os.Exit(1)
	}

	// Validate -c option
	if args[0] != "-c" {
		fmt.Fprintf(os.Stderr, "Error: first argument must be -c\n")
		os.Exit(1)
	}

	// Parse the number of bytes using custom atoi
	n := atoi(args[1])
	if n < 0 {
		fmt.Fprintf(os.Stderr, "Error: invalid number of bytes\n")
		os.Exit(1)
	}

	files := args[2:]
	if len(files) == 0 {
		fmt.Fprintf(os.Stderr, "Error: at least one file required\n")
		os.Exit(1)
	}

	hasError := false
	printedAny := false
	hadPreviousError := false

	// Process each file
	for _, filename := range files {
		// Read file content first to check if it exists
		content, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "open %s: no such file or directory\n", filename)
			hasError = true
			hadPreviousError = true
			continue
		}

		// Print separator for multiple files (only if we've printed something before)
		if printedAny {
			fmt.Println()
		} else if hadPreviousError {
			// If we had errors before but this is the first successful file, add a newline
			fmt.Println()
		}

		// Print header for multiple files
		if len(files) > 1 {
			fmt.Printf("==> %s <==\n", filename)
		}

		// Handle the tail -c logic
		if n > len(content) {
			// If requested bytes more than file size, print entire file
			fmt.Print(string(content))
		} else {
			// Print last n bytes
			fmt.Print(string(content[len(content)-n:]))
		}

		printedAny = true
		hadPreviousError = false
	}

	if hasError {
		os.Exit(1)
	}
}
