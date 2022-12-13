package distress

import "aoc2022/pkg/numbers"

func Compare(p1, p2 *Package) int {
	// Both are integers
	if p1.value != -1 && p2.value != -1 {
		if p1.value < p2.value {
			return -1
		} else if p1.value == p2.value {
			return 0
		}
		return 1
		// Only p1 is an integer
	} else if p1.value != -1 {
		var tmp Package
		tmp.value = -1
		tmp.contents = make([]*Package, 1)
		tmp.contents[0] = p1
		return Compare(&tmp, p2)
		// Only p2 is an integer
	} else if p2.value != -1 {
		var tmp Package
		tmp.value = -1
		tmp.contents = make([]*Package, 1)
		tmp.contents[0] = p2
		return Compare(p1, &tmp)
	}
	// Both are lists
	l1, l2 := len(p1.contents), len(p2.contents)
	l := numbers.Min(l1, l2)
	for i := 0; i < l; i++ {
		r := Compare(p1.contents[i], p2.contents[i])
		if r != 0 {
			return r
		}
	}
	if l1 < l2 {
		// p1 runs out of items first
		return -1
	} else if l1 > l2 {
		// p2 runs out of items first
		return 1
	}
	// Both are identical
	return 0
}
