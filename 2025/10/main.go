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

// Global variable to track the minimum button combination found
var minButtonsCombination [][]int = nil

func toggleMachinesState(machinesState []int, machine int) []int {
	// Create a copy to avoid modifying the original state
	newMachineState := make([]int, len(machinesState))
	copy(newMachineState, machinesState)

	if newMachineState[machine] == 1 {
		newMachineState[machine] = 0
	} else {
		newMachineState[machine] = 1
	}

	return newMachineState
}

// toggleButton applies a button press by toggling all affected indicator lights
func toggleButton(machinesState []int, button []int) []int {
	newMachineState := machinesState
	for _, machine := range button {
		newMachineState = toggleMachinesState(newMachineState, machine)
	}

	return newMachineState
}

// tryButtonsCombinations uses depth-first search to find button combinations
// that transform the machine from initial state (all lights off) to desired state.
// This is a brute-force approach with depth limiting and pruning.
func tryButtonsCombinations(desiredMachinesState []int, buttons [][]int, currentMachinesState []int, currentlyPressedButtons [][]int, maxDepth int) {
	// Base case: check if we found a valid solution
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

	// Depth limit reached, stop exploring this branch
	if len(currentlyPressedButtons) >= maxDepth {
		return
	}

	// Try pressing each button and recursively explore
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

// findFewestButtonsCombination finds the minimum number of button presses needed
// to configure the indicator lights to the desired state (Part 1 solution).
// Uses iterative deepening: tries depth 1, then 2, then 3, etc. until solution found.
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

// gcd computes the greatest common divisor using Euclidean algorithm.
// Used in Gaussian elimination to keep numbers manageable.
func gcd(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Matrix represents an augmented matrix for solving linear equations.
// For part 2, each row is an equation like: button0*x0 + button1*x1 + ... = requirement
// The last column holds the right-hand side (requirements).
type Matrix struct {
	rows int
	cols int
	data [][]int
}

func NewMatrix(rows, cols int) *Matrix {
	data := make([][]int, rows)
	for i := range data {
		data[i] = make([]int, cols)
	}
	return &Matrix{rows: rows, cols: cols, data: data}
}

func (m *Matrix) Copy() *Matrix {
	newM := NewMatrix(m.rows, m.cols)
	for i := 0; i < m.rows; i++ {
		copy(newM.data[i], m.data[i])
	}
	return newM
}

// GaussianElimination performs Gaussian elimination to row echelon form using integer arithmetic.
// Uses GCD-based row operations to avoid fractions and keep all values as integers.
// This transforms the matrix so we can identify free variables and solve the system.
func (m *Matrix) GaussianElimination() {
	pivotRow := 0

	// Process each column to create row echelon form
	for col := 0; col < m.cols-1 && pivotRow < m.rows; col++ {
		// Find a pivot: a row with non-zero entry in this column
		foundPivot := false
		for row := pivotRow; row < m.rows; row++ {
			if m.data[row][col] != 0 {
				// Swap rows
				m.data[pivotRow], m.data[row] = m.data[row], m.data[pivotRow]
				foundPivot = true
				break
			}
		}

		if !foundPivot {
			// No pivot in this column, it will be a free variable
			continue
		}

		// Eliminate this column in all other rows using row operations
		for row := 0; row < m.rows; row++ {
			if row == pivotRow {
				continue
			}

			if m.data[row][col] == 0 {
				// Already zero, no need to eliminate
				continue
			}

			// Use GCD to determine multipliers that keep numbers small and avoid fractions
			a := m.data[pivotRow][col]
			b := m.data[row][col]
			g := gcd(a, b)

			mulPivot := b / g
			mulRow := a / g

			// Row operation: row = row * mulRow - pivotRow * mulPivot
			// This eliminates the entry in column 'col' for this row
			for c := 0; c < m.cols; c++ {
				m.data[row][c] = m.data[row][c]*mulRow - m.data[pivotRow][c]*mulPivot
			}

			// Simplify the row by dividing by GCD if possible
			rowGcd := 0
			for c := 0; c < m.cols; c++ {
				if m.data[row][c] != 0 {
					if rowGcd == 0 {
						rowGcd = m.data[row][c]
						if rowGcd < 0 {
							rowGcd = -rowGcd
						}
					} else {
						rowGcd = gcd(rowGcd, m.data[row][c])
					}
				}
			}
			if rowGcd > 1 {
				for c := 0; c < m.cols; c++ {
					m.data[row][c] /= rowGcd
				}
			}
		}

		pivotRow++
	}
}

// GetFreeVariables identifies which variables are free (can be set arbitrarily).
// A free variable is one that doesn't have a leading 1 in its column after Gaussian elimination.
func (m *Matrix) GetFreeVariables() []int {
	freeVars := make([]bool, m.cols-1)
	for i := range freeVars {
		freeVars[i] = true
	}

	// Mark variables that are leading (not free)
	for row := 0; row < m.rows; row++ {
		// Find the first non-zero entry (leading variable) in this row
		for col := 0; col < m.cols-1; col++ {
			if m.data[row][col] != 0 {
				// Check if this is the only non-zero entry in this column (excluding other rows)
				onlyOne := true
				for r := 0; r < m.rows; r++ {
					if r != row && m.data[r][col] != 0 {
						onlyOne = false
						break
					}
				}
				if onlyOne {
					freeVars[col] = false
				}
				break
			}
		}
	}

	var result []int
	for i, isFree := range freeVars {
		if isFree {
			result = append(result, i)
		}
	}
	return result
}

// SolveWithFreeVars solves the linear system given specific values for free variables.
// Uses back-substitution to find values for all other variables.
func (m *Matrix) SolveWithFreeVars(freeVars []int, freeValues []int) []int {
	numVars := m.cols - 1
	solution := make([]int, numVars)

	// Initialize solution with the given free variable values
	freeVarMap := make(map[int]int)
	for i, varIdx := range freeVars {
		solution[varIdx] = freeValues[i]
		freeVarMap[varIdx] = freeValues[i]
	}

	// Solve for remaining (non-free) variables using back-substitution
	for row := m.rows - 1; row >= 0; row-- {
		// Find the leading variable
		leadingVar := -1
		for col := 0; col < numVars; col++ {
			if m.data[row][col] != 0 {
				leadingVar = col
				break
			}
		}

		// Empty row or already satisfied
		if leadingVar == -1 {
			continue
		}

		// Skip if this is a free variable (already set)
		if _, isFree := freeVarMap[leadingVar]; isFree {
			continue
		}

		// Calculate: leadingCoeff * x = rhs - sum(other terms)
		rhs := m.data[row][numVars]
		for col := 0; col < numVars; col++ {
			if col != leadingVar {
				rhs -= m.data[row][col] * solution[col]
			}
		}

		leadingCoeff := m.data[row][leadingVar]

		// Check if solution is integer (must divide evenly)
		if rhs%leadingCoeff != 0 {
			return nil
		}

		solution[leadingVar] = rhs / leadingCoeff
	}

	return solution
}

// findMinimalSolution finds the minimal number of button presses needed to satisfy
// the joltage requirements.
// Solves a system of linear equations using Gaussian elimination and searches over free variables.
func findMinimalSolution(buttons [][]int, requirements []int) int {
	if len(buttons) == 0 || len(requirements) == 0 {
		return 0
	}

	numButtons := len(buttons)
	numCounters := len(requirements)

	// Build the augmented matrix for the system of equations
	// Each row represents one counter: sum(button_presses[i] for buttons affecting this counter) = requirement
	// Each column represents one button variable (how many times to press it)
	// The last column is the right-hand side (the requirement value)
	matrix := NewMatrix(numCounters, numButtons+1)

	// Populate the matrix based on which buttons affect which counters
	for counter := 0; counter < numCounters; counter++ {
		for button := 0; button < numButtons; button++ {
			// Check if this button affects this counter (coefficient is 1 if yes, 0 if no)
			for _, affectedCounter := range buttons[button] {
				if affectedCounter == counter {
					matrix.data[counter][button] = 1
					break
				}
			}
		}
		matrix.data[counter][numButtons] = requirements[counter]
	}

	// Transform to row echelon form
	matrix.GaussianElimination()

	// Identify which button variables can be set freely
	freeVars := matrix.GetFreeVariables()

	if len(freeVars) == 0 {
		// No free variables: unique solution exists, solve directly
		solution := matrix.SolveWithFreeVars(nil, nil)
		if solution == nil {
			return -1
		}

		// Verify all button presses are non-negative (can't press negative times)
		for _, val := range solution {
			if val < 0 {
				return -1
			}
		}

		// Sum up total button presses
		total := 0
		for _, val := range solution {
			total += val
		}
		return total
	}

	// Multiple solutions exist: search over free variable space to find minimum
	// Set search bounds based on maximum requirement value
	maxVal := 0
	for _, req := range requirements {
		if req > maxVal {
			maxVal = req
		}
	}

	minPresses := -1

	// Recursive search function to try all combinations of free variable values
	var searchFreeVars func(index int, values []int)
	searchFreeVars = func(index int, values []int) {
		if index == len(freeVars) {
			// All free variables assigned, try solving with this combination
			solution := matrix.SolveWithFreeVars(freeVars, values)
			if solution == nil {
				return
			}

			// Check all non-negative
			valid := true
			total := 0
			for _, val := range solution {
				if val < 0 {
					valid = false
					break
				}
				total += val
			}

			if valid && (minPresses == -1 || total < minPresses) {
				minPresses = total
			}
			return
		}

		// Iterate through possible values for this free variable
		for val := 0; val <= maxVal; val++ {
			values[index] = val
			searchFreeVars(index+1, values)

			// Prune: if current value already exceeds best solution, no point trying higher
			if minPresses != -1 && val > minPresses {
				break
			}
		}
	}

	freeValues := make([]int, len(freeVars))
	searchFreeVars(0, freeValues)

	return minPresses
}

func solveSecond(filename string) {
	fmt.Println("Solving second task with file: ", filename)

	_, buttons, requirements := readFile(filename)

	totalButtonPresses := 0

	for i := 0; i < len(requirements); i++ {
		buttonsForLine := buttons[i]
		requirementsForLine := requirements[i]

		result := findMinimalSolution(buttonsForLine, requirementsForLine)

		if result > 0 {
			totalButtonPresses += result
		}
	}

	fmt.Println("Total button presses:", totalButtonPresses)
}

func main() {

	solveFirst("input1.txt")
	solveFirst("input2.txt")

	fmt.Println()

	solveSecond("input1.txt")
	solveSecond("input2.txt")

	fmt.Println()

}
