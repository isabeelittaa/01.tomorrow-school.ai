package piscine

func StringToIntSlice(str string) []int {
	var intSlice []int

	for _, r := range str {
		intSlice = append(intSlice, int(r))
	}

	return intSlice
}
