package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const minDial = 0
const MaxDial = 100

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

func calculateDialPosition(filename string) {
	data := readFile(filename)

	dial := 50

	counter := 0

	fmt.Println("Input : ", data)

	for _, line := range data {
		//fmt.Println(line)

		direction := string(line[0])
		number, err := strconv.Atoi(line[1:])

		if err != nil {
			log.Fatalf("error converting string to int: %s", err)
		}

		//fmt.Println(direction, number)

		if direction == "R" {
			dial += number

			for dial >= MaxDial {
				dial = dial - MaxDial
			}

		} else if direction == "L" {
			dial -= number

			for dial < minDial {
				dial = MaxDial + dial
			}
		} else {
			log.Fatalf("invalid direction: %s", direction)
		}

		if dial == 0 {
			counter += 1
		}

	}

	fmt.Println("Password: ", counter)

}

func main() {

	calculateDialPosition("input2.txt")

}
