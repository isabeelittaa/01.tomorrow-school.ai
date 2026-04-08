package piscine

func RockAndRoll(n int) string {
	if n < 0 {
		return "error: number is negative\n"
	}

	isDivisibleBy2 := n%2 == 0
	isDivisibleBy3 := n%3 == 0

	if isDivisibleBy2 && isDivisibleBy3 {
		return "rock and roll\n"
	} else if isDivisibleBy2 {
		return "rock\n"
	} else if isDivisibleBy3 {
		return "roll\n"
	} else {
		return "error: non divisible\n"
	}
}
