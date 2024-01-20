package grid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLintPoints(t *testing.T) {
	l := NewLine(Vec{0, 0}, Vec{10, 10})
	ps := l.Points()
	assert.Equal(t, 11, len(ps))
	assert.Equal(t, Point{0, 0}, ps[0])
	assert.Equal(t, Point{1, 1}, ps[1])
	assert.Equal(t, Point{10, 10}, ps[10])
	l2 := NewLine(Vec{10, 10}, Vec{0, 0})
	ps = l2.Points()
	assert.Equal(t, 11, len(ps))
	assert.Equal(t, Point{10, 10}, ps[0])
	l3 := NewLine(Vec{0, 0}, Vec{0, 10})
	ps = l3.Points()
	assert.Equal(t, 11, len(ps))
	assert.Equal(t, Point{0, 0}, ps[0])
	assert.Equal(t, Point{0, 10}, ps[10])
	l4 := NewLine(Vec{0, 10}, Vec{0, 0})
	ps = l4.Points()
	assert.Equal(t, 11, len(ps))
	assert.Equal(t, Point{0, 10}, ps[0])
}

func TestLineCross(t *testing.T) {
	l1 := NewLine(Vec{0, 0}, Vec{10, 10})
	l2 := NewLine(Vec{0, 10}, Vec{10, 0})
	ok := l1.Cross(*l2)
	assert.True(t, ok)
	l1 = NewLine(Vec{0, 0}, Vec{10, 0})
	l2 = NewLine(Vec{1, 0}, Vec{1, 10})
	ok = l1.Cross(*l2)
	assert.True(t, ok)
	l1 = NewLine(Vec{0, 0}, Vec{10, 0})
	l2 = NewLine(Vec{1, 1}, Vec{1, 10})
	ok = l1.Cross(*l2)
	assert.False(t, ok)

	l1 = NewLine(Vec{0, 0}, Vec{0, 10})
	l2 = NewLine(Vec{0, 0}, Vec{0, 9})
	ok = l1.Cross(*l2)
	assert.True(t, ok)

	// (0,0)
	// (0,1)
	//
	l1 = NewLine(Vec{0, 0}, Vec{0, 0})
	l2 = NewLine(Vec{-1, 0}, Vec{5, 0})
	assert.True(t, l1.Cross(*l2))
}
