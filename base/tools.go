package base

func Min[T byte | int | int64 | float32 | float64](x T, y T) (min T) {
	min = x
	if x > y {
		min = y
	}
	return
}

func Max[T byte | int | int64 | float32 | float64](x T, y T) (max T) {
	max = x
	if x < y {
		max = y
	}
	return
}
