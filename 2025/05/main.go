package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFile(filename string) ([][]int, []int) {
	var ranges [][]int
	var ingredients []int
	var passedBlankLine bool = false

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
		if line == "" {
			passedBlankLine = true
			continue
		}

		if passedBlankLine {
			newIngredient, err := strconv.Atoi(line)
			if err != nil {
				log.Fatalf("error converting string to int: %s", err)
			}

			ingredients = append(ingredients, newIngredient)
		} else {
			tempRange := strings.Split(line, "-")
			minRange, err := strconv.Atoi(tempRange[0])
			if err != nil {
				log.Fatalf("error converting string to int: %s", err)
			}

			maxRange, err := strconv.Atoi(tempRange[1])
			if err != nil {
				log.Fatalf("error converting string to int: %s", err)
			}

			ranges = append(ranges, []int{minRange, maxRange})
		}

	}

	// Check for errors during the scan
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	return ranges, ingredients
}

func isFresh(ingredient int, ranges [][]int) bool {
	for _, r := range ranges {
		minRange, maxRange := r[0], r[1]

		if ingredient >= minRange && ingredient <= maxRange {
			return true
		}
	}

	return false
}

func countFreshIngredients(ingredients []int, ranges [][]int) int {
	count := 0

	for _, ingredient := range ingredients {
		if isFresh(ingredient, ranges) {
			count++
		}
	}

	return count
}

func solveFirst(filename string) {
	fmt.Println("Solving first task with file: ", filename)

	ranges, ingredients := readFile(filename)

	freshCount := countFreshIngredients(ingredients, ranges)

	fmt.Println("Number of fresh ingredients: ", freshCount)

	fmt.Println()
}

func combineOverlappingRanges(ranges [][]int) [][]int {

	for i := 0; i < len(ranges); i++ {
		minRange1, maxRange1 := ranges[i][0], ranges[i][1]

		for j := i + 1; j < len(ranges); j++ {
			minRange2, maxRange2 := ranges[j][0], ranges[j][1]

			// Check for overlaps
			if maxRange1 >= minRange2 && minRange1 <= maxRange2 {
				newMin := minRange1
				if minRange2 < newMin {
					newMin = minRange2
				}

				newMax := maxRange1
				if maxRange2 > newMax {
					newMax = maxRange2
				}

				ranges[i] = []int{newMin, newMax}

				// Remove the j-th range as it has been merged
				ranges = append(ranges[:j], ranges[j+1:]...)
				i = 0 // Restart from the beginning
				break
			}
		}
	}

	return ranges

}

func countItemInRanges(ranges [][]int) int {
	count := 0

	for _, r := range ranges {
		minRange, maxRange := r[0], r[1]
		count += (maxRange - minRange + 1)
	}

	return count
}

func solveSecond(filename string) {
	fmt.Println("Solving second task with file: ", filename)

	ranges, _ := readFile(filename)

	combinedRanges := combineOverlappingRanges(ranges)

	countFreshItems := countItemInRanges(combinedRanges)

	fmt.Println("Number of fresh ingredients: ", countFreshItems)

	fmt.Println()
}

func main() {

	solveFirst("input1.txt")
	solveFirst("input2.txt")

	fmt.Println()

	solveSecond("input1.txt")
	solveSecond("input2.txt")

}
