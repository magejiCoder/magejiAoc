package sequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsCircleSeq(t *testing.T) {
	notCircle := []int{1, 2, -3, 3, -2, 0, 4}
	ok := IsCircle[int](notCircle)
	assert.False(t, ok)
	circle := []int{1, 2, 3, 2, 3}
	ok = IsCircle[int](circle)
	assert.True(t, ok)
	ok = IsCircleStable[int](circle)
	assert.True(t, ok)
	fakeCircle := []int{1, 2, 3, 2, 4}
	ok = IsCircle[int](fakeCircle)
	assert.True(t, ok)
	ok = IsCircleStable[int](fakeCircle)
	assert.False(t, ok)
}
