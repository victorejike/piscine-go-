package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func printString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func main() {
	args := os.Args[1:]

	// If no arguments, read from stdin
	if len(args) == 0 {
		_, err := io.Copy(os.Stdout, os.Stdin)
		if err != nil {
			printString("ERROR: " + err.Error() + "\n")
			os.Exit(1)
		}
		return
	}

	// Process each file argument
	for _, filename := range args {
		file, err := os.Open(filename)
		if err != nil {
			printString("ERROR: " + err.Error() + "\n")
			os.Exit(1)
		}

		_, err = io.Copy(os.Stdout, file)
		if err != nil {
			printString("ERROR: " + err.Error() + "\n")
			file.Close()
			os.Exit(1)
		}
		file.Close()
	}
}
