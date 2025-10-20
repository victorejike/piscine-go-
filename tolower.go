package piscine

func ToLower(s string) string {
	result := ""
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			r = r + 32
		}
		result += string(r)
	}
	return result
}
