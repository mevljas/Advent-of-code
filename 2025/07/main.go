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

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, strings.Split(line, ""))
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	return result
}

func countBeams(grid [][]string, line int, nextLocations map[int]int) int {

	if line >= len(grid) {
		return 0
	}

	var newNextLocations = make(map[int]int)
	splitCount := 0
	currentLine := grid[line]
	lineLength := len(currentLine)

	for j := 0; j < lineLength; j++ {
		symbol := grid[line][j]
		hasLine := nextLocations[j] == 1

		if symbol == "S" {
			// Here laser beam starts
			newNextLocations[j] = 1
			break

		} else if hasLine {
			if symbol == "^" {
				// Here laser beam splits
				splitCount++

				if j > 0 {
					newNextLocations[j-1] = 1
				}
				if j < lineLength-1 {
					newNextLocations[j+1] = 1
				}

			} else {
				// Here laser beam continues
				newNextLocations[j] = 1
			}
		}
	}

	return countBeams(grid, line+1, newNextLocations) + splitCount
}

var memoizationTree [][]int

func countTimelines(grid [][]string, line int, currentIndex int) int {

	if line >= len(grid) {
		return 1
	}

	TimelinesCount := 0
	currentLine := grid[line]
	lineLength := len(currentLine)

	if currentIndex == -1 {
		// Find the starting point

		for j := 0; j < lineLength; j++ {
			symbol := grid[line][j]
			if symbol == "S" {
				return countTimelines(grid, line+1, j)
			}
		}
	}

	// Check memoization tree
	if memoizationTree[line][currentIndex] != -1 {
		return memoizationTree[line][currentIndex]
	}

	symbol := grid[line][currentIndex]

	if symbol == "^" {
		// Here laser beam splits

		if currentIndex > 0 {
			// left branch

			TimelinesCount += countTimelines(grid, line+1, currentIndex-1)
		}
		if currentIndex < lineLength-1 {
			// right branch

			TimelinesCount += countTimelines(grid, line+1, currentIndex+1)
		}

	} else {
		// Here laser beam continues

		TimelinesCount += countTimelines(grid, line+1, currentIndex)
	}

	// Save to memoization tree
	memoizationTree[line][currentIndex] = TimelinesCount

	return TimelinesCount
}

func solveFirst(filename string) {
	fmt.Println("Solving first task with file: ", filename)

	input := readFile(filename)

	splitCount := countBeams(input, 0, make(map[int]int))

	fmt.Println("Number of splits: ", splitCount)

	fmt.Println()
}

func initMemoizationTree(length int, width int) {
	memoizationTree = make([][]int, length)
	for i := 0; i < length; i++ {
		memoizationTree[i] = make([]int, width)
		for j := 0; j < width; j++ {
			memoizationTree[i][j] = -1
		}
	}
}

func solveSecond(filename string) {
	fmt.Println("Solving second task with file: ", filename)

	input := readFile(filename)

	inputLength := len(input)
	inputWidth := len(input[0])

	initMemoizationTree(inputLength, inputWidth)

	timelinesCount := countTimelines(input, 0, -1)

	fmt.Println("Number of timelines: ", timelinesCount)

	fmt.Println()
}

func main() {

	solveFirst("input1.txt")
	solveFirst("input2.txt")

	fmt.Println()

	solveSecond("input1.txt")
	solveSecond("input2.txt")

}
