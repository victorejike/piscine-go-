package piscine

func Split(s, sep string) []string {
	var result []string
	word := ""
	sepLen := len(sep)

	for i := 0; i < len(s); {
		// Check if we found the separator
		if i+sepLen <= len(s) && s[i:i+sepLen] == sep {
			// Add the current word to the result
			result = append(result, word)
			word = ""
			i += sepLen // Skip over the separator
		} else {
			word += string(s[i])
			i++
		}
	}

	// Add the final word
	result = append(result, word)

	return result
}
