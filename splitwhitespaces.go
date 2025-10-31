package piscine

func SplitWhiteSpaces(s string) []string {
	var words []string
	word := ""

	for i := 0; i < len(s); i++ {
		if s[i] != ' ' && s[i] != '\t' && s[i] != '\n' {
			word += string(s[i])
		} else if word != "" {
			words = append(words, word)
			word = ""
		}
	}

	// Add last word if there’s one
	if word != "" {
		words = append(words, word)
	}

	return words
}
