package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

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
func getGridSize(redTiles [][]int) (int, int) {
	maxX, maxY := 0, 0
	for _, tile := range redTiles {
		x := tile[0]
		y := tile[1]
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
	}
	return maxX + 1, maxY + 1
}

func createGrid(redTiles [][]int, width, height int) [][]int {
	grid := make([][]int, height)
	for i := range grid {
		grid[i] = make([]int, width)
	}

	for _, tile := range redTiles {
		x := tile[0]
		y := tile[1]
		grid[y][x] = 1 // Mark red tile
	}
	return grid
}

func printGrid(grid [][]int) {
	for _, row := range grid {
		for _, cell := range row {
			if cell == 0 {
				fmt.Print(". ")
			} else if cell == 1 {
				fmt.Print("# ")
			} else if cell == 2 {
				fmt.Print("X ")
			}
		}
		fmt.Println()
	}
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

	//fmt.Println("Red tiles: ", redTiles)

	//width, height := getGridSize(redTiles)
	//fmt.Printf("Grid size: width=%d, heigh t=%d\n", width, height)

	//grid := createGrid(redTiles, width, height)

	//printGrid(grid)

	coords, size := findBiggestRectangle(redTiles)
	fmt.Printf("Biggest rectangle coordinates: (%d, %d) to (%d, %d) "+
		"with size %.0f\n", coords[0], coords[1], coords[2], coords[3], size)

	fmt.Println()
}

func shouldColorTileGreen(x, y int, greenTiles [][]int) bool {
	for _, tile := range greenTiles {
		// Check if the tile is in the same row or column
		if tile[0] == x || tile[1] == y {
			return true
		}
	}

	return false

}

func markGreenTiles(grid [][]int, redTiles [][]int) [][]int {
	var greenTiles [][]int

	// mark outer tiles
	for _, redTile := range redTiles {
		// Check if there is another red tile in the same row or column
		x, y := redTile[0], redTile[1]

		greenTiles = append(greenTiles, []int{x, y})

		for _, otherTile := range redTiles {

			if otherTile[1] == y && otherTile[0] != x {
				//	There is another red tile in the same row
				//	Color the tiles between them green
				start := int(math.Min(float64(x), float64(otherTile[0]))) + 1
				end := int(math.Max(float64(x), float64(otherTile[0])))
				for i := start; i < end; i++ {
					grid[y][i] = 2 // Mark green tile
					greenTiles = append(greenTiles, []int{x, y})
				}
			}
			if otherTile[0] == x && otherTile[1] != y {
				//	There is another red tile in the same column
				//	Color the tiles between them green
				start := int(math.Min(float64(y), float64(otherTile[1]))) + 1
				end := int(math.Max(float64(y), float64(otherTile[1])))
				for i := start; i < end; i++ {
					grid[i][x] = 2 // Mark green tile
					greenTiles = append(greenTiles, []int{x, y})
				}
			}
		}

	}

	height := len(grid)
	width := len(grid[0])

	// mark inner tiles

	for y := 1; y < height-1; y++ {
		for x := 1; x < width-1; x++ {
			if grid[y][x] == 0 && shouldColorTileGreen(x, y, greenTiles) {
				grid[y][x] = 2 // Mark green tile
			}
		}
	}

	return grid

}

func isRectanglePossible(grid [][]int, x1, y1, x2, y2 int) bool {
	// Check whether the rectangle defined by (x1, y1) and (x2, y2)
	// contains only red(1) and green(2) tiles
	minX := int(math.Min(float64(x1), float64(x2)))
	maxX := int(math.Max(float64(x1), float64(x2)))
	minY := int(math.Min(float64(y1), float64(y2)))
	maxY := int(math.Max(float64(y1), float64(y2)))

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if grid[y][x] == 0 {
				return false
			}
		}
	}

	return true
}

func findBiggestAppropriateRectangle(appropriateTiles [][]int, grid [][]int) ([4]int, float64) {
	var maxSize float64 = 0
	var coords [4]int

	for i := 0; i < len(appropriateTiles); i++ {
		for j := i + 1; j < len(appropriateTiles); j++ {
			x1, y1 := appropriateTiles[i][0], appropriateTiles[i][1]
			x2, y2 := appropriateTiles[j][0], appropriateTiles[j][1]

			if !isRectanglePossible(grid, x1, y1, x2, y2) {
				continue
			}

			size := calcRectangleSize(x1, y1, x2, y2)

			if size > maxSize {
				maxSize = size
				coords = [4]int{x1, y1, x2, y2}
			}
		}
	}

	return coords, maxSize

}

func solveSecond(filename string) {
	fmt.Println("Solving second task with file: ", filename)

	redTiles := readFile(filename)

	//fmt.Println("Red tiles: ", redTiles)

	width, height := getGridSize(redTiles)
	fmt.Printf("Grid size: width=%d, heigh t=%d\n", width, height)

	grid := createGrid(redTiles, width, height)
	fmt.Println("Grid created")

	markedGrid := markGreenTiles(grid, redTiles)
	fmt.Println("Marking green tiles done.")
	//printGrid(markedInnerGrid)

	coords, size := findBiggestAppropriateRectangle(redTiles, markedGrid)

	fmt.Printf("Biggest rectangle coordinates: (%d, %d) to (%d, %d) "+
		"with size %.0f\n", coords[0], coords[1], coords[2], coords[3], size)
}

func main() {

	//solveFirst("input1.txt")
	//solveFirst("input2.txt")

	fmt.Println()

	//solveSecond("input1.txt")
	solveSecond("input2.txt")

}
