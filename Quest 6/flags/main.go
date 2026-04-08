package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 || hasHelp(args) {
		printHelp()
		return
	}

	insertStr := ""
	orderFlag := false
	mainStr := ""

	// Parse flags and arguments
	for _, a := range args {
		if a == "--order" || a == "-o" {
			orderFlag = true
		} else if len(a) > 9 && a[:9] == "--insert=" {
			insertStr = a[9:]
		} else if len(a) > 3 && a[:3] == "-i=" {
			insertStr = a[3:]
		} else {
			mainStr = a
		}
	}

	// Combine main string with insert flag (if any)
	result := mainStr + insertStr

	// Sort characters manually if -o or --order is given
	if orderFlag {
		result = manualSort(result)
	}

	// Print final result
	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

// Print help menu
func printHelp() {
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("\t This flag inserts the string into the string passed as argument.")
	fmt.Println("--order")
	fmt.Println("  -o")
	fmt.Println("\t This flag will behave like a boolean, if it is called it will order the argument.")
}

// Detect --help or -h
func hasHelp(args []string) bool {
	for _, a := range args {
		if a == "--help" || a == "-h" {
			return true
		}
	}
	return false
}

// Manual ASCII bubble sort
func manualSort(s string) string {
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if runes[j] > runes[j+1] {
				runes[j], runes[j+1] = runes[j+1], runes[j]
			}
		}
	}
	return string(runes)
}
