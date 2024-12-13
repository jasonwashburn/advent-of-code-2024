package day07

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/jasonwashburn/advent-of-code-2024/utils"
)

type Equation struct {
	bigNum  int
	lilNums []int
}

func Solve() {
	// input := utils.ReadInput("./day07/sample.txt")
	input := utils.ReadInput("./day07/input.txt")

	var equations []Equation
	var goodEquations []Equation
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		bigNum, lilNums := processLine(line)

		// reverse the numbers since we're going to be working back from 0
		slices.Reverse(lilNums)

		equations = append(equations, Equation{bigNum: bigNum, lilNums: lilNums})
	}
	fmt.Println(equations)

	sumOfGoodValues := 0
	for _, equation := range equations {
		if findOperations(equation.bigNum, equation.lilNums, equation.bigNum, 0) {
			goodEquations = append(goodEquations, equation)
			sumOfGoodValues += equation.bigNum
		}
	}

	fmt.Println(goodEquations)
	fmt.Println("Part 1: Sum = ", sumOfGoodValues)
}

func findOperations(bigNum int, lilNums []int, currentNum int, index int) bool {
	// base case
	if index == len(lilNums)-1 {
		// fmt.Printf("Base case reached: %+v, %+v, %+v, %+v\n", bigNum, lilNums, currentNum, index)
		return lilNums[index] == currentNum
	}
	nextNum := lilNums[index]
	index += 1
	if findOperations(bigNum, lilNums, currentNum-nextNum, index) || (nextNum != 0 && currentNum%nextNum == 0 && findOperations(bigNum, lilNums, currentNum/nextNum, index)) {
		return true
	}

	return false
}

func processLine(line string) (bigNum int, lilNums []int) {
	firstSplit := strings.Split(strings.TrimSpace(line), ":")
	bigNum, err := strconv.Atoi(firstSplit[0])
	if err != nil {
		log.Panicf("Unable to convert bigNum %s to int", firstSplit[0])
	}
	for _, char := range strings.Split(strings.TrimSpace(firstSplit[1]), " ") {
		num, err := strconv.Atoi(char)
		if err != nil {
			log.Panicf("Unable to convert char %s to int", char)
		}
		lilNums = append(lilNums, num)
	}

	return bigNum, lilNums
}

func concatNums(a, b int) int {
	numA := strconv.Itoa(a)
	numB := strconv.Itoa(b)
	output, err := strconv.Atoi(numA + numB)
	if err != nil {
		log.Panic("Unable to concat nums")
	}

	return output
}
