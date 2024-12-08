package day06

import (
	"fmt"
	"log"
	"strconv"
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

	visited, _ := walkGrid(grid, startPosition)
	fmt.Println("Part 1 - Unique positions visited: ", len(visited))
	fmt.Println("-------------------------------------------------")

	// Part 2

	// Loop through each of the unique positions from before
	var loopCausingPositions []string
	loopCount := 0
	for strPos := range visited {
		loopCount += 1
		fmt.Printf("Starting loop %04d of %04d\n", loopCount, len(visited))
		pos := strings.Split(strPos, ",")
		rowIndex, err := strconv.Atoi(pos[0])
		if err != nil {
			log.Panic(err)
		}
		colIndex, err := strconv.Atoi(pos[1])
		if err != nil {
			log.Panic(err)
		}

		// ...and place an obstacle at that location then try to walk again
		grid[rowIndex][colIndex] = "O"
		_, err = walkGrid(grid, startPosition)
		// record any positions which cause a loop,
		if err != nil {
			loopCausingPositions = append(loopCausingPositions, strPos)
		}
		// then switch the grid back to original,
		grid[rowIndex][colIndex] = "."
		// and go to next position
	}
	fmt.Println("Part 2 - Number of blocking positions found: ", len(loopCausingPositions))
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

func walkGrid(grid [][]string, startPosition []int) (positions map[string]string, err error) {
	direction := Up
	currentPosition := Position{rowIndex: startPosition[0], colIndex: startPosition[1]}
	var newPosition Position
	positions = make(map[string]string)
	stringPos := fmt.Sprintf("%d,%d", newPosition.rowIndex, newPosition.colIndex)
	positions[stringPos] = Up
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
			// fmt.Println("Exiting from: ", currentPosition, " while moving ", direction)
			break
		}

		// check to see what's in the newPosition
		newPositionContent := grid[newPosition.rowIndex][newPosition.colIndex]
		if newPositionContent == "#" || newPositionContent == "O" {
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
			// check to see if we're looping (same position, same direction as before)
			if positions[stringPos] == direction {
				// we've been here before, going in same direction
				return positions, fmt.Errorf("loop detected at %s, while heading %s", stringPos, direction)
			}
			// otherwise, keep truckin...
			positions[stringPos] = direction
		} else {
			log.Panic("encountered unexpected content: ", newPositionContent)
		}

	}
	return positions, nil
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
