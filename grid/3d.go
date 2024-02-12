package grid

import "golang.org/x/exp/constraints"

type PointUnit interface {
	constraints.Integer | constraints.Float
}

type Point3D[T PointUnit] struct {
	X T
	Y T
	Z T
}

type Vector3D[T PointUnit] struct {
	X T
	Y T
	Z T
}

// Plane3D represents a 3D plane in the form of Ax + By + Cz + D = 0
type Plane3D[T PointUnit] struct {
	A T
	B T
	C T
	D T
}

// NewPlane3D returns a new Plane3D from a point and a normal vector
func NewPlane3D[T PointUnit](p Point3D[T], n Vector3D[T]) Plane3D[T] {
	return Plane3D[T]{
		A: n.X,
		B: n.Y,
		C: n.Z,
		D: -n.X*p.X - n.Y*p.Y - n.Z*p.Z,
	}
}

// NormalVector returns the normal vector of the plane defined by the two points
func NormalVector[T PointUnit](p1, p2 Point3D[T]) Vector3D[T] {
	return Vector3D[T]{
		X: p1.Y*p2.Z - p1.Z*p2.Y,
		Y: p1.Z*p2.X - p1.X*p2.Z,
		Z: p1.X*p2.Y - p1.Y*p2.X,
	}
}

type XYZVec struct {
	X int
	Y int
	Z int
}

func (v XYZVec) Distance(v2 XYZVec) int {
	return Abs(v.X, v2.X) + Abs(v.Y, v2.Y) + Abs(v.Z, v2.Z)
}

func (v XYZVec) Add(v2 XYZVec) XYZVec {
	return XYZVec{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
		Z: v.Z + v2.Z,
	}
}

func (v XYZVec) Neighbors6() []XYZVec {
	return []XYZVec{
		{X: v.X + 1, Y: v.Y, Z: v.Z},
		{X: v.X - 1, Y: v.Y, Z: v.Z},
		{X: v.X, Y: v.Y + 1, Z: v.Z},
		{X: v.X, Y: v.Y - 1, Z: v.Z},
		{X: v.X, Y: v.Y, Z: v.Z + 1},
		{X: v.X, Y: v.Y, Z: v.Z - 1},
	}
}
