package utils

import (
	"log"
	"os"
)

func ReadInput(filename string) string {
	contents, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Unable to read from file: %s", filename)
	}
	output := string(contents)
	return output
}

func Abs(num int) int {
	if num < 0 {
		return -num
	} else {
		return num
	}
}
