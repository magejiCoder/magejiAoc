package matrix

import (
	"github.com/magejiCoder/magejiAoc/grid"
	"golang.org/x/exp/constraints"
)

type insider[T constraints.Ordered] struct {
	cornerPass map[grid.Vec]struct{}
	match      func(Point[T]) bool
}

type insiderOption[T constraints.Ordered] func(*insider[T])

// IsPointInside returns true if the point is inside the matrix
func (m Matrix[T]) IsPointInside(p Point[T], opts ...insiderOption[T]) bool {
	var intersectCount int
	i, j := p.X, p.Y
	var opt insider[T]
	for _, o := range opts {
		o(&opt)
	}
	// 从当前点到左上边界的射线上的交点
	// 1. 不同横竖的射线: 需要计算太多重合的时候方向的case，在这个场景下不如直接用斜线
	// 2. 把转角点看成曲线，则有: 7,L 对于 k = -1 的斜线来说，都是凹点，属于内侧;
	//    F,J 对于 k = -1 的斜线来说，都是凸点，属于外侧;
	//    S 既是凸点也是凹点
	for {
		if i < 0 || j < 0 {
			break
		}
		if opt.cornerPass != nil {
			if _, ok := opt.cornerPass[grid.Vec{X: i, Y: j}]; ok {
				i = i - 1
				j = j - 1
				continue
			}
		}
		if opt.match != nil && opt.match(Point[T]{X: i, Y: j, Value: m.Get(i, j)}) {
			intersectCount++
		}
		i--
		j--
	}
	// 奇数: 在平面中
	// 偶数: 不在平面中
	return intersectCount%2 > 0 && intersectCount > 0
}
