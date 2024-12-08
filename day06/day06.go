package day06

import (
	"fmt"
	"log"
	"strings"

	"github.com/jasonwashburn/advent-of-code-2024/utils"
)

const (
	Up    = "up"
	Right = "right"
	Down  = "down"
	Left  = "left"
)

func main() {
	Solve()
}

func Solve() {
	// input := utils.ReadInput("./day06/sample.txt")
	input := utils.ReadInput("./day06/input.txt")
	grid, startPosition := createGrid(input)

	visited := walkGrid(grid, startPosition)
	fmt.Println("Unique positions visited: ", len(visited))

	// Part 2

	// Loop through each of the unique positions from before
	// and place an obstacle at that location then try to walk again
	// record any positions which cause a loop,
	// then switch the grid back to original,
	// and go to next position
}

func createGrid(input string) (grid [][]string, startPosition []int) {
	for rowIndex, line := range strings.Split(input, "\n") {
		var row []string
		for colIndex, char := range strings.TrimSpace(line) {
			if string(char) == "^" {
				startPosition = []int{rowIndex, colIndex}
			}
			row = append(row, string(char))
		}
		if len(row) != 0 {
			grid = append(grid, row)
		}
	}
	// replace starting position with a "."
	grid[startPosition[0]][startPosition[1]] = "."
	fmt.Println("NumRows: ", len(grid), " NumCols: ", len(grid[0]))
	return
}

func positionInsideGrid(grid [][]string, position Position) bool {
	numRows := len(grid)
	numCols := len(grid[0])

	rowIndex := position.rowIndex
	colIndex := position.colIndex

	return rowIndex >= 0 && rowIndex < numRows && colIndex >= 0 && colIndex < numCols
}

type Position struct {
	rowIndex int
	colIndex int
}

func walkGrid(grid [][]string, startPosition []int) map[string]bool {
	direction := Up
	currentPosition := Position{rowIndex: startPosition[0], colIndex: startPosition[1]}
	var newPosition Position

	positions := make(map[string]bool)
	stringPos := fmt.Sprintf("%d,%d", newPosition.rowIndex, newPosition.colIndex)
	positions[stringPos] = true
	for {
		switch direction {
		case Up:
			newPosition = Position{currentPosition.rowIndex - 1, currentPosition.colIndex}
		case Right:
			newPosition = Position{currentPosition.rowIndex, currentPosition.colIndex + 1}
		case Down:
			newPosition = Position{currentPosition.rowIndex + 1, currentPosition.colIndex}
		case Left:
			newPosition = Position{currentPosition.rowIndex, currentPosition.colIndex - 1}
		default:
			log.Panicf("invalid direction detected at beginning of loop: %s", direction)
		}
		if !positionInsideGrid(grid, newPosition) {
			fmt.Println("Exiting from: ", currentPosition, " while moving ", direction)
			break
		}

		// check to see what's in the newPosition
		newPositionContent := grid[newPosition.rowIndex][newPosition.colIndex]
		if newPositionContent == "#" {
			// if it's a #, turn right, stay in curent position and continue loop
			var err error
			direction, err = turnRight(direction)
			if err != nil {
				log.Panic(err)
			}
			continue
		} else if newPositionContent == "." {
			// if it's a ., update current position, add new position to list of positions
			// also change content to an X so we can track where we've been
			currentPosition = newPosition
			stringPos := fmt.Sprintf("%d,%d", newPosition.rowIndex, newPosition.colIndex)
			positions[stringPos] = true
		} else {
			log.Panic("encountered unexpected content: ", newPositionContent)
		}

	}
	return positions
}

func turnRight(dir string) (string, error) {
	if dir == Up {
		return Right, nil
	} else if dir == Right {
		return Down, nil
	} else if dir == Down {
		return Left, nil
	} else if dir == Left {
		return Up, nil
	}
	return "", fmt.Errorf("%s is not a valid direction", dir)
}
