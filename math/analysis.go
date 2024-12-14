package math

import "math"

// StdDev calculates the standard deviation of a slice of integers.
// to describe the amount of variation or dispersion of a set of values.
// 计算标准差，描述一段数值的变化或分散程度
func StdDev(a []int) float64 {
	mean := 0.0
	for _, v := range a {
		mean += float64(v)
	}
	mean /= float64(len(a))
	variance := 0.0
	for _, v := range a {
		diff := float64(v) - mean
		variance += diff * diff
	}
	return math.Sqrt(variance / float64(len(a)))
}
