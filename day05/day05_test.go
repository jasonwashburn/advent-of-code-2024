package day05

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessPairLine(t *testing.T) {
	got := processPairLine("47|53\n")
	want := []int{47, 53}
	assert.Equal(t, want, got)
}
