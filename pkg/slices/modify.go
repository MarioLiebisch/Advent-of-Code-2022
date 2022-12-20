package slices

func Insert[T any](slice []T, pos int, value T) []T {
	return append(slice[:pos], append([]T{value}, slice[pos:]...)...)
}

func Remove[T any](slice []T, pos int) []T {
	return append(slice[:pos], slice[pos+1:]...)
}

func Move[T any](slice []T, src, dst int) []T {
	value := slice[src]
	l := len(slice) - 1

	if dst >= l {
		dst %= l
	}

	for dst < 0 {
		dst += l
	}

	return Insert(Remove(slice, src), dst, value)
}

func MoveBy[T any](slice []T, src, distance int) []T {
	return Move(slice, src, src+distance)
}
