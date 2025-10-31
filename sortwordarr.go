package piscine

func SortWordArr(a []string) {
	n := len(a)
	if n <= 1 {
		return
	}

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if a[j] > a[j+1] {
				// Swap elements
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}
