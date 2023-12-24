package geometry

import (
	"github.com/magejiCoder/magejiAoc/math"
	"github.com/magejiCoder/magejiAoc/matrix"
)

// Area return the area of the polygon
// uses [shoelace formula](https://en.wikipedia.org/wiki/Shoelace_formula)
func Area(p []matrix.Point[byte]) int {
	var sum int
	for i := 0; i < len(p)-1; i++ {
		sum += p[i].X*p[i+1].Y - p[i].Y*p[i+1].X
	}
	sum += p[len(p)-1].X*p[0].Y - p[len(p)-1].Y*p[0].X
	return sum / 2
}

func Perimeter(p []matrix.Point[byte]) int {
	var sum int
	for i := 0; i < len(p)-1; i++ {
		sum += math.Abs(p[i].X-p[i+1].X) + math.Abs(p[i].Y-p[i+1].Y)
	}
	sum += math.Abs(p[len(p)-1].X-p[0].X) + math.Abs(p[len(p)-1].Y-p[0].Y)
	return sum
}

// count return the number of points inside the polygon
// uses [pick's theorem](https://en.wikipedia.org/wiki/Pick%27s_theorem)
func Count(area int, perimeter int) int {
	if area-perimeter/2+1 > 0 {
		return area - perimeter/2 + 1
	}
	return math.Abs(area - perimeter/2 - 1)
}
