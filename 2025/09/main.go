package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Represent green tiles using polygon edges and interior detection
type TileSet struct {
	edges   [][4]int // List of line segments [x1,y1,x2,y2]
	polygon [][]int  // The polygon vertices in order
}

type Pair struct {
	i, j        int
	maxPossible float64
}

func readFile(filename string) [][]int {
	var result [][]int

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		chars := strings.Split(line, ",")
		x, _ := strconv.Atoi(chars[0])
		y, _ := strconv.Atoi(chars[1])
		result = append(result, []int{x, y})
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	return result
}
func calcRectangleSize(x1, y1, x2, y2 int) float64 {
	width := math.Abs(float64(x2 - x1))
	height := math.Abs(float64(y2 - y1))
	return (width + 1) * (height + 1)
}

func findBiggestRectangle(redTiles [][]int) ([4]int, float64) {
	var maxSize float64 = 0
	var coords [4]int
	for i := 0; i < len(redTiles); i++ {
		for j := i + 1; j < len(redTiles); j++ {
			x1, y1 := redTiles[i][0], redTiles[i][1]
			x2, y2 := redTiles[j][0], redTiles[j][1]

			size := calcRectangleSize(x1, y1, x2, y2)

			if size > maxSize {
				maxSize = size
				coords = [4]int{x1, y1, x2, y2}
			}
		}
	}

	return coords, maxSize

}

func solveFirst(filename string) {
	fmt.Println("Solving first task with file: ", filename)

	redTiles := readFile(filename)
	coords, size := findBiggestRectangle(redTiles)
	fmt.Printf("Biggest rectangle coordinates: (%d, %d) to (%d, %d) "+
		"with size %.0f\n", coords[0], coords[1], coords[2], coords[3], size)

	fmt.Println()
}

func buildGreenTileSet(redTiles [][]int) *TileSet {
	ts := &TileSet{
		edges:   make([][4]int, 0),
		polygon: redTiles,
	}

	// Build edges by connecting consecutive red tiles
	for i := 0; i < len(redTiles); i++ {
		curr := redTiles[i]
		next := redTiles[(i+1)%len(redTiles)] // wrap around

		ts.edges = append(ts.edges, [4]int{curr[0], curr[1], next[0], next[1]})
	}

	return ts
}

// Check if point is on an edge
func (ts *TileSet) isOnEdge(x, y int) bool {
	for _, edge := range ts.edges {
		x1, y1, x2, y2 := edge[0], edge[1], edge[2], edge[3]

		// Check if point is on horizontal edge
		if y1 == y2 && y == y1 {
			minX := int(math.Min(float64(x1), float64(x2)))
			maxX := int(math.Max(float64(x1), float64(x2)))
			if x >= minX && x <= maxX {
				return true
			}
		}

		// Check if point is on vertical edge
		if x1 == x2 && x == x1 {
			minY := int(math.Min(float64(y1), float64(y2)))
			maxY := int(math.Max(float64(y1), float64(y2)))
			if y >= minY && y <= maxY {
				return true
			}
		}
	}
	return false
}

// Check if point is inside the polygon
func (ts *TileSet) isInsidePolygon(x, y int) bool {
	n := len(ts.polygon)
	inside := false

	j := n - 1
	for i := 0; i < n; i++ {
		xi, yi := ts.polygon[i][0], ts.polygon[i][1]
		xj, yj := ts.polygon[j][0], ts.polygon[j][1]

		if ((yi > y) != (yj > y)) && (x < (xj-xi)*(y-yi)/(yj-yi)+xi) {
			inside = !inside
		}
		j = i
	}

	return inside
}

// Checks if a given point (x, y) is either on the edge or inside the polygon of the TileSet.
func (ts *TileSet) contains(x, y int) bool {
	return ts.isOnEdge(x, y) || ts.isInsidePolygon(x, y)
}

func isRectanglePossible(ts *TileSet, x1, y1, x2, y2 int) bool {
	// Check whether the rectangle defined by (x1, y1) and (x2, y2)
	// contains only green tiles (tiles in tileset)
	minX := int(math.Min(float64(x1), float64(x2)))
	maxX := int(math.Max(float64(x1), float64(x2)))
	minY := int(math.Min(float64(y1), float64(y2)))
	maxY := int(math.Max(float64(y1), float64(y2)))

	width := maxX - minX + 1
	height := maxY - minY + 1

	// For smaller rectangles, check all points
	if width*height <= 10000 {
		for y := minY; y <= maxY; y++ {
			for x := minX; x <= maxX; x++ {
				if !ts.contains(x, y) {
					return false
				}
			}
		}
		return true
	}

	// For larger rectangles, we need a different strategy
	// Check if all 4 corners are inside
	if !ts.contains(minX, minY) || !ts.contains(maxX, maxY) {
		return false
	}
	if !ts.contains(minX, maxY) || !ts.contains(maxX, minY) {
		return false
	}

	// Check all 4 edges comprehensively
	// Top and bottom edges
	for x := minX; x <= maxX; x++ {
		if !ts.contains(x, minY) || !ts.contains(x, maxY) {
			return false
		}
	}

	// Left and right edges
	for y := minY; y <= maxY; y++ {
		if !ts.contains(minX, y) || !ts.contains(maxX, y) {
			return false
		}
	}

	// Sample interior more carefully - check multiple rows and columns
	sampleStep := int(math.Max(float64(width)/20, float64(height)/20))
	if sampleStep < 1 {
		sampleStep = 1
	}

	for y := minY; y <= maxY; y += sampleStep {
		for x := minX; x <= maxX; x += sampleStep {
			if !ts.contains(x, y) {
				return false
			}
		}
	}

	return true
}

func findBiggestAppropriateRectangle(redTiles [][]int, ts *TileSet) ([4]int, float64) {
	var maxSize float64 = 0
	var coords [4]int

	n := len(redTiles)

	// Optimization: Sort by potential area contribution and check most promising pairs first

	fmt.Println("Generating candidate pairs...")
	candidates := make([]Pair, 0, n*n/2)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			x1, y1 := redTiles[i][0], redTiles[i][1]
			x2, y2 := redTiles[j][0], redTiles[j][1]
			size := calcRectangleSize(x1, y1, x2, y2)
			candidates = append(candidates, Pair{i, j, size})
		}
	}

	// Sort candidates by size descending
	fmt.Printf("Sorting %d candidates by size...\n", len(candidates))
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].maxPossible > candidates[j].maxPossible
	})

	fmt.Println("Checking candidates in order of decreasing size...")
	checked := 0
	for _, pair := range candidates {
		checked++
		if checked%10000 == 0 {
			fmt.Printf("Checked %d/%d candidates, current max: %.0f\n", checked, len(candidates), maxSize)
		}

		// Early termination: if no remaining candidate can beat current max, stop
		if pair.maxPossible <= maxSize {
			fmt.Printf("Early termination at %d/%d candidates\n", checked, len(candidates))
			break
		}

		x1, y1 := redTiles[pair.i][0], redTiles[pair.i][1]
		x2, y2 := redTiles[pair.j][0], redTiles[pair.j][1]

		if isRectanglePossible(ts, x1, y1, x2, y2) {
			maxSize = pair.maxPossible
			coords = [4]int{x1, y1, x2, y2}
			fmt.Printf("Found valid rectangle with size %.0f\n", maxSize)
		}
	}

	return coords, maxSize
}

func solveSecond(filename string) {
	fmt.Println("Solving second task with file: ", filename)

	redTiles := readFile(filename)
	fmt.Printf("Read %d red tiles\n", len(redTiles))

	fmt.Println("Building green tile set...")
	ts := buildGreenTileSet(redTiles)
	fmt.Printf("Green tile set built\n")

	fmt.Println("Finding biggest appropriate rectangle...")
	coords, size := findBiggestAppropriateRectangle(redTiles, ts)

	fmt.Printf("Biggest rectangle coordinates: (%d, %d) to (%d, %d) "+
		"with size %.0f\n", coords[0], coords[1], coords[2], coords[3], size)
}

func main() {

	solveFirst("input1.txt")
	solveFirst("input2.txt")

	fmt.Println()

	solveSecond("input1.txt")
	solveSecond("input2.txt")

}
