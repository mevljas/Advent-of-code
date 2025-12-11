package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func readFile(filename string) ([][]int, [][][]int, [][]int) {
	var machines [][]int
	var buttons [][][]int
	var requirements [][]int

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		var lineButtons [][]int

		for i, part := range parts {
			partLen := len(part)
			if i == 0 {
				trimmed := part[1 : partLen-1]
				var lineMachines []int
				for character := range trimmed {
					if trimmed[character] == '#' {
						lineMachines = append(lineMachines, 1)
					} else {
						lineMachines = append(lineMachines, 0)
					}
				}
				machines = append(machines, lineMachines)
			} else if strings.HasPrefix(part, "(") {
				value := part[1 : partLen-1]
				numbers := strings.Split(value, ",")
				var button []int
				for _, numStr := range numbers {
					number, err := strconv.Atoi(numStr)
					if err != nil {
						log.Fatalf("failed to convert string to int: %s", err)
					}
					button = append(button, number)
				}

				lineButtons = append(lineButtons, button)

			} else if strings.HasPrefix(part, "{") {
				value := part[1 : partLen-1]
				numbers := strings.Split(value, ",")
				var lineRequirements []int
				for _, numStr := range numbers {
					number, err := strconv.Atoi(numStr)
					if err != nil {
						log.Fatalf("failed to convert string to int: %s", err)
					}
					lineRequirements = append(lineRequirements, number)
				}

				requirements = append(requirements, lineRequirements)
			}

		}

		buttons = append(buttons, lineButtons)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	return machines, buttons, requirements
}

var minButtonsCombination [][]int = nil

func toggleMachinesState(machinesState []int, machine int) []int {
	newMachineState := machinesState

	if newMachineState[machine] == 1 {
		newMachineState[machine] = 0
	} else {
		newMachineState[machine] = 1
	}

	return newMachineState
}

func toggleButton(machinesState []int, button []int) []int {
	newMachineState := machinesState
	for _, machine := range button {
		newMachineState = toggleMachinesState(newMachineState, machine)
	}

	return newMachineState
}

func tryButtonsCombinations(desiredMachinesState []int, buttons [][]int, currentMachinesState []int, currentlyPressedButtons [][]int, maxDepth int) {
	if minButtonsCombination == nil && len(desiredMachinesState) == len(currentMachinesState) && reflect.DeepEqual(desiredMachinesState, currentMachinesState) {
		minButtonsCombination = currentlyPressedButtons
		return
	}

	if len(currentlyPressedButtons) < len(minButtonsCombination) && len(desiredMachinesState) == len(currentMachinesState) && reflect.DeepEqual(desiredMachinesState, currentMachinesState) {
		minButtonsCombination = currentlyPressedButtons
		return
	}

	if len(currentlyPressedButtons) >= maxDepth {
		return
	}

	for i := 0; i < len(buttons); i++ {
		button := buttons[i]
		newMachinesState := toggleButton(currentMachinesState, button)
		newPressedButtons := append(currentlyPressedButtons, button)

		tryButtonsCombinations(desiredMachinesState, buttons, newMachinesState, newPressedButtons, maxDepth)

	}

	return
}

// Find the fewest combination of buttons that can satisfy the machine requirements
func findFewestButtonsCombination(desiredMachinesState []int, buttons [][]int) [][]int {

	maxDepth := len(buttons) + 1

	for true {

		fmt.Println("Trying depth: ", maxDepth)

		currentMachinesState := make([]int, len(desiredMachinesState))
		minButtonsCombination = nil
		var currentlyPressedButtons [][]int

		//	Pick a random button to press next
		for depth := 1; depth < maxDepth; depth++ {
			tryButtonsCombinations(desiredMachinesState, buttons, currentMachinesState, currentlyPressedButtons, depth)

			if minButtonsCombination != nil {
				return minButtonsCombination
			}

		}
		maxDepth *= 2
	}

	fmt.Println("No combination found")

	return nil

}

func solveFirst(filename string) {
	fmt.Println("Solving first task with file: ", filename)

	machines, buttons, _ := readFile(filename)

	fmt.Println("Machines: ", machines)
	fmt.Println("Buttons: ", buttons)

	for i := 0; i < len(machines); i++ {
		desiredMachinesState := machines[i]
		buttonsForLine := buttons[i]

		result := findFewestButtonsCombination(desiredMachinesState, buttonsForLine)

		fmt.Println("Result: ", result)
	}

	fmt.Println()

}

func main() {

	solveFirst("input1.txt")

	fmt.Println()

}
