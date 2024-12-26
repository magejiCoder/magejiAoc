package seq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombination(t *testing.T) {
	seq := []int{1, 2, 3, 4}
	n := 2
	result := Combination(seq, n)
	assert.Equal(t, 6, len(result))
}
