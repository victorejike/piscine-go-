package piscine

import "github.com/01-edu/z01"

// Print the solution for the 8 queens puzzle
func printSolution(solution []int) {
	for _, val := range solution {
		z01.PrintRune(rune(val + '0')) // Print the column as a rune
	}
	z01.PrintRune('\n') // Print newline after each solution
}

// Check if it is safe to place the queen at position (row, col)
func isSafe(row, col int, columns [8]bool, leftDiag [15]bool, rightDiag [15]bool) bool {
	if columns[col] || leftDiag[row-col+7] || rightDiag[row+col] {
		return false
	}
	return true
}

// Recursively try to place queens on the chessboard
func placeQueens(row int, solution []int, columns [8]bool, leftDiag [15]bool, rightDiag [15]bool) {
	if row == 8 { // If all 8 queens are placed, print the solution
		printSolution(solution)
		return
	}

	for col := 0; col < 8; col++ {
		if isSafe(row, col, columns, leftDiag, rightDiag) {
			// Place the queen and mark the column and diagonals
			columns[col] = true
			leftDiag[row-col+7] = true
			rightDiag[row+col] = true
			solution[row] = col + 1 // Store the solution (1-indexed)

			// Recursively try to place queens in the next row
			placeQueens(row+1, solution, columns, leftDiag, rightDiag)

			// Backtrack: Unmark the column and diagonals
			columns[col] = false
			leftDiag[row-col+7] = false
			rightDiag[row+col] = false
		}
	}
}

// The main function to solve the 8 queens puzzle
func EightQueens() {
	// Use fixed-size arrays instead of `make`
	var columns [8]bool    // To track columns where queens are placed
	var leftDiag [15]bool  // To track left diagonals (r - c + 7)
	var rightDiag [15]bool // To track right diagonals (r + c)
	var solution [8]int    // To store the solution (1-indexed)

	// Start solving from the first row
	placeQueens(0, solution[:], columns, leftDiag, rightDiag)
}
