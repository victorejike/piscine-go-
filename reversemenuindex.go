package piscine

func ReverseMenuIndex(menu []string) []string {
	length := len(menu)
	reversed := make([]string, length) // create a slice of the same size

	for i := 0; i < length; i++ {
		reversed[i] = menu[length-1-i]
	}

	return reversed
}
