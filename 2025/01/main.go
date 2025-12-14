package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const minDial = 0
const MaxDial = 99

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

func calculateOldPassword(data []string) {
	dial := 50
	zeroCounter := 0

	for _, line := range data {

		direction := string(line[0])
		number, err := strconv.Atoi(line[1:])

		if err != nil {
			log.Fatalf("error converting string to int: %s", err)
		}

		if direction == "R" {
			dial += number

			for dial > MaxDial {
				dial = dial - MaxDial - 1
			}

		} else if direction == "L" {
			dial -= number

			for dial < minDial {
				dial = MaxDial + dial + 1
			}
		} else {
			log.Fatalf("invalid direction: %s", direction)
		}

		if dial == 0 {
			zeroCounter += 1
		}

	}

	fmt.Println("Old password: ", zeroCounter)

}

func calculateNewPassword(data []string) {
	dial := 50
	zeroCounter := 0

	for _, line := range data {

		direction := string(line[0])
		number, err := strconv.Atoi(line[1:])

		if err != nil {
			log.Fatalf("error converting string to int: %s", err)
		}

		previousDial := dial
		zeroCounter += number / 100
		number = number % 100

		if direction == "R" {
			dial += number

		} else if direction == "L" {
			dial -= number

		}

		if dial > MaxDial {
			dial = dial - MaxDial - 1
			if previousDial != 0 {
				zeroCounter += 1
			}
		} else if dial < minDial {
			dial = MaxDial + dial + 1
			if previousDial != 0 {
				zeroCounter += 1
			}
		} else if dial == 0 {
			zeroCounter += 1
		}

	}

	fmt.Println("New password: ", zeroCounter)

}

func main() {

	input1 := readFile("input1.txt")
	input2 := readFile("input2.txt")
	input3 := readFile("input3.txt")

	fmt.Println("Input 1:")
	calculateOldPassword(input1)
	calculateNewPassword(input1)
	fmt.Println()

	fmt.Println("Input 2:")
	calculateOldPassword(input2)
	calculateNewPassword(input2)
	fmt.Println()

	fmt.Println("Input 3:")
	calculateOldPassword(input3)
	calculateNewPassword(input3)

}
