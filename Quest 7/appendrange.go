package piscine

func AppendRange(min, max int) []int {
	if min >= max {
		return nil
	}
	s := []int{}
	for i := min; i < max; i++ {
		s = append(s, i)
	}
	return s
}
