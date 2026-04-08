package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	upper := false

	if len(args) > 0 && args[0] == "--upper" {
		upper = true
		args = args[1:]
	}

	// If there are no arguments left, produce no output at all.
	if len(args) == 0 {
		return
	}

	for _, arg := range args {
		n := Atoi(arg)
		if n < 1 || n > 26 {
			z01.PrintRune(' ')
			continue
		}
		var r rune
		if upper {
			r = rune('A' + n - 1)
		} else {
			r = rune('a' + n - 1)
		}
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

// Custom Atoi (no strconv)
func Atoi(s string) int {
	result := 0
	for _, r := range s {
		if r < '0' || r > '9' {
			return 0
		}
		result = result*10 + int(r-'0')
	}
	return result
}
