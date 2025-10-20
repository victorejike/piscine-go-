package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	newWord := true

	for i, c := range runes {
		if isAlnum(c) {
			if newWord {
				// Capitalize letter
				if c >= 'a' && c <= 'z' {
					runes[i] = c - 'a' + 'A'
				}
				// digits stay the same
				newWord = false
			} else {
				// Lowercase letter
				if c >= 'A' && c <= 'Z' {
					runes[i] = c - 'A' + 'a'
				}
			}
		} else {
			// Non-alphanumeric resets newWord flag
			newWord = true
		}
	}
	return string(runes)
}

func isAlnum(c rune) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9')
}
