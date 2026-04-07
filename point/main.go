package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	setPoint(points)

	output := []rune{
		120, // x
		32,  // space
		61,  // =
		32,  // space
		52,  // 4
		50,  // 2
		44,  // ,
		32,  // space
		121, // y
		32,  // space
		61,  // =
		32,  // space
		50,  // 2
		49,  // 1
		10,  // newline
	}

	for _, c := range output {
		z01.PrintRune(c)
	}
}
