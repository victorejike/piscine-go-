package piscine

func BasicAtoi(s string) int {
	result := 0

	for _, char := range s {
		// Convert rune to digit by subtracting '0'
		digit := int(char - '0')
		// Shift current result left and add new digit
		result = result*10 + digit
	}

	return result
}
