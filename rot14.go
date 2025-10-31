package piscine

func Rot14(s string) string {
	result := []rune{}

	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			r = ((r-'a'+14)%26 + 'a')
		} else if r >= 'A' && r <= 'Z' {
			r = ((r-'A'+14)%26 + 'A')
		}
		result = append(result, r)
	}

	return string(result)
}
