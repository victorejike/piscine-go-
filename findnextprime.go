package piscine

func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	}
	for {
		if isPrime(nb) {
			return nb
		}
		nb++
	}
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	// Check divisibility up to √n, skipping even numbers
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}
