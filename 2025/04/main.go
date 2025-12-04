package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readFile(filename string) [][]string {
	var result [][]string

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	// Create a new scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Loop through the file and read each line
	for scanner.Scan() {
		line := scanner.Text() // Get the line as a string
		result = append(result, strings.Split(line, ""))
	}

	// Check for errors during the scan
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	return result
}

func isPositionAccessible(board [][]string, x int, y int) bool {

	boardHeight := len(board)
	boardWidth := len(board[0])

	adjacentRolls := 0

	// left
	if x > 0 && board[y][x-1] == "@" {
		adjacentRolls++
	}
	// right
	if x < boardWidth-1 && board[y][x+1] == "@" {
		adjacentRolls++
	}
	// up
	if y > 0 && board[y-1][x] == "@" {
		adjacentRolls++
	}
	// down
	if y < boardHeight-1 && board[y+1][x] == "@" {
		adjacentRolls++
	}
	// diagonal up-left
	if x > 0 && y > 0 && board[y-1][x-1] == "@" {
		adjacentRolls++
	}
	// diagonal up-right
	if x < boardWidth-1 && y > 0 && board[y-1][x+1] == "@" {
		adjacentRolls++
	}
	// diagonal down-left
	if x > 0 && y < boardHeight-1 && board[y+1][x-1] == "@" {
		adjacentRolls++
	}
	// diagonal down-right
	if x < boardWidth-1 && y < boardHeight-1 && board[y+1][x+1] == "@" {
		adjacentRolls++
	}

	return adjacentRolls < 4
}

func solveFirst(filename string) {
	fmt.Println("Solving first task with file: ", filename)

	board := readFile(filename)
	boardHeight := len(board)
	boardWidth := len(board[0])
	accessibleRolls := 0

	for i := 0; i < boardHeight; i++ {
		for j := 0; j < boardWidth; j++ {
			if board[i][j] == "@" {
				if isPositionAccessible(board, j, i) {
					accessibleRolls++
				}
			}
		}
	}

	fmt.Println("Number of accessible rolls: ", accessibleRolls)

	fmt.Println()
}

func solveSecond(filename string) {
	fmt.Println("Solving second task with file: ", filename)

	board := readFile(filename)
	boardHeight := len(board)
	boardWidth := len(board[0])
	removedRolls := 0
	hasRemovedRolls := true

	for loop := true; loop; loop = hasRemovedRolls {
		hasRemovedRolls = false
		for i := 0; i < boardHeight; i++ {
			for j := 0; j < boardWidth; j++ {
				if board[i][j] == "@" {
					if isPositionAccessible(board, j, i) {
						board[i][j] = "x"
						removedRolls++
						hasRemovedRolls = true
					}
				}

			}
		}
	}

	fmt.Println("Number of removed rolls: ", removedRolls)

	fmt.Println()
}

func main() {

	solveFirst("input1.txt")
	solveFirst("input2.txt")
	fmt.Println()
	solveSecond("input1.txt")
	solveSecond("input2.txt")

}
