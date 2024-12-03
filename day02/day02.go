package day02

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/jasonwashburn/advent-of-code-2024/utils"
)

func main() {
	Solve()
}

func Solve() {
	// input := utils.ReadInput("./day02/sample.txt")
	input := utils.ReadInput("./day02/input.txt")
	var safeCount int
	for _, line := range strings.Split(input, "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}

		safe := processLine(line)
		if safe {
			safeCount++
		}
	}
	fmt.Printf("There are %d safe lines\n", safeCount)

	// Part 2
	safeCount = 0
	for _, line := range strings.Split(input, "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}

		safe := processLine(line)
		if safe {
			safeCount++
		} else {
			if processLineWithDampener(line) {
				safeCount++
			}
		}

	}
	fmt.Printf("There are %d safe lines with dampener\n", safeCount)
}

func checkNumsAreSafe(nums []int) bool {
	var startingDirection bool
	var currDirection bool
	var safe bool

	startingDirection, safe = checkIncrAndSafe(nums[0], nums[1])
	if safe {
		for i := 1; i < len(nums)-1; i++ {
			currDirection, safe = checkIncrAndSafe(nums[i], nums[i+1])
			if currDirection != startingDirection || !safe {
				safe = false
				break
			}
		}
	}
	return safe
}

func removeAt(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}

	newSlice := make([]int, 0, len(slice)-1)
	newSlice = append(newSlice, slice[:index]...)
	newSlice = append(newSlice, slice[index+1:]...)

	return newSlice
}

func processLine(line string) bool {
	nums := splitLine(line)
	return checkNumsAreSafe(nums)
}

func processLineWithDampener(line string) bool {
	var safe bool = false
	nums := splitLine(line)
	fmt.Println("Original Line: ", nums)

	for dampenedIndex := range nums {
		testCase := removeAt(nums, dampenedIndex)
		fmt.Println("Test Case: ", dampenedIndex, testCase)
		safe = checkNumsAreSafe(testCase)
		if safe {
			fmt.Println("Safe")
			break
		}
		fmt.Println("Not Safe")
	}
	return safe
}

func splitLine(line string) []int {
	var output []int
	line = strings.TrimSpace(line)

	for _, num := range strings.Split(line, " ") {
		n, err := strconv.Atoi(num)
		if err != nil {
			log.Panicf("unable to convert %s to int", num)
		}
		output = append(output, n)
	}
	return output
}

func checkIncrAndSafe(first, second int) (incr bool, safe bool) {
	minDiff := 1
	maxDiff := 3
	incr = first < second

	diff := utils.Abs(first - second)
	safe = diff >= minDiff && diff <= maxDiff
	return
}
