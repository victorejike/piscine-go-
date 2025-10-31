package piscine

func ActiveBits(n int) int {
	if n < 0 {
		n = -n
	}
	count := 0
	for n != 0 {
		if n&1 == 1 {
			count++
		}
		n >>= 1
	}
	return count
}
