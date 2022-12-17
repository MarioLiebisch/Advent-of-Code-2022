package tetris

import "aoc2022/pkg/io"

func ReadChamber(file string) *Chamber {
	var res Chamber
	res.jets = io.ReadLines(file)[0]
	for i := 0; i < 7; i++ {
		res.structure[i] = make(map[int64]bool)
	}
	return &res
}

func (c *Chamber) Print() {
	for y := c.height - 1; y >= 0; y-- {
		for x := 0; x < 7; x++ {
			if c.structure[x][y] {
				print("#")
			} else {
				print(".")
			}
		}
		println()
	}
	println()
}
