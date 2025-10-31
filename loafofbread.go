package piscine

func LoafOfBread(str string) string {
	// If the input is empty or contains only spaces, return just a newline.
	if str == "" {
		return "\n"
	}

	runes := []rune(str)
	// Count non-space runes
	nonSpaceCount := 0
	for _, r := range runes {
		if r != ' ' {
			nonSpaceCount++
		}
	}
	// If there are zero non-space characters, treat like empty input.
	if nonSpaceCount == 0 {
		return "\n"
	}
	if nonSpaceCount < 5 {
		return "Invalid Output\n"
	}

	var chunks []string
	pos := 0
	L := len(runes)

	for pos < L {
		tmp := ""
		count := 0
		// collect up to 5 non-space runes
		for pos < L && count < 5 {
			if runes[pos] != ' ' {
				tmp += string(runes[pos])
				count++
			}
			pos++
		}
		if tmp != "" {
			chunks = append(chunks, tmp)
		}
		// skip the next character if any
		if pos < L {
			pos++
		}
	}

	// join chunks with spaces without importing packages
	result := ""
	for i, c := range chunks {
		if i > 0 {
			result += " "
		}
		result += c
	}
	result += "\n"
	return result
}
