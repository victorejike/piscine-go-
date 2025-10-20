package piscine

func TrimAtoi(s string) int {
	result := 0
	negative := false
	foundDigit := false

	for _, c := range s {
		if c == '-' && !foundDigit {
			negative = true
		} else if c >= '0' && c <= '9' {
			digit := int(c - '0')
			result = result*10 + digit
			foundDigit = true
		}
	}

	if negative {
		return -result
	}
	return result
}
