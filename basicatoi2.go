package piscine

func BasicAtoi2(s string) int {
	result := 0

	for _, char := range s {
		// Check if character is a digit
		if char < '0' || char > '9' {
			return 0
		}
		// Convert rune to digit by subtracting '0'
		digit := int(char - '0')
		// Shift current result left and add new digit
		result = result*10 + digit
	}

	return result
}
