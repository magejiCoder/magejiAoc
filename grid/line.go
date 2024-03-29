package grid

import (
	"sort"
)

type Line struct {
	from Vec
	to   Vec
}

func NewLine(v1, v2 Vec) *Line {
	return &Line{from: v1, to: v2}
}

func (l *Line) Vec() Vec {
	return Vec{l.to.X - l.from.X, l.to.Y - l.from.Y}
}

func (l *Line) Points() []Point {
	v := l.Vec()
	// k == 0
	if v.X == 0 {
		if v.Y < 0 {
			v = v.Neg()
		}
		ps := make([]Point, v.Y+1)
		for i := 0; i <= v.Y; i++ {
			ps[i] = Point{l.from.X, l.from.Y + i}
		}
		return ps
	}
	// k == 0
	if v.Y == 0 {
		if v.X < 0 {
			v = v.Neg()
		}
		ps := make([]Point, v.X+1)
		for i := 0; i <= v.X; i++ {
			ps[i] = Point{l.from.X + i, l.from.Y}
		}
		return ps
	}
	// k == 1
	if v.X == v.Y {
		if v.X < 0 {
			v = v.Neg()
		}
		ps := make([]Point, v.X+1)
		for i := 0; i <= v.X; i++ {
			ps[i] = Point{l.from.X + i, l.from.Y + i}
		}
		return ps
	}
	panic("not a line")
}

func Multiply(v1 Vec, v2 Vec) int {
	return v1.X*v2.Y - v1.Y*v2.X
}

func (l Line) Cross(l2 Line) bool {
	v1 := l.Vec()
	v2 := l2.Vec()
	// v1 is point , v2 is line, v1 is on the line
	if v1.X == 0 && v1.Y == 0 {
		return l2.from.X <= l.from.X && l.from.X <= l2.to.X &&
			l2.from.Y <= l.from.Y && l.from.Y <= l2.to.Y
	}
	// v2 is point , v1 is line
	if v2.X == 0 && v2.Y == 0 {
		return l.from.X <= l2.from.X && l2.from.X <= l.to.X &&
			l.from.Y <= l2.from.Y && l2.from.Y <= l.to.Y
	}
	return Multiply(v1, l2.from.Sub(l.from))*Multiply(v1, l2.to.Sub(l.from)) <= 0 &&
		Multiply(v2, l.from.Sub(l2.from))*Multiply(v2, l.to.Sub(l2.from)) <= 0
}

func (l *Line) ToX() int {
	return l.to.X
}

func (l *HorizontalLine) Print() {
	println(l.from.X, l.to.X)
}

func (l *Line) OnDraw(fn func(v Vec) error) error {
	if l.from.X == l.to.X {
		if l.from.Y > l.to.Y {
			l.from, l.to = l.to, l.from
		}
		for i := l.from.Y; i <= l.to.Y; i++ {
			if err := fn(Vec{l.from.X, i}); err != nil {
				return err
			}
		}
	} else {
		if l.from.X > l.to.X {
			l.from, l.to = l.to, l.from
		}
		for i := l.from.X; i <= l.to.X; i++ {
			if err := fn(Vec{i, l.from.Y}); err != nil {
				return err
			}
		}
	}
	return nil
}

type HorizontalLine struct {
	*Line
}

func NewHorizontalLine(x0, x1, y int) *HorizontalLine {
	return &HorizontalLine{NewLine(Vec{x0, y}, Vec{x1, y})}
}

func (l *HorizontalLine) Add(l2 *HorizontalLine) []*HorizontalLine {
	if l.from.Y != l2.from.Y {
		panic("not horizontal")
	}
	if l2.from.X > l.to.X+1 {
		return []*HorizontalLine{l, l2}
	}
	if l.from.X > l2.to.X+1 {
		return []*HorizontalLine{l2, l}
	}
	if l.from.X > l2.from.X {
		l.from.X = l2.from.X
	}
	if l.to.X < l2.to.X {
		l.to.X = l2.to.X
	}
	return []*HorizontalLine{l}
}

func (l *HorizontalLine) Len() int {
	return l.to.X - l.from.X + 1
}

type HorizontalLines []*HorizontalLine

func (hls HorizontalLines) Len() int {
	return len(hls)
}

func (hls HorizontalLines) Less(i, j int) bool {
	return hls[i].from.X < hls[j].from.X
}

func (hls HorizontalLines) Swap(i, j int) {
	hls[i], hls[j] = hls[j], hls[i]
}

func Merge(hls HorizontalLines) HorizontalLines {
	if len(hls) == 1 {
		return hls
	}
	sort.Sort(hls)
	merged := []*HorizontalLine{hls[0]}
	res := []*HorizontalLine{}
	for i := 1; i < len(hls); i++ {
		if i == len(hls)-1 {
			res = append(res, merged...)
		}
		merged = merged[len(merged)-1].Add(hls[i])
		if len(merged) == 2 {
			res = append(res, merged[0])
			merged = merged[1:]
		}
	}

	return res
}
