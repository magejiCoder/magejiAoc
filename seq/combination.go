package seq

func Combination[T any](seq []T, n int) [][]T {
	if n == 0 {
		return [][]T{{}}
	}
	if n == 1 {
		result := make([][]T, len(seq))
		for i, e := range seq {
			result[i] = []T{e}
		}
		return result
	}
	if n == len(seq) {
		return [][]T{seq}
	}
	result := [][]T{}
	for i, e := range seq {
		sub := Combination(seq[i+1:], n-1)
		for _, s := range sub {
			result = append(result, append([]T{e}, s...))
		}
	}
	return result
}
