package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) < 2 {
		return true
	}

	ascending := true
	descending := true

	for i := 1; i < len(a); i++ {
		result := f(a[i-1], a[i])

		if result > 0 {
			ascending = false
		}
		if result < 0 {
			descending = false
		}

		if !ascending && !descending {
			return false
		}
	}

	return ascending || descending
}
