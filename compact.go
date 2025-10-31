package piscine

func Compact(ptr *[]string) int {
	slice := *ptr
	newSlice := make([]string, 0)

	for _, v := range slice {
		if v != "" { // Only keep non-empty strings
			newSlice = append(newSlice, v)
		}
	}

	*ptr = newSlice
	return len(newSlice)
}
