package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadInput(t *testing.T) {
	got := ReadInput(0)
	want := "1 2\n3 4\n"

	assert.Equal(t, want, got)
}
