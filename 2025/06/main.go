package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
		result = append(result, strings.Fields(line))
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	return result
}

func readFileWithSpaces(filename string) []string {
	var result []string

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	return result
}

func findLongestInstruction(instructions []string) int {
	longest := 0

	for i := 0; i < len(instructions); i++ {
		if len(instructions[i]) > longest {
			longest = len(instructions[i])
		}
	}

	return longest
}

func fixInstructions(instructions []string) [][]string {

	var longestInstruction int = findLongestInstruction(instructions)

	for i := 0; i < longestInstruction; i++ {
		hasValueSomewhere := false

		// Find if there is any value in this column
		for j := 0; j < len(instructions); j++ {
			item := string(instructions[j][i])

			if item == "+" || item == "*" || item == " " {
				continue
			}
			hasValueSomewhere = true
			break
		}

		// Fill the column if necessary
		if hasValueSomewhere {
			for j := 0; j < len(instructions)-1; j++ {
				item := string(instructions[j][i])

				if item == " " {
					instructions[j] = instructions[j][:i] + "#" + instructions[j][i+1:]
				}
			}
		}

	}

	// Spit the instruction into 2D array
	var fixedInstructions [][]string
	for j := 0; j < len(instructions); j++ {
		fixedInstructions = append(fixedInstructions, strings.Fields(instructions[j]))
	}

	return fixedInstructions

}

func solveFirst(filename string) {
	fmt.Println("Solving first task with file: ", filename)

	sum := 0

	instructions := readFile(filename)

	numbersLen := len(instructions) - 1
	indexOfSymbols := numbersLen
	problemSize := len(instructions[0])

	for i := 0; i < problemSize; i++ {
		currentSymbol := instructions[indexOfSymbols][i]
		currentResult := 0

		if currentSymbol == "*" {
			currentResult = 1
		}

		for j := 0; j < numbersLen; j++ {
			number, err := strconv.Atoi(instructions[j][i])

			if err != nil {
				log.Fatalf("error converting string to int: %s", err)
			}

			if currentSymbol == "+" {
				currentResult += number
			} else if currentSymbol == "*" {
				currentResult *= number
			}

		}

		sum += currentResult

	}

	fmt.Println("Sum is: ", sum)

	fmt.Println()
}

func solveSecond(filename string) {
	fmt.Println("Solving second task with file: ", filename)

	sum := 0
	instructions := readFileWithSpaces(filename)
	fixedInstructions := fixInstructions(instructions)

	instructionsSize := len(fixedInstructions) - 1
	indexOfSymbols := instructionsSize
	numbersSize := len(fixedInstructions[0])

	for i := numbersSize - 1; i >= 0; i-- {
		currentSymbol := fixedInstructions[indexOfSymbols][i]
		columnWidth := len(fixedInstructions[0][i])
		columnSum := 0

		if currentSymbol == "*" {
			columnSum = 1
		}

		for j := 0; j < columnWidth; j++ {

			concatString := ""

			for k := instructionsSize - 1; k >= 0; k-- {
				item := string(fixedInstructions[k][i][j])

				if item == "#" {
					continue
				}

				concatString = item + concatString
			}

			number, err := strconv.Atoi(concatString)
			if err != nil {
				log.Fatalf("error converting string to int: %s", err)
			}

			if currentSymbol == "+" {
				columnSum += number
			} else if currentSymbol == "*" {
				columnSum *= number
			}
		}

		sum += columnSum

	}

	fmt.Println("Sum is: ", sum)

}

func main() {

	solveFirst("input1.txt")
	solveFirst("input2.txt")

	fmt.Println()

	solveSecond("input1.txt")
	solveSecond("input2.txt")

}
