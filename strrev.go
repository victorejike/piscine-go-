package piscine

func StrRev(s string) string {
	// Convert string to slice of runes to handle Unicode correctly
	runes := []rune(s)

	// Reverse the slice of runes
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Convert back to string and return
	return string(runes)
}
