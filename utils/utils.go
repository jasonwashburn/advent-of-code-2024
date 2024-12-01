package utils

import (
	"fmt"
	"os"
)

func ReadInput(day int64) string {
	filename := fmt.Sprintf("../day%02d/input.txt", day)
	contents, err := os.ReadFile(filename)
	fmt.Printf("Contents: \n%b", contents)
	if err != nil {
		fmt.Errorf("Unable to read from file: %s", filename)
	}
	output := string(contents)
	return output
}
