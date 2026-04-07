package piscine

func ActiveBits(n int) int {
	count := 0
	if n < 0 {
		n = -n
	}

	for n > 0 {
		if n%2 == 1 {
			count++
		}
		n /= 2
	}
	return count
}
