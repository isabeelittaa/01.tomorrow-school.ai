package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	le := max - min
	s := make([]int, le)
	for i := 0; i < le; i++ {
		s[i] = min + i
	}
	return s
}
