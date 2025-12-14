package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readFile(filename string) []string {
	var result []string

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
		result = append(result, line)
	}

	// Check for errors during the scan
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	return result
}

func findMax2Batteries(bank string) int {
	maxSum := 0
	for i := 0; i < len(bank)-1; i++ {
		first := bank[i]

		for j := i + 1; j < len(bank); j++ {
			second := bank[j]

			concatNumber := string([]byte{first, second})
			sum, err := strconv.Atoi(concatNumber)

			if err != nil {
				log.Fatalf("error converting string to int: %s", err)
			}

			if sum > maxSum {
				maxSum = sum
			}
		}
	}

	return maxSum
}

// Global best found so far
var globalBest int

func findMax12Batteries(bank string, usedBatteries string) int {
	if bank == "" && usedBatteries == "" {
		return 0
	}

	if bank == "" || len(usedBatteries) == 12 {
		currentSum, err := strconv.Atoi(usedBatteries)
		if err != nil {
			log.Fatalf("error converting string to int: %s", err)
		}
		if currentSum > globalBest {
			globalBest = currentSum
		}
		return currentSum
	}

	// Pruning: calculate the maximum possible value we could
	remainingSlots := 12 - len(usedBatteries)
	if remainingSlots > len(bank) {
		remainingSlots = len(bank)
	}

	if remainingSlots > 0 {
		// Build the best possible number: current digits + all 9s
		bestPossible := usedBatteries
		for i := 0; i < remainingSlots; i++ {
			bestPossible += "9"
		}
		maxPossible, err := strconv.Atoi(bestPossible)

		if err != nil {
			log.Fatalf("error converting string to int: %s", err)
		}

		if maxPossible <= globalBest {
			return 0 // This branch can't beat the best, prune it
		}
	}

	// Take the first battery
	firstBattery := bank[0]
	remainingBank := bank[1:]
	newUsedBatteries := usedBatteries + string(firstBattery)

	// Sum if we use the first battery
	sumWithBattery := 0

	// Sum if we don't use the first battery
	sumWithoutBattery := 0

	// Try taking the battery first if it's a high digit
	if firstBattery >= '5' {
		sumWithBattery = findMax12Batteries(remainingBank, newUsedBatteries)
		sumWithoutBattery = findMax12Batteries(remainingBank, usedBatteries)
	} else {
		sumWithoutBattery = findMax12Batteries(remainingBank, usedBatteries)
		sumWithBattery = findMax12Batteries(remainingBank, newUsedBatteries)
	}

	result := sumWithBattery
	if sumWithoutBattery > sumWithBattery {
		result = sumWithoutBattery
	}

	return result
}

func solveFirst(filename string) {
	fmt.Println("Solving first task with file: ", filename)

	banks := readFile(filename)

	joltageSum := 0
	for _, bank := range banks {
		joltageSum += findMax2Batteries(bank)

	}

	fmt.Printf("Max joltage sum is: %d", joltageSum)

	fmt.Println()
}

func solveSecond(filename string) {
	fmt.Println("Solving second task with file: ", filename)

	banks := readFile(filename)

	joltageSum := 0
	for _, bank := range banks {
		globalBest = 0 // Reset for each bank
		joltageSum += findMax12Batteries(bank, "")
	}

	fmt.Printf("Max joltage sum is: %d", joltageSum)
	fmt.Println()
}

func main() {

	solveFirst("input1.txt")
	solveFirst("input2.txt")
	fmt.Println()
	solveSecond("input1.txt")
	solveSecond("input2.txt")

}
