package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadInput(t *testing.T) {
	got := ReadInput("../day00/input.txt")
	want := "1 2\n3 4\n"

	assert.Equal(t, want, got)
}

func TestAbs(t *testing.T) {
	tests := []struct {
		name     string
		num      int
		expected int
	}{
		{
			name:     "positive num",
			num:      10,
			expected: 10,
		},
		{
			name:     "negative num",
			num:      -10,
			expected: 10,
		},
		{
			name:     "zero",
			num:      0,
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Abs(test.num)
			want := test.expected
			assert.Equal(t, want, got)
		})
	}
}
