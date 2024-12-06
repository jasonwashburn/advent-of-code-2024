package day03

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/jasonwashburn/advent-of-code-2024/utils"
)

func main() {
	Solve()
}

func Solve() {
	input := utils.ReadInput("./day03/input.txt")
	// input := utils.ReadInput("./day03/input.txt")
	total := processInstructions(input)
	fmt.Println("The total is: ", total)

	// Part 2
	total = processDayTwo(input)
	fmt.Println("The part 2 total is: ", total)
}

func processDayTwo(input string) int {
	total := 0
	tokenPattern := `(don't\(\)|do\(\)|mul\(\d+,\d+\))`
	mulPattern := `mul\((\d+),(\d+)\)`

	tokenRe := regexp.MustCompile(tokenPattern)
	mulRe := regexp.MustCompile(mulPattern)

	matches := tokenRe.FindAllString(input, -1)

	enabled := true
	for _, match := range matches {
		if match == "don't()" {
			enabled = false
		} else if match == "do()" {
			enabled = true
		} else {
			if !enabled {
				continue
			}
			mulNums := mulRe.FindAllStringSubmatch(match, -1)

			numOne, err := strconv.Atoi(mulNums[0][1])
			if err != nil {
				log.Fatal("Unable to convert numOne: ", err)
			}
			numTwo, err := strconv.Atoi(mulNums[0][2])
			if err != nil {
				log.Fatal("Unable to convert numTwo: ", err)
			}
			total += numOne * numTwo
		}
	}

	return total
}

func removeSubstrings(input string, indices [][]int) string {
	var newText string
	i := 0
	end := len(input)

	for _, index := range indices {
		j := index[0]
		newText += input[i:j]
		i = index[1] + 1
	}

	newText += input[i:end]
	return newText
}

func buildNewStringFromIndicies(input string, indices [][]int) string {
	var newString string
	for _, index := range indices {
		newString += input[index[0]:index[1]]
	}
	return newString
}

func processInstructions(input string) int {
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	matches := re.FindAllString(input, -1)
	fmt.Println("Number of matches found: ", len(matches))

	var total int
	for _, match := range matches {
		numOneStr := strings.Split(strings.TrimLeft(match, "mul("), ",")[0]
		numTwoStr := strings.TrimRight(strings.Split(match, ",")[1], ")")

		numOne, err := strconv.Atoi(numOneStr)
		if err != nil {
			log.Panic("Unable to convert string to num: ", numOneStr)
		}

		numTwo, err := strconv.Atoi(numTwoStr)
		if err != nil {
			log.Panic("Unable to convert string to num: ", numTwoStr)
		}
		total += numOne * numTwo

	}
	return total
}
