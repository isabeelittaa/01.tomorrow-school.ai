package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	n := len(a)
	if n <= 1 {
		return true
	}

	hasAscendingViolation := false
	hasDescendingViolation := false

	for i := 0; i < n-1; i++ {
		comparison := f(a[i], a[i+1])

		if comparison > 0 {
			hasAscendingViolation = true
		} else if comparison < 0 {
			hasDescendingViolation = true
		}

		if hasAscendingViolation && hasDescendingViolation {
			return false
		}
	}
	return true
}
