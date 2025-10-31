package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	// Step 1: Convert nbr from baseFrom to decimal
	value := 0
	baseFromLen := len(baseFrom)

	for i := 0; i < len(nbr); i++ {
		value = value*baseFromLen + indexOf(baseFrom, nbr[i])
	}

	// Step 2: Convert decimal to baseTo
	if value == 0 {
		return string(baseTo[0])
	}

	baseToLen := len(baseTo)
	var result []byte

	for value > 0 {
		remainder := value % baseToLen
		result = append([]byte{baseTo[remainder]}, result...)
		value /= baseToLen
	}

	return string(result)
}

// Helper function to get the index of a rune in base string
func indexOf(base string, c byte) int {
	for i := 0; i < len(base); i++ {
		if base[i] == c {
			return i
		}
	}
	return -1
}
