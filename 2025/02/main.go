package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile(filename string) string {

	b, err := os.ReadFile(filename) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'

	return str
}

func splitData(input string) [][]string {

	var ranges [][]string

	commas := strings.Split(input, ",")
	for _, stringRange := range commas {
		rangeSplit := strings.Split(stringRange, "-")
		ranges = append(ranges, rangeSplit)
	}

	return ranges
}

func filterData(data [][]string) [][]string {
	var filtered [][]string

	for _, pair := range data {
		start, end := pair[0], pair[1]
		if strings.HasPrefix(start, "0") || strings.HasPrefix(end, "0") {
			continue
		}
		filtered = append(filtered, pair)
	}

	return filtered
}

func checkIfValid(id string) bool {

	if len(id) < 2 {
		return true
	}

	if len(id)%2 != 0 {
		return true
	}

	firstHalf := id[0 : len(id)/2]
	secondHalf := id[len(id)/2:]

	if firstHalf == secondHalf {
		return false
	}

	return true

}

func countInvalidIds(data [][]string) int {
	invalidCount := 0
	sum := 0

	for _, pair := range data {
		start, err := strconv.Atoi(pair[0])

		if err != nil {
			fmt.Println("Error converting start:", err)
			continue
		}

		end, err := strconv.Atoi(pair[1])

		if err != nil {
			fmt.Println("Error converting end:", err)
			continue
		}

		for i := start; i <= end; i++ {
			id := strconv.Itoa(i)
			if !checkIfValid(id) {
				invalidCount++
				sum += i
			}
		}

	}

	fmt.Println("Total Invalid IDs: ", invalidCount)

	return sum

}

func checkIfValidV2(id string) bool {

	if len(id) < 2 {
		return true
	}

	for i := 2; i <= len(id); i++ {
		// Split the id in i parts and check if all parts are the same
		if len(id)%i != 0 {
			continue
		}

		partLength := len(id) / i
		part := id[0:partLength]
		allSame := true
		for j := 1; j < i; j++ {
			if id[j*partLength:(j+1)*partLength] != part {
				allSame = false
				break
			}
		}
		if allSame {
			return false
		}
	}

	return true

}

func countInvalidIdsV2(data [][]string) int {
	invalidCount := 0
	sum := 0

	for _, pair := range data {
		start, err := strconv.Atoi(pair[0])

		if err != nil {
			fmt.Println("Error converting start:", err)
			continue
		}

		end, err := strconv.Atoi(pair[1])

		if err != nil {
			fmt.Println("Error converting end:", err)
			continue
		}

		for i := start; i <= end; i++ {
			id := strconv.Itoa(i)
			if !checkIfValidV2(id) {
				invalidCount++
				sum += i
			}
		}

	}

	fmt.Println("Total Invalid IDs: ", invalidCount)

	return sum

}

func solveFirst(filename string) {
	fmt.Println("Solving first task with file: ", filename)

	input := readFile(filename)
	//fmt.Println("Input: ", input)

	split := splitData(input)
	//fmt.Println("Split Data: ", split)

	filtered := filterData(split)
	//fmt.Println("Filtered Data: ", filtered)

	sum := countInvalidIds(filtered)
	fmt.Println("Sum of Invalid IDs: ", sum)

	fmt.Println()
}

func solveSecond(filename string) {
	fmt.Println("Solving second task with file: ", filename)

	input := readFile(filename)
	//fmt.Println("Input: ", input)

	split := splitData(input)
	//fmt.Println("Split Data: ", split)

	filtered := filterData(split)
	//fmt.Println("Filtered Data: ", filtered)

	sum := countInvalidIdsV2(filtered)
	fmt.Println("Sum of Invalid IDs: ", sum)

	fmt.Println()
}

func main() {

	solveFirst("input1.txt")
	solveFirst("input2.txt")

	solveSecond("input1.txt")
	solveSecond("input2.txt")

}
