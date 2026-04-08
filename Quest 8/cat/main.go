package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 1 {
		_, err := io.Copy(os.Stdout, os.Stdin)
		if err != nil {
			for _, r := range []rune("ERROR: " + err.Error()) {
				z01.PrintRune(r)
			}
			z01.PrintRune('\n')
			os.Exit(1)
		}
		return
	}

	for _, fileName := range os.Args[1:] {
		file, err := os.Open(fileName)
		if err != nil {
			// Убираем излишнюю информацию
			for _, r := range []rune("ERROR: " + err.Error()) {
				z01.PrintRune(r)
			}
			z01.PrintRune('\n')
			os.Exit(1)
		}
		defer file.Close()

		_, err = io.Copy(os.Stdout, file)
		if err != nil {
			for _, r := range []rune("ERROR: " + err.Error()) {
				z01.PrintRune(r)
			}
			z01.PrintRune('\n')
			os.Exit(1)
		}
	}
}
