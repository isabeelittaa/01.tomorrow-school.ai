package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[0]
	// Remove any directory path (keep only program name)
	start := 0
	for i := range name {
		if name[i] == '/' {
			start = i + 1
		}
	}
	for _, r := range name[start:] {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
