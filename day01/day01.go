package day01

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/jasonwashburn/advent-of-code-2024/utils"
)

func main() {
	Solve()
}

// read input
// split into lists
// add up lists
// get difference

func Solve() {
	input := utils.ReadInput("./day01/input.txt")

	var listOne []int
	var listTwo []int

	lines := strings.Split(input, "\n")
	for i, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		lineSlice := strings.Split(strings.TrimSpace(line), "   ")
		if len(lineSlice) < 2 {
			log.Fatalf("length of lineSlice is less than 2 on input line: %d", i)
		}
		numOne, err := strconv.Atoi(lineSlice[0])
		if err != nil {
			log.Fatalf("Unable to convert numOne to int on line: %d", i)
		}
		numTwo, err := strconv.Atoi(lineSlice[1])
		if err != nil {
			log.Fatalf("Unable to convert numTwo to int on line: %d", i)
		}
		listOne = append(listOne, numOne)
		listTwo = append(listTwo, numTwo)
	}

	slices.Sort(listOne)
	slices.Sort(listTwo)

	var difference int
	for i := 0; i < len(listOne); i++ {
		difference += utils.Abs(listOne[i] - listTwo[i])
	}

	fmt.Printf("The difference between list one and list two is %d\n", difference)

	// Day two
	listTwoMap := make(map[int]int)
	for _, n := range listTwo {
		listTwoMap[n] += 1
	}

	similarityScore := 0

	for _, n := range listOne {
		total := n * listTwoMap[n]
		similarityScore += total
	}

	fmt.Printf("The total similarityScore is %d\n", similarityScore)
}

func sumSlice(numbers []int) int {
	var total int
	for _, n := range numbers {
		total += n
	}
	return total
}
