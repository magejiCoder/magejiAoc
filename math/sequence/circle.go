package sequence

import (
	"errors"
	"fmt"

	"golang.org/x/exp/constraints"
)

var ErrNoCircle = errors.New("no circle found")

func IsCircle[T constraints.Integer](seq []T) bool {
	return NewCircleSeq(seq).IsCircle(false)
}

func IsCircleStable[T constraints.Integer](seq []T) bool {
	return NewCircleSeq(seq).IsCircle(true)
}

func NewCircleSeq[T constraints.Integer](seq []T) *circleSeq[T] {
	return &circleSeq[T]{seq: seq}
}

type circleSeq[T constraints.Integer] struct {
	seq []T
}

// findCircle uses **Floyd's tortoise and hare** algorithm (fixed to avoid the one-step trap).
func (c *circleSeq[T]) FindCircle(isStrict bool) (int, int, error) {
	if len(c.seq) < 2 {
		return -1, 0, ErrNoCircle
	}
	x0, x1 := -1, -1
	for i := 0; i < len(c.seq); i++ {
		for j := i + 1; j < len(c.seq); j++ {
			if c.seq[i] == c.seq[j] {
				if isStrict {
					if j+1 < len(c.seq) && c.seq[i+1] == c.seq[j+1] {
						x0, x1 = i, j
						break
					}
				} else {
					x0, x1 = i, j
					break
				}
			}
		}
	}
	if x0 == -1 || x1 == -1 {
		return -1, -1, ErrNoCircle
	}
	return x0, x1 - x0, nil
}

func (c *circleSeq[T]) IsCircle(strict bool) bool {
	_, _, err := c.FindCircle(strict)
	return err == nil
}

func (c *circleSeq[T]) At(i int) (int, error) {
	start, gap, err := c.FindCircle(true)
	if err != nil {
		return 0, fmt.Errorf("at index %d: %w", i, err)
	}
	if i < start {
		return int(c.seq[i]), nil
	}
	return int(c.seq[start+(i-start)%gap]), nil
}
