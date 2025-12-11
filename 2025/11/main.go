package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/dominikbraun/graph"
)

// readFile parses the input file and builds an adjacency list representation of the graph.
func readFile(filename string) map[string][]string {
	connections := make(map[string][]string)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		first := parts[0]
		key := first[0 : len(first)-1]
		connections[key] = parts[1:]

	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	return connections
}

// topologicalSort performs Kahn's algorithm to order nodes so that
// for every edge A -> B, A comes before B in the ordering.
// Returns nil if the graph contains a cycle (not a DAG).
//
// Algorithm:
// 1. Count incoming edges (in-degree) for each node
// 2. Start with nodes that have no incoming edges
// 3. Remove each node and decrement in-degree of its neighbors
// 4. When a neighbor's in-degree becomes 0, add it to the queue
func topologicalSort(connections map[string][]string) []string {
	// Collect all nodes and count their incoming edges
	nodes := make(map[string]bool)
	inDegree := make(map[string]int)

	for node, neighbors := range connections {
		nodes[node] = true
		for _, neighbor := range neighbors {
			nodes[neighbor] = true
			inDegree[neighbor]++ // neighbor has one more incoming edge
		}
	}

	// Start with nodes that have no incoming edges (in-degree = 0)
	var queue []string
	for node := range nodes {
		if inDegree[node] == 0 {
			queue = append(queue, node)
		}
	}

	var result []string
	for len(queue) > 0 {
		// Take next node with no remaining incoming edges
		current := queue[0]
		queue = queue[1:]
		result = append(result, current)

		// "Remove" this node by decrementing in-degree of its neighbors
		for _, neighbor := range connections[current] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// If we couldn't process all nodes, there's a cycle
	if len(result) != len(nodes) {
		return nil // Cycle detected - not a DAG
	}

	return result
}

// State represents the DP state: current node and which checkpoints have been visited.
// hasFft: true if we've already passed through the 'fft' node
// hasDac: true if we've already passed through the 'dac' node
type State struct {
	node   string
	hasFft bool
	hasDac bool
}

// countPathsDAG uses dynamic programming on a DAG to count valid paths.
// This is the key optimization: instead of enumerating all paths (exponential),
// we compute the count using DP (linear in nodes × states).
//
// The idea:
//   - Process nodes in reverse topological order (from 'out' back to 'svr')
//   - For each node and each state (hasFft, hasDac), compute how many valid paths
//     exist from that node to 'out'
//   - A path is valid if it passes through 'fft' before 'dac'
//
// Time complexity: O(V × 4) where V = number of nodes, 4 = number of (hasFft, hasDac) combinations
func countPathsDAG(connections map[string][]string, start, target string, topologicalOrder []string) int {
	// Get topological ordering of nodes
	topoOrder := topologicalSort(connections)
	if topoOrder == nil {
		fmt.Println("Graph has a cycle - cannot use DP approach")
		return -1
	}

	// Reverse the order so we process from target ('out') backwards to start ('svr')
	// This way, when we compute dp[node], all dp[neighbor] values are already computed
	for i, j := 0, len(topoOrder)-1; i < j; i, j = i+1, j-1 {
		topoOrder[i], topoOrder[j] = topoOrder[j], topoOrder[i]
	}

	// dp[state] = number of valid paths from state.node to target,
	//             given that we've already visited fft (if hasFft) and dac (if hasDac)
	dp := make(map[State]int)

	// Base case: we're at the target node
	// A path is only valid if we've visited BOTH fft and dac by the time we reach target
	dp[State{target, true, true}] = 1   // valid: has both fft and dac
	dp[State{target, true, false}] = 0  // invalid: missing dac
	dp[State{target, false, true}] = 0  // invalid: missing fft
	dp[State{target, false, false}] = 0 // invalid: missing both

	// Process nodes in reverse topological order (from target back to start)
	for _, node := range topoOrder {
		if node == target {
			continue // Already handled as base case
		}

		// For each combination of (hasFft, hasDac) flags
		for _, fft := range []bool{false, true} {
			for _, dac := range []bool{false, true} {
				state := State{node, fft, dac}

				// After visiting this node, update the flags
				// If this node is 'fft', set hasFft = true
				// If this node is 'dac', set hasDac = true
				newFft := fft || (node == "fft")
				newDac := dac || (node == "dac")

				// KEY CONSTRAINT: if we're visiting 'dac', 'fft' must already be visited!
				// This enforces the "fft before dac" requirement
				if node == "dac" && !fft {
					dp[state] = 0 // Invalid: visiting dac without having visited fft
					continue
				}

				// Sum up paths through all neighbors
				// dp[current] = sum of dp[neighbor] for all outgoing edges
				total := 0
				for _, neighbor := range connections[node] {
					nextState := State{neighbor, newFft, newDac}
					total += dp[nextState]
				}
				dp[state] = total
			}
		}
	}

	// Answer: count paths from start node, starting with hasFft=false, hasDac=false
	// (we haven't visited either checkpoint yet when we begin)
	return dp[State{start, false, false}]
}

// findAllPaths uses DFS to enumerate all paths from current position to target.
// WARNING: This is exponential in complexity - only use for small graphs!
// For large graphs, use countPathsDAG instead.
func findAllPaths(connections map[string][]string, currentPath []string, visitedDevices map[string]bool, targetDevice string) [][]string {

	allPaths := [][]string{}
	currentDevice := currentPath[len(currentPath)-1]
	visitedDevicesCopy := make(map[string]bool)

	for k, v := range visitedDevices {
		visitedDevicesCopy[k] = v
	}
	visitedDevicesCopy[currentDevice] = true

	if targetDevice == currentDevice {
		allPaths = append(allPaths, currentPath)
		return allPaths
	}

	for _, nextDevice := range connections[currentDevice] {

		if !visitedDevices[nextDevice] {

			newPath := append(currentPath, nextDevice)
			pathsFromNext := findAllPaths(connections, newPath, visitedDevicesCopy, targetDevice)
			allPaths = append(allPaths, pathsFromNext...)
		}
	}

	return allPaths

}

// solveFirst solves part 1: count all paths from 'you' to 'out'
// Uses simple DFS path enumeration (works for small inputs)
func solveFirst(filename string) {
	fmt.Println("Solving first task with file: ", filename)

	connections := readFile(filename)

	startDevice := "you"
	targetDevice := "out"

	visitedDevices := make(map[string]bool)
	initialPath := []string{startDevice}

	allPaths := findAllPaths(connections, initialPath, visitedDevices, targetDevice)

	fmt.Println("Number of paths: ", len(allPaths))
	fmt.Println()
}

func createGraph(connections map[string][]string) graph.Graph[string, string] {
	g := graph.New(graph.StringHash, graph.Directed(), graph.PreventCycles())

	for from, tos := range connections {
		g.AddVertex(from)
		for _, to := range tos {
			g.AddEdge(from, to)
		}
	}

	return g
}

// solveSecond solves part 2: count paths from 'svr' to 'out' that pass through 'fft' before 'dac'
// Uses DP on DAG for O(V) complexity instead of exponential brute force.
// The input graph has ~574 nodes and trillions of valid paths - DP is essential!
func solveSecond(filename string) {
	fmt.Println("Solving second task with file: ", filename)

	connections := readFile(filename)

	startDevice := "svr"
	targetDevice := "out"

	g := createGraph(connections)

	//For a deterministic topological ordering, use StableTopologicalSort.
	topologicalOrder, _ := graph.TopologicalSort(g)
	fmt.Println("Topological order: ", topologicalOrder)

	// Use DP approach - works because the graph is a DAG (no cycles)
	// Time complexity: O(V * 4) where V = number of nodes
	count := countPathsDAG(connections, startDevice, targetDevice, topologicalOrder)

	fmt.Println("Number of paths: ", count)
	fmt.Println()
}

func main() {
	solveFirst("input1.txt")
	solveFirst("input2.txt")

	fmt.Println()

	solveSecond("input3.txt")
	solveSecond("input2.txt")
}
