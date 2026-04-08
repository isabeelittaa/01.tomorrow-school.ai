package piscine

func Abort(a, b, c, d, e int) int {
	lis := []int{a, b, c, d, e}
	for i := 0; i < len(lis)-1; i++ {
		for j := i + 1; j < len(lis); j++ {
			if lis[i] > lis[j] {
				lis[i], lis[j] = lis[j], lis[i]
			}
		}
	}

	return lis[2]
}
