package day04

import (
	"fmt"
	"strings"

	"github.com/jasonwashburn/advent-of-code-2024/utils"
)

func main() {
	Solve()
}

func createGrid(input string) [][]string {
	var grid [][]string
	for _, line := range strings.Split(input, "\n") {
		var row []string
		for _, char := range strings.TrimSpace(line) {
			row = append(row, string(char))
		}
		if len(row) != 0 {
			grid = append(grid, row)
		}
	}
	return grid
}

func Solve() {
	// input := utils.ReadInput("./day04/sample.txt")
	input := utils.ReadInput("./day04/input.txt")
	grid := createGrid(input)

	var count int
	for rowIndex, line := range grid {
		for colIndex := range line {
			rightCount := searchRight(line, colIndex)
			leftCount := searchLeft(line, colIndex)
			downCount := searchDown(grid, rowIndex, colIndex)
			upCount := searchUp(grid, rowIndex, colIndex)
			upRightCount := searchUpRight(grid, rowIndex, colIndex)
			downRightCount := searchDownRight(grid, rowIndex, colIndex)
			downLeftCount := searchDownLeft(grid, rowIndex, colIndex)
			// fmt.Println(rowIndex, ",", colIndex, ": ", downLeftCount)
			upLeftCount := searchUpLeft(grid, rowIndex, colIndex)

			count += (rightCount + leftCount + downCount + upCount + upRightCount + downRightCount + downLeftCount + upLeftCount)
		}
	}

	fmt.Println("Total number found: ", count)

	ProcessPartTwo()
}

func ProcessPartTwo() {
	// input := utils.ReadInput("./day04/sample.txt")
	input := utils.ReadInput("./day04/input.txt")
	grid := createGrid(input)
	var count int
	for rowIndex, row := range grid {
		for colIndex, char := range row {
			if char == `A` {
				if colIndex-1 < 0 || colIndex+1 > len(row)-1 {
					// Not enough left or right
					continue
				}
				if rowIndex-1 < 0 || rowIndex+1 > len(grid)-1 {
					// Not enough up or down
					continue
				}
				upLeft := grid[rowIndex-1][colIndex-1]
				upRight := grid[rowIndex-1][colIndex+1]
				downRight := grid[rowIndex+1][colIndex+1]
				downLeft := grid[rowIndex+1][colIndex-1]
				center := char

				slash := strings.Join([]string{downLeft, center, upRight}, "")
				backslash := strings.Join([]string{downRight, center, upLeft}, "")

				slashMatch := slash == "MAS" || slash == "SAM"
				backslashMatch := backslash == "MAS" || backslash == "SAM"

				if slashMatch && backslashMatch {
					count += 1
				}

			}
		}
	}
	fmt.Println("Part 2, found: ", count)
}

func searchRight(row []string, colIndex int) int {
	numFound := 0
	rowLength := len(row)

	if NotEnoughRight(rowLength, colIndex) {
		return numFound
	}
	substring := strings.Join(row[colIndex:colIndex+4], "")
	if substring == "XMAS" {
		numFound += 1
	}
	return numFound
}

func searchLeft(row []string, colIndex int) int {
	numFound := 0

	if NotEnoughLeft(colIndex) {
		return numFound
	}

	substring := strings.Join(row[colIndex-3:colIndex+1], "")
	if substring == "SAMX" {
		numFound += 1
	}
	return numFound
}

func searchDown(grid [][]string, rowIndex, colIndex int) int {
	numFound := 0
	numRows := len(grid)

	if NotEnoughDown(numRows, rowIndex) {
		return numFound
	}

	first := string(grid[rowIndex][colIndex])
	second := string(grid[rowIndex+1][colIndex])
	third := string(grid[rowIndex+2][colIndex])
	fourth := string(grid[rowIndex+3][colIndex])

	substring := first + second + third + fourth

	if substring == "XMAS" {
		numFound += 1
	}

	return numFound
}

func searchUp(grid [][]string, rowIndex, colIndex int) int {
	numFound := 0

	if NotEnoughUp(rowIndex) {
		return numFound
	}

	first := string(grid[rowIndex][colIndex])
	second := string(grid[rowIndex-1][colIndex])
	third := string(grid[rowIndex-2][colIndex])
	fourth := string(grid[rowIndex-3][colIndex])

	substring := first + second + third + fourth

	if substring == "XMAS" {
		numFound += 1
	}

	return numFound
}

func searchUpRight(grid [][]string, rowIndex, colIndex int) int {
	numFound := 0
	rowLength := len(grid[rowIndex])

	if NotEnoughUp(rowIndex) || NotEnoughRight(rowLength, colIndex) {
		return numFound
	}

	first := string(grid[rowIndex][colIndex])
	second := string(grid[rowIndex-1][colIndex+1])
	third := string(grid[rowIndex-2][colIndex+2])
	fourth := string(grid[rowIndex-3][colIndex+3])

	substring := first + second + third + fourth

	if substring == "XMAS" {
		numFound += 1
	}

	return numFound
}

func searchDownRight(grid [][]string, rowIndex, colIndex int) int {
	numFound := 0
	rowLength := len(grid[rowIndex])
	numRows := len(grid)

	if NotEnoughDown(numRows, rowIndex) || NotEnoughRight(rowLength, colIndex) {
		return numFound
	}

	first := string(grid[rowIndex][colIndex])
	second := string(grid[rowIndex+1][colIndex+1])
	third := string(grid[rowIndex+2][colIndex+2])
	fourth := string(grid[rowIndex+3][colIndex+3])

	substring := first + second + third + fourth

	if substring == "XMAS" {
		numFound += 1
	}

	return numFound
}

func searchDownLeft(grid [][]string, rowIndex, colIndex int) int {
	numFound := 0
	numRows := len(grid)

	if NotEnoughDown(numRows, rowIndex) || NotEnoughLeft(colIndex) {
		return numFound
	}

	first := string(grid[rowIndex][colIndex])
	second := string(grid[rowIndex+1][colIndex-1])
	third := string(grid[rowIndex+2][colIndex-2])
	fourth := string(grid[rowIndex+3][colIndex-3])

	substring := first + second + third + fourth

	if substring == "XMAS" {
		numFound += 1
	}

	return numFound
}

func searchUpLeft(grid [][]string, rowIndex, colIndex int) int {
	numFound := 0

	if NotEnoughUp(rowIndex) || NotEnoughLeft(colIndex) {
		return numFound
	}

	first := string(grid[rowIndex][colIndex])
	second := string(grid[rowIndex-1][colIndex-1])
	third := string(grid[rowIndex-2][colIndex-2])
	fourth := string(grid[rowIndex-3][colIndex-3])

	substring := first + second + third + fourth
	if substring == "XMAS" {
		numFound += 1
	}

	return numFound
}

func NotEnoughDown(numRows, rowIndex int) bool {
	return numRows-rowIndex < 4
}

func NotEnoughUp(rowIndex int) bool {
	return rowIndex-3 < 0
}

func NotEnoughLeft(colIndex int) bool {
	return colIndex-3 < 0
}

func NotEnoughRight(rowLength, colIndex int) bool {
	return rowLength-colIndex < 4
}
