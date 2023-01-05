package utilities

func MaxMin[T float64 | int | uint](value, max, min T) T {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
