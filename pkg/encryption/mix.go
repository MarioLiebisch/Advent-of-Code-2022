package encryption

import (
	"aoc2022/pkg/slices"
)

func mod(a, b int) int {
	// if a is negative, apply the rules for positive a,
	// negate it again, and then add b to get a positive value
	if a < 0 {
		return -(-a % b) + b
	}
	return a % b
}

func (f *File) ApplyKey(key int) {
	for i := range *f {
		(*f)[i].value *= key
	}
}

func (f *File) Mix() {
	l := len(*f) - 1

	for i := range *f {
		for from, e := range *f {
			if i == e.origin {
				to := mod(from+e.value, l)
				slices.Insert(slices.Remove(*f, from), to, e)
				break
			}
		}
	}
}
