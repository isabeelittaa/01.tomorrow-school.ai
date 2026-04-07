package main

import (
	"fmt"
	"os"
)

const SIZE = 9

func parseInput(input []string) ([][]int, error) {
	var board [][]int
	for _, line := range input {
		if len(line) != SIZE {
			return nil, fmt.Errorf("invalid input length")
		}
		var row []int
		for _, char := range line {
			if char == '.' {
				row = append(row, 0) // Empty cell
			} else if char >= '1' && char <= '9' {
				row = append(row, int(char-'0'))
			} else {
				return nil, fmt.Errorf("invalid character in input")
			}
		}
		board = append(board, row)
	}
	return board, nil
}

// Check if a number can be placed at the given row and column
func isValid(board [][]int, row, col, num int) bool {
	// Check the row
	for c := 0; c < SIZE; c++ {
		if board[row][c] == num {
			return false
		}
	}

	// Check the column
	for r := 0; r < SIZE; r++ {
		if board[r][col] == num {
			return false
		}
	}

	// Check the 3x3 subgrid
	startRow := (row / 3) * 3
	startCol := (col / 3) * 3
	for r := startRow; r < startRow+3; r++ {
		for c := startCol; c < startCol+3; c++ {
			if board[r][c] == num {
				return false
			}
		}
	}

	return true
}

// Backtracking function to solve the sudoku
func solve(board [][]int) bool {
	for row := 0; row < SIZE; row++ {
		for col := 0; col < SIZE; col++ {
			if board[row][col] == 0 { // If cell is empty
				for num := 1; num <= SIZE; num++ {
					if isValid(board, row, col, num) {
						board[row][col] = num
						if solve(board) {
							return true
						}
						board[row][col] = 0 // Backtrack
					}
				}
				return false // No valid number found, backtrack
			}
		}
	}
	return true // Sudoku solved
}

// Print the board in the required format
func printBoard(board [][]int) {
	for _, row := range board {
		for _, num := range row {
			fmt.Printf("%d ", num)
		}
		fmt.Println("$")
	}
}

func main() {
	if len(os.Args) < SIZE {
		fmt.Println("Error")
		return
	}

	// Parse the input
	board, err := parseInput(os.Args[1:])
	if err != nil {
		fmt.Println("Error")
		return
	}

	// Solve the Sudoku
	if solve(board) {
		printBoard(board)
	} else {
		fmt.Println("Error")
	}
}
