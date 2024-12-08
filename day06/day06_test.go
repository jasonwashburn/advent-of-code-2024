package day06

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateGrid(t *testing.T) {
	input := `.#.
  ..^
  #..`
	expectedGrid := [][]string{
		{".", "#", "."},
		{".", ".", "^"},
		{"#", ".", "."},
	}
	expectedStartingPosition := []int{1, 2}

	gotGrid, gotStartingPosition := createGrid(input)
	assert.Equal(t, expectedGrid, gotGrid)
	assert.Equal(t, expectedStartingPosition, gotStartingPosition)
}

func TestGridHasPosition(t *testing.T) {
	grid := [][]string{
		{".", "#", "."},
		{".", ".", "^"},
		{"#", ".", "."},
	}
	tests := []struct {
		pos      Position
		expected bool
	}{
		{
			pos:      Position{rowIndex: 0, colIndex: 0},
			expected: true,
		},
		{
			pos:      Position{rowIndex: 2, colIndex: 2},
			expected: true,
		},
		{
			pos:      Position{rowIndex: 3, colIndex: 3},
			expected: false,
		},
		{
			pos:      Position{rowIndex: -1, colIndex: 0},
			expected: false,
		},
		{
			pos:      Position{rowIndex: 3, colIndex: 0},
			expected: false,
		},
		{
			pos:      Position{rowIndex: 0, colIndex: -1},
			expected: false,
		},
		{
			pos:      Position{rowIndex: 0, colIndex: 3},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%+v", test.pos), func(t *testing.T) {
			assert.Equal(t, test.expected, positionInsideGrid(grid, test.pos))
		})
	}
}

func TestTurnRight(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: Up, expected: Right},
		{input: Right, expected: Down},
		{input: Down, expected: Left},
		{input: Left, expected: Up},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("turnRight(%s)", test.input), func(t *testing.T) {
			result, err := turnRight(test.input)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, result)
		})
	}
}
