package fx

// simpleLagrange is a simple implementation of Lagrange interpolation.
// this is for y = a*n^2 + b*n + c
// and for n = [0,1,2] and y = [y0,y1,y2]
// we have a = (y2 - 2*y1 + y0)/2, b = (y2 - y0)/2, c = y1
func simpleLagrange(y0, y1, y2 int) (int, int, int) {
	return (y2 - 2*y1 + y0) / 2, (y2 - y0) / 2, y1
}
