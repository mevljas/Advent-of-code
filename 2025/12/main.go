package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Region struct {
	Size     string
	Presents []int
}

func readFile(filename string) (map[int][][]string, []Region) {
	shapes := map[int][][]string{}
	var regions []Region
	currentShapeIndx := -1
	var newShape []string

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "x") {
			//	region
			split := strings.Split(line, " ")
			regionName := split[0]
			regionName = regionName[:len(regionName)-1]

			var presents []int

			for i := 1; i < len(split); i++ {
				presentsCount, err := strconv.Atoi(split[i])
				if err != nil {
					log.Fatal(err)
				}

				presents = append(presents, presentsCount)
			}

			regions = append(regions, Region{Size: regionName, Presents: presents})

		} else if line != "" && string(line[1]) == ":" {
			//	shape
			number, err := strconv.Atoi(string(line[0]))
			if err != nil {
				log.Fatal(err)
			}
			currentShapeIndx = number
		} else if len(line) > 0 {
			newShape = append(newShape, line)

		} else {
			if currentShapeIndx != -1 {
				var shapeLine [][]string
				for _, row := range newShape {
					shapeRow := []string{}
					for _, ch := range row {
						shapeRow = append(shapeRow, string(ch))
					}
					shapeLine = append(shapeLine, shapeRow)
				}
				shapes[currentShapeIndx] = shapeLine
			}
			currentShapeIndx = -1
			newShape = []string{}
		}

	}

	// Handle the last shape if file doesn't end with empty line
	if currentShapeIndx != -1 && len(newShape) > 0 {
		var shapeLine [][]string
		for _, row := range newShape {
			shapeRow := []string{}
			for _, ch := range row {
				shapeRow = append(shapeRow, string(ch))
			}
			shapeLine = append(shapeLine, shapeRow)
		}
		shapes[currentShapeIndx] = shapeLine
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	return shapes, regions
}

func createRegionMatrix(size string) [][]string {
	split := strings.Split(size, "x")
	width, err := strconv.Atoi(split[0])
	if err != nil {
		log.Fatal(err)
	}

	height, err := strconv.Atoi(split[1])
	if err != nil {
		log.Fatal(err)
	}

	// Create height rows, each with width columns
	region := make([][]string, height)
	for i := range region {
		region[i] = make([]string, width)
		for j := range region[i] {
			region[i][j] = "."
		}
	}

	return region
}

func createRegionMatrices(regions []Region) map[string][][]string {
	regionMatrices := map[string][][]string{}

	for _, region := range regions {
		if _, exists := regionMatrices[region.Size]; !exists {
			regionMatrices[region.Size] = createRegionMatrix(region.Size)
		}
	}

	return regionMatrices
}

func rotateShapeCopy(shape [][]string, degrees int) [][]string {
	if degrees%90 != 0 {
		log.Fatal("Can only rotate in 90 degree increments")
	}
	shapeRows := len(shape)
	shapeCols := len(shape[0])
	var rotatedShape [][]string

	switch degrees % 360 {
	case 0:
		rotatedShape = make([][]string, shapeRows)
		for i := 0; i < shapeRows; i++ {
			rotatedShape[i] = make([]string, shapeCols)
			copy(rotatedShape[i], shape[i])
		}
	case 90:
		rotatedShape = make([][]string, shapeCols)
		for i := 0; i < shapeCols; i++ {
			rotatedShape[i] = make([]string, shapeRows)
			for j := 0; j < shapeRows; j++ {
				rotatedShape[i][j] = shape[shapeRows-1-j][i]
			}
		}
	case 180:
		rotatedShape = make([][]string, shapeRows)
		for i := 0; i < shapeRows; i++ {
			rotatedShape[i] = make([]string, shapeCols)
			for j := 0; j < shapeCols; j++ {
				rotatedShape[i][j] = shape[shapeRows-1-i][shapeCols-1-j]
			}
		}
	case 270:
		rotatedShape = make([][]string, shapeCols)
		for i := 0; i < shapeCols; i++ {
			rotatedShape[i] = make([]string, shapeRows)
			for j := 0; j < shapeRows; j++ {
				rotatedShape[i][j] = shape[j][shapeCols-1-i]
			}
		}
	}

	return rotatedShape
}

func canFitPresentsIntoRegionRec(shapes [][][]string, shapeIndx int, region [][]string, emptyCount int) bool {
	if shapeIndx >= len(shapes) {
		return true
	}

	// Early termination: count remaining cells needed
	cellsNeeded := 0
	for k := shapeIndx; k < len(shapes); k++ {
		shape := shapes[k]
		for _, row := range shape {
			for _, cell := range row {
				if cell == "#" {
					cellsNeeded++
				}
			}
		}
	}

	if cellsNeeded > emptyCount {
		return false
	}

	//	Try all rotations and positions
	shape := shapes[shapeIndx]
	regionRows := len(region)
	regionCols := len(region[0])

	// Find first empty cell to start from
	startRow, startCol := 0, 0
	found := false
	for i := 0; i < regionRows && !found; i++ {
		for j := 0; j < regionCols && !found; j++ {
			if region[i][j] == "." {
				startRow, startCol = i, j
				found = true
			}
		}
	}

	for rotation := 0; rotation < 360; rotation += 90 {
		rotatedShape := rotateShapeCopy(shape, rotation)

		shapeRows := len(rotatedShape)
		shapeCols := len(rotatedShape[0])

		// Start from the first empty cell position
		for i := startRow; i <= regionRows-shapeRows; i++ {
			colStart := 0
			if i == startRow {
				colStart = startCol
			}

			for j := colStart; j <= regionCols-shapeCols; j++ {
				canPlace := true

				for r := 0; r < shapeRows; r++ {
					for c := 0; c < shapeCols; c++ {
						if rotatedShape[r][c] == "#" && region[i+r][j+c] == "#" {
							canPlace = false
							break
						}
					}
					if !canPlace {
						break
					}
				}

				if canPlace {
					//	Place shape and count cells used
					cellsUsed := 0
					for r := 0; r < shapeRows; r++ {
						for c := 0; c < shapeCols; c++ {
							if rotatedShape[r][c] == "#" {
								region[i+r][j+c] = "#"
								cellsUsed++
							}
						}
					}

					if canFitPresentsIntoRegionRec(shapes, shapeIndx+1, region, emptyCount-cellsUsed) {
						return true
					}

					//	Remove shape
					for r := 0; r < shapeRows; r++ {
						for c := 0; c < shapeCols; c++ {
							if rotatedShape[r][c] == "#" {
								region[i+r][j+c] = "."
							}
						}
					}
				}
			}
		}
	}

	return false

}

func canFitAllPresentsIntoRegion(allShapes map[int][][]string, regionSize string, presents []int, regionMatrices map[string][][]string) bool {

	regionMatrixCopy := make([][]string, len(regionMatrices[regionSize]))
	for i := range regionMatrices[regionSize] {
		regionMatrixCopy[i] = make([]string, len(regionMatrices[regionSize][i]))
		copy(regionMatrixCopy[i], regionMatrices[regionSize][i])
	}

	type ShapeWithSize struct {
		shape [][]string
		size  int
	}

	var shapesWithSize []ShapeWithSize
	totalCellsNeeded := 0

	for i := 0; i < len(presents); i++ {
		count := presents[i]
		if count > 0 {
			shapeSize := 0
			for _, row := range allShapes[i] {
				for _, cell := range row {
					if cell == "#" {
						shapeSize++
					}
				}
			}

			for j := 0; j < count; j++ {
				shapesWithSize = append(shapesWithSize, ShapeWithSize{shape: allShapes[i], size: shapeSize})
				totalCellsNeeded += shapeSize
			}
		}
	}

	// Sort shapes by size (largest first) for better pruning
	sort.Slice(shapesWithSize, func(i, j int) bool {
		return shapesWithSize[i].size > shapesWithSize[j].size
	})

	var shapesToFit [][][]string
	for _, s := range shapesWithSize {
		shapesToFit = append(shapesToFit, s.shape)
	}

	// Check if total cells needed exceeds region size -> impossible
	regionCells := len(regionMatrixCopy) * len(regionMatrixCopy[0])
	if totalCellsNeeded > regionCells {
		return false
	}

	if !canFitPresentsIntoRegionRec(shapesToFit, 0, regionMatrixCopy, regionCells) {
		return false
	}

	return true
}

func countDoableRegions(shapes map[int][][]string, regions []Region, regionMatrices map[string][][]string) int {
	count := 0

	for i, region := range regions {
		canFit := canFitAllPresentsIntoRegion(shapes, region.Size, region.Presents, regionMatrices)
		fmt.Printf("Region %d: %s with presents %v: %v\n", i+1, region.Size, region.Presents, canFit)
		if canFit {
			count++
		}
	}

	return count
}

func solveFirst(filename string) {
	fmt.Println("Solving first task with file: ", filename)

	shapes, regions := readFile(filename)

	fmt.Println("Shapes: ", shapes)
	fmt.Println("Regions: ", regions)

	count := countDoableRegions(shapes, regions, createRegionMatrices(regions))
	fmt.Println("Number of regions that can fit presents: ", count)

	fmt.Println()
}

func main() {
	solveFirst("input1.txt")
	solveFirst("input2.txt")

	fmt.Println()

}
