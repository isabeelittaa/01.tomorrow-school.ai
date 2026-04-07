package piscine

// ForEach applies the function f to each element of the slice a
func ForEach(f func(int), a []int) {
	for _, v := range a {
		f(v) // Call the function f on each element v in the slice a
	}
}
