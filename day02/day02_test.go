package day02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveAt(t *testing.T) {
	baseSlice := []int{0, 1, 2, 3, 4}
	assert.Equal(t, []int{1, 2, 3, 4}, removeAt(baseSlice, 0))
	assert.Equal(t, []int{0, 2, 3, 4}, removeAt(baseSlice, 1))
	assert.Equal(t, []int{0, 1, 3, 4}, removeAt(baseSlice, 2))
	assert.Equal(t, []int{0, 1, 2, 4}, removeAt(baseSlice, 3))
	assert.Equal(t, []int{0, 1, 2, 3}, removeAt(baseSlice, 4))
}

func TestProcessLine(t *testing.T) {
	tests := []struct {
		name string
		line string
		safe bool
	}{
		{name: "Test 1", line: "7 6 4 2 1", safe: true},
		{name: "Test 2", line: "1 2 7 8 9", safe: false},
		{name: "Test 3", line: "9 7 6 2 1", safe: false},
		{name: "Test 4", line: "1 3 2 4 5", safe: false},
		{name: "Test 5", line: "8 6 4 4 1", safe: false},
		{name: "Test 6", line: "1 3 6 7 9", safe: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := processLine(test.line)
			want := test.safe

			assert.Equal(t, want, got)
		})
	}
}

func TestProcessLineWithDampener(t *testing.T) {
	tests := []struct {
		name string
		line string
		safe bool
	}{
		{name: "Test 1", line: "7 6 4 2 1", safe: true},
		{name: "Test 2", line: "1 2 7 8 9", safe: false},
		{name: "Test 3", line: "9 7 6 2 1", safe: false},
		{name: "Test 4", line: "1 3 2 4 5", safe: true},
		{name: "Test 5", line: "8 6 4 4 1", safe: true},
		{name: "Test 6", line: "1 3 6 7 9", safe: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := processLineWithDampener(test.line)
			want := test.safe

			assert.Equal(t, want, got)
		})
	}
}

func TestSplitLine(t *testing.T) {
	tests := []struct {
		name string
		line string
		want []int
	}{
		{name: "Test 1", line: "7 6 4 2 1", want: []int{7, 6, 4, 2, 1}},
		{name: "Test 2", line: "1 2 7 8 9", want: []int{1, 2, 7, 8, 9}},
		{name: "Test 3", line: "9 7 6 2 1", want: []int{9, 7, 6, 2, 1}},
		{name: "Test 4", line: "1 3 2 4 5", want: []int{1, 3, 2, 4, 5}},
		{name: "Test 5", line: "8 6 4 4 1", want: []int{8, 6, 4, 4, 1}},
		{name: "Test 6", line: "1 3 6 7 9", want: []int{1, 3, 6, 7, 9}},
		{name: "Test 7", line: "1 3 6 7 9 10 11", want: []int{1, 3, 6, 7, 9, 10, 11}},
		{name: "Test 8", line: " 1 3 6 7 9 ", want: []int{1, 3, 6, 7, 9}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := splitLine(test.line)
			want := test.want

			assert.Equal(t, want, got)
		})
	}
}

func TestCheckIncrAndSafe(t *testing.T) {
	tests := []struct {
		name   string
		first  int
		second int
		incr   bool
		safe   bool
	}{
		{
			name:   "safe increase",
			first:  1,
			second: 4,
			incr:   true,
			safe:   true,
		},
		{
			name:   "unsafe increase",
			first:  1,
			second: 5,
			incr:   true,
			safe:   false,
		},
		{
			name:   "safe decrease",
			first:  4,
			second: 1,
			incr:   false,
			safe:   true,
		},
		{
			name:   "unsafe decrease",
			first:  5,
			second: 1,
			incr:   false,
			safe:   false,
		},
		{
			name:   "unsafe no change",
			first:  1,
			second: 1,
			incr:   false,
			safe:   false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotIncr, gotSafe := checkIncrAndSafe(test.first, test.second)
			wantIncr := test.incr
			wantSafe := test.safe

			assert.Equal(t, wantIncr, gotIncr)
			assert.Equal(t, wantSafe, gotSafe)
		})
	}
}
