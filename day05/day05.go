package day05

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

type UpdateRecord struct {
	pages []int
}

type pageSet map[int]struct{}

func (ps pageSet) contains(n int) bool {
	if _, ok := ps[n]; ok {
		return true
	}
	return false
}

func Solve() {
	// input := utils.ReadInput("./day05/sample.txt")
	input := utils.ReadInput("./day05/input.txt")
	mustComeAfters := make(map[int]pageSet)
	var updateRecords []UpdateRecord
	var badUpdateRecords []UpdateRecord

	sectionOne := true
	for _, line := range strings.Split(input, "\n") {
		if sectionOne {
			if strings.TrimSpace(line) == "" {
				// Go to section 2
				sectionOne = false
				fmt.Println("Start Section Two")
				continue
			}
			// Process section 1H
			pair := processPairLine(line)
			before := pair[0]
			after := pair[1]

			if _, ok := mustComeAfters[before]; ok {
				mustComeAfters[before][after] = struct{}{}
			} else {
				{
				}
				newEntry := make(pageSet)
				newEntry[after] = struct{}{}
				mustComeAfters[before] = newEntry
			}

		} else {
			// Process Section 2
			if line == "" {
				continue
			}

			numSlice := []int{}
			for _, char := range strings.Split(strings.TrimSpace(line), ",") {
				num, err := strconv.Atoi(char)
				if err != nil {
					log.Panic("Unable to convert character to num: ", char)
				}
				numSlice = append(numSlice, num)
			}

			updateRecords = append(updateRecords, UpdateRecord{pages: numSlice})
		}
	}

	// Go through records

	returnTotal := 0
	for _, record := range updateRecords {
		goodRecord := true
		pages := record.pages
		for currentPageIndex, currentPage := range pages {
			for _, laterPage := range pages[currentPageIndex+1:] {
				afterRules := mustComeAfters[laterPage]
				// if currentPage is in afterRules for later pages, it violates order
				// because it is currently before the later page under test
				if afterRules.contains(currentPage) {
					// Out of order
					goodRecord = false
					break
				}

			}
			if !goodRecord {
				// found a bad page, no reason to continue through the rest for this
				// record
				//
				// add record to list of bad records for processing in part 2
				break
			}
		}
		if goodRecord {
			// This record is in order
			middleIndex := len(pages) / 2
			middlePageNum := pages[middleIndex]
			returnTotal += middlePageNum
		} else {
			badUpdateRecords = append(badUpdateRecords, record)
		}
	}
	fmt.Println("Part 1 Total is : ", returnTotal)

	fmt.Println("Part 2...........")

	badMiddleTotal := 0
	for _, badRecord := range badUpdateRecords {
		reorderedPages := reorderPages(badRecord.pages, mustComeAfters)
		middleNum := reorderedPages[len(reorderedPages)/2]
		badMiddleTotal += middleNum
	}
	fmt.Println("Part 2 Total: ", badMiddleTotal)
}

func reorderPages(pages []int, mustComeAfters map[int]pageSet) []int {
	outOfOrder := true
	for outOfOrder {
		goodRecord := true
		for currentPageIndex, currentPage := range pages {
			for laterPageIndex, laterPage := range pages[currentPageIndex+1:] {
				afterRules := mustComeAfters[laterPage]
				// if currentPage is in afterRules for later pages, it violates order
				// because it is currently before the later page under test
				if afterRules.contains(currentPage) {
					// Out of order so we must make a swap
					pages[currentPageIndex] = laterPage
					pages[currentPageIndex+1+laterPageIndex] = currentPage
					goodRecord = false
					break
				}
			}
			if !goodRecord {
				break
			}
		}
		if goodRecord {
			outOfOrder = false
		}
	}

	return pages
}

func processPairLine(pair string) []int {
	pair = strings.TrimSpace(pair)
	split := strings.Split(pair, "|")
	numOne, err := strconv.Atoi(split[0])
	if err != nil {
		log.Panic("Unable to convert string to int: ", split[0])
	}
	numTwo, err := strconv.Atoi(split[1])
	if err != nil {
		log.Panic("Unable to convert string to int: ", split[0])
	}

	return []int{numOne, numTwo}
}
