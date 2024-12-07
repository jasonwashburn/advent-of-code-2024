package day04

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const input string = `....XXMAS.
.SAMXMS...
...S..A...
..A.A.MS.X
XMASAMX.MM
X.....XA.A
S.S.S.S.SS
.A.A.A.A.A
..M.M.M.MM
.X.X.XMASX
..........
XMAS......
MM........
A.A.......
S..S......`

func TestCreateGrid(t *testing.T) {
	input := `123
  456
  789`
	expected := [][]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}
	assert.Equal(t, expected, createGrid(input))
}

var grid = createGrid(input)

func TestSearchRight(t *testing.T) {
	var count int
	for _, row := range grid {
		for colIndex := range row {
			count += searchRight(row, colIndex)
		}
	}
	expected := 4
	assert.Equal(t, expected, count)
}

func TestSearchLeft(t *testing.T) {
	var count int
	for _, row := range grid {
		for colIndex := range row {
			count += searchLeft(row, colIndex)
		}
	}
	expected := 2
	assert.Equal(t, expected, count)
}

func TestSearchUp(t *testing.T) {
	var count int
	for rowIndex, row := range grid {
		for colIndex := range row {
			count += searchUp(grid, rowIndex, colIndex)
		}
	}
	expected := 2
	assert.Equal(t, expected, count)
}

func TestSearchDown(t *testing.T) {
	var count int
	for rowIndex, row := range grid {
		for colIndex := range row {
			count += searchDown(grid, rowIndex, colIndex)
		}
	}
	expected := 2
	assert.Equal(t, expected, count)
}

func TestSearchUpRight(t *testing.T) {
	var count int
	for rowIndex, row := range grid {
		for colIndex := range row {
			count += searchUpRight(grid, rowIndex, colIndex)
		}
	}
	expected := 4
	assert.Equal(t, expected, count)
}

func TestSearchDownRight(t *testing.T) {
	var count int
	for rowIndex, row := range grid {
		for colIndex := range row {
			count += searchDownRight(grid, rowIndex, colIndex)
		}
	}
	expected := 2
	assert.Equal(t, expected, count)
}

func TestSearchDownLeft(t *testing.T) {
	var count int
	for rowIndex, row := range grid {
		for colIndex := range row {
			count += searchDownLeft(grid, rowIndex, colIndex)
		}
	}
	expected := 1
	assert.Equal(t, expected, count)
}

func TestSearchUpLeft(t *testing.T) {
	var count int
	for rowIndex, row := range grid {
		for colIndex := range row {
			count += searchUpLeft(grid, rowIndex, colIndex)
		}
	}
	expected := 4
	assert.Equal(t, expected, count)
}

func TestNotEnoughLeft(t *testing.T) {
	// [0,1,2,3,4]
	tests := []struct {
		index    int
		expected bool
	}{
		{
			0, true,
		},
		{
			1, true,
		},
		{
			2, true,
		},
		{
			3, false,
		},
		{
			4, false,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Index: %d", test.index), func(t *testing.T) {
			assert.Equal(t, test.expected, NotEnoughLeft(test.index))
		})
	}
}

func TestNotEnoughDown(t *testing.T) {
	// 0 [0,1,2,3,4]
	// 1 [0,1,2,3,4]
	// 2 [0,1,2,3,4]
	// 3 [0,1,2,3,4]
	// 4 [0,1,2,3,4]
	numRows := 5
	tests := []struct {
		rowIndex int
		expected bool
	}{
		{
			0, false,
		},
		{
			1, false,
		},
		{
			2, true,
		},
		{
			3, true,
		},
		{
			4, true,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Index: %d", test.rowIndex), func(t *testing.T) {
			assert.Equal(t, test.expected, NotEnoughDown(numRows, test.rowIndex))
		})
	}
}
