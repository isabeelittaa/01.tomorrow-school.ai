package piscine

func IterativeFactorial(nb int) int {
	num := nb
	fact := 1
	if num >= 0 {
		for iter := 1; iter <= num; iter++ {
			fact *= iter
		}
		return fact
	} else {
		return -1
	}
}
