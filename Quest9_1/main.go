package main

import (
	"fmt"

	piscine "piscinego"
)

// PrintNbr is the function passed to ForEach to print each number
func PrintNbr(n int) {
	fmt.Print(n)
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	piscine.ForEach(PrintNbr, a) // This will print: 123456
}
