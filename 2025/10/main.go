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
	// Create a copy 
	newMachineState := make([]int, len(machinesState))
	copy(newMachineState, machinesState)

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
	// Check if we found a solution
	if len(desiredMachinesState) == len(currentMachinesState) && reflect.DeepEqual(desiredMachinesState, currentMachinesState) {
		if minButtonsCombination == nil || len(currentlyPressedButtons) < len(minButtonsCombination) {
			// Make a copy of the solution
			solution := make([][]int, len(currentlyPressedButtons))
			copy(solution, currentlyPressedButtons)
			minButtonsCombination = solution
		}
		return
	}

	// Prune if we've already found a better solution
	if minButtonsCombination != nil && len(currentlyPressedButtons) >= len(minButtonsCombination) {
		return
	}

	if len(currentlyPressedButtons) >= maxDepth {
		return
	}

	for i := 0; i < len(buttons); i++ {
		button := buttons[i]
		newMachinesState := toggleButton(currentMachinesState, button)
		// Create a copy of currentlyPressedButtons to avoid shared state
		newPressedButtons := make([][]int, len(currentlyPressedButtons)+1)
		copy(newPressedButtons, currentlyPressedButtons)
		newPressedButtons[len(currentlyPressedButtons)] = button

		tryButtonsCombinations(desiredMachinesState, buttons, newMachinesState, newPressedButtons, maxDepth)
	}

	return
}

// Find the fewest combination of buttons that can satisfy the machine requirements
func findFewestButtonsCombination(desiredMachinesState []int, buttons [][]int) [][]int {
	minButtonsCombination = nil
	currentMachinesState := make([]int, len(desiredMachinesState))
	var currentlyPressedButtons [][]int

	// Try increasing depths
	for depth := 1; depth <= len(buttons)*2; depth++ {
		fmt.Println("Trying depth: ", depth)

		tryButtonsCombinations(desiredMachinesState, buttons, currentMachinesState, currentlyPressedButtons, depth)

		if minButtonsCombination != nil {
			return minButtonsCombination
		}
	}

	fmt.Println("No combination found")
	return nil
}

func solveFirst(filename string) {
	fmt.Println("Solving first task with file: ", filename)

	machines, buttons, _ := readFile(filename)

	fmt.Println("Machines: ", machines)
	fmt.Println("Buttons: ", buttons)

	totalButtonPresses := 0

	for i := 0; i < len(machines); i++ {
		desiredMachinesState := machines[i]
		buttonsForLine := buttons[i]

		result := findFewestButtonsCombination(desiredMachinesState, buttonsForLine)

		fmt.Println("Result: ", result)
		if result != nil {
			totalButtonPresses += len(result)
		}
	}

	fmt.Println("\nTotal button presses:", totalButtonPresses)

}

func main() {

	solveFirst("input1.txt")
	solveFirst("input2.txt")

	fmt.Println()

}
