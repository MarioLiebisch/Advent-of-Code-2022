package numbers

import "sort"

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func SliceMax(array []int) int {
	var max int = 0
	for _, v := range array {
		if v > max {
			max = v
		}
	}
	return max
}

func SliceMaxN(array []int, n int) []int {
	sort.Ints(array)
	return array[len(array)-n:]
}
