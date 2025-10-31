package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)
	start := 0
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			item := str[start:i]
			summary[item]++
			start = i + 1
		}
	}
	// Add the last item if there's any
	if start <= len(str) {
		item := str[start:]
		summary[item]++
	}
	return summary
}
