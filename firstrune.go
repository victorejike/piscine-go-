package piscine

func FirstRune(s string) rune {
	for _, r := range s {
		return r // return the first rune and exit the loop
	}
	return 0 // in case the string is empty
}
