package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/dominikbraun/graph"
	"github.com/golang/geo/r3"
)

type myVector struct {
	vector r3.Vector
	name   int
}

func readFile(filename string) []myVector {
	var result []myVector
	counter := 0

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		coordinates := strings.Split(line, ",")
		x, err := strconv.ParseFloat(coordinates[0], 32)
		if err != nil {
			log.Fatalf("failed to convert x coordinate: %s", err)
		}

		y, err := strconv.ParseFloat(coordinates[1], 32)
		if err != nil {
			log.Fatalf("failed to convert y coordinate: %s", err)
		}

		z, err := strconv.ParseFloat(coordinates[2], 32)
		if err != nil {
			log.Fatalf("failed to convert z coordinate: %s", err)
		}

		result = append(result, myVector{name: counter, vector: r3.Vector{x, y, z}})
		counter++
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	return result
}

func calculateDistance(a, b myVector) float64 {

	//return vec3.Distance(&a.vector, &b.vector)
	//return vec3.SquareDistance(&a.vector, &b.vector)

	//Calculate Euclidean distance between two 3D points
	x1, y1, z1 := a.vector.X, a.vector.Y, a.vector.Z
	x2, y2, z2 := b.vector.X, b.vector.Y, b.vector.Z

	//Calculate the sum of squared differences
	sum_sq_diff := (x2-x1)*(x2-x1) + (y2-y1)*(y2-y1) + (z2-z1)*(z2-z1)

	//Take the square root
	distance := sum_sq_diff

	return distance

}

type vectorDistance struct {
	from     myVector
	to       myVector
	distance float64
}

func calculateDistances(points []myVector) []vectorDistance {
	var distances []vectorDistance

	for i := 0; i < len(points); i++ {

		for j := i + 1; j < len(points); j++ {

			distance := calculateDistance(points[i], points[j])
			distanceObj := vectorDistance{from: points[i], to: points[j], distance: distance}
			distances = append(distances, distanceObj)

		}
	}

	return distances

}

func sortDistances(distances []vectorDistance) {
	sort.Slice(distances, func(i, j int) bool {
		return distances[i].distance < distances[j].distance
	})
}

func createGraph(nodes []myVector) graph.Graph[int, int] {
	g := graph.New(graph.IntHash)

	for _, node := range nodes {
		err := g.AddVertex(node.name)
		if err != nil {
			return nil
		}
	}

	return g
}

func connectGraph(g graph.Graph[int, int], connections []vectorDistance) graph.Graph[int, int] {
	counter := 0
	for _, connection := range connections {
		first := connection.from
		second := connection.to

		err := g.AddEdge(first.name, second.name)
		if err != nil {
			return nil
		}
		counter++

		if counter >= 1000 {
			break
		}
	}

	return g
}

func findConnectedComponents(g graph.Graph[int, int], nodes []myVector) [][]int {
	var components [][]int
	visited := make(map[int]bool)

	for _, node := range nodes {
		var component []int

		if visited[node.name] {
			continue
		}

		graph.DFS(g, node.name, func(value int) bool {
			if !visited[value] {
				visited[value] = true
				component = append(component, value)
			}
			return false
		})

		if len(component) > 0 {
			components = append(components, component)
		}
	}

	return components

}

func sortComponentsBySize(components [][]int) {
	sort.Slice(components, func(i, j int) bool {
		return len(components[i]) > len(components[j])
	})
}

func calculateCircuitSize(component [][]int) int {
	size := 1
	componentsLen := len(component)
	if componentsLen > 3 {
		componentsLen = 3
	}

	for i := 0; i < componentsLen; i++ {
		size *= len(component[i])
	}

	return size
}

func solveFirst(filename string) {
	fmt.Println("Solving first task with file: ", filename)

	nodes := readFile(filename)

	connections := calculateDistances(nodes)

	sortDistances(connections)

	myGraph := createGraph(nodes)

	connectedGraph := connectGraph(myGraph, connections)

	components := findConnectedComponents(connectedGraph, nodes)

	sortComponentsBySize(components)

	circuitSize := calculateCircuitSize(components)

	fmt.Println("Circuit size: ", circuitSize)

	//file, _ := os.Create("./mygraph.gv")
	//_ = draw.DOT(connectedGraph, file)

	fmt.Println()
}

func areAllNodesConnected(g graph.Graph[int, int], nodesCount int) bool {
	connectedNodesCount := 0

	graph.BFS(g, 0, func(value int) bool {
		connectedNodesCount++
		return false

	})

	if nodesCount == connectedNodesCount {
		return true
	}
	return false
}

func connectFullGraph(g graph.Graph[int, int], connections []vectorDistance, nodesCount int) int {
	for _, connection := range connections {
		first := connection.from
		second := connection.to

		err := g.AddEdge(first.name, second.name)
		if err != nil {
			log.Fatal("Failed to add edge: ", err)
			return 0
		}

		if areAllNodesConnected(g, nodesCount) {
			return int(first.vector.X * second.vector.X)
		}

	}

	return 0
}

func solveSecond(filename string) {
	fmt.Println("Solving second task with file: ", filename)

	nodes := readFile(filename)

	connections := calculateDistances(nodes)

	sortDistances(connections)

	myGraph := createGraph(nodes)

	result := connectFullGraph(myGraph, connections, len(nodes))

	fmt.Println("Result: ", result)
	fmt.Println()
}

func main() {

	solveFirst("input1.txt")
	solveFirst("input2.txt")

	fmt.Println()

	solveSecond("input1.txt")
	solveSecond("input2.txt")

}
