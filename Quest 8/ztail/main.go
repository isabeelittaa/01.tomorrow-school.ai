package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 || os.Args[1] != "-c" {
		fmt.Println("Usage: go run . -c <num> <file1> [file2...]")
		os.Exit(1)
	}

	numChars := 0
	fmt.Sscanf(os.Args[2], "%d", &numChars)

	if numChars <= 0 {
		fmt.Println("Invalid number of characters:", os.Args[2])
		os.Exit(1)
	}

	exitCode := 0

	for i, filename := range os.Args[3:] {

		file, err := os.Open(filename)
		if err != nil {
			// Prints error message without a preceding separator newline.
			fmt.Println(err)
			exitCode = 1
			continue
		}
		defer file.Close()

		// Fix 1: Print separator newline only for the second and subsequent
		// files that were successfully opened (i.e., not before an os.Open error).
		if i > 0 {
			fmt.Println()
		}

		fi, err := file.Stat()
		if err != nil {
			fmt.Printf("stat %s: %v\n", filename, err)
			exitCode = 1
			continue
		}

		fmt.Printf("==> %s <==\n", filename)

		fileSize := fi.Size()
		// Fix 2: Calculate offset to get the last 'numChars' bytes (not numChars-1)
		offset := fileSize - int64(numChars)
		if offset < 0 {
			offset = 0
		}

		_, err = file.Seek(offset, 0)
		if err != nil {
			fmt.Printf("seek %s: %v\n", filename, err)
			exitCode = 1
			continue
		}

		// Fix 2: Create buffer to hold 'numChars' bytes (not numChars-1)
		buf := make([]byte, numChars)
		n, err := file.Read(buf)
		if err != nil && err.Error() != "EOF" {
			fmt.Printf("read %s: %v\n", filename, err)
			exitCode = 1
			continue
		}

		fmt.Print(string(buf[:n]))

		// Trailing fmt.Println() was removed in the previous step and remains removed.
	}

	if exitCode != 0 {
		os.Exit(1)
	}
}
