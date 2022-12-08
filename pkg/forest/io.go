package forest

import "aoc2022/pkg/io"

func FromFile(file string) *Forest {
	var f Forest
	for _, line := range io.ReadLines(file) {
		var nl []int
		for _, c := range line {
			nl = append(nl, int(c-'0'))
		}
		f.trees = append(f.trees, nl)
	}
	f.height = len(f.trees)
	if f.height > 0 {
		f.width = len(f.trees[0])
	}
	return &f
}
