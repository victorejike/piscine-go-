package main

import (
	"fmt"
	"os"
)

func main() {
	// Check the number of arguments
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("File name missing")
		return
	}

	if len(args) > 1 {
		fmt.Println("Too many arguments")
		return
	}

	// Read and display the file content
	filename := args[0]
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Print the file content
	fmt.Print(string(content))
}
