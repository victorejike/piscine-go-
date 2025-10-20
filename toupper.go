package piscine

func ToUpper(s string) string {
	result := ""
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			r = r - 32
		}
		result += string(r)
	}
	return result
}
