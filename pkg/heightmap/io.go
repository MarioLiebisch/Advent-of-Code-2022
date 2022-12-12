package heightmap

import "aoc2022/pkg/io"

func Read(file string) *HeightMap {
	var res HeightMap
	lines := io.ReadLines(file)
	res.height = len(lines)
	if res.height > 0 {
		res.width = len(lines[0])
		res.data = make([][]HeightTile, res.height)
		for y, line := range lines {
			res.data[y] = make([]HeightTile, res.width)
			for x, r := range line {
				h := int(r) - int('a')
				if r == 'S' {
					h = 0
					res.start = &res.data[y][x]
				} else if r == 'E' {
					h = 25
					res.end = &res.data[y][x]
				}
				res.data[y][x].X = x
				res.data[y][x].Y = y
				res.data[y][x].H = h
				res.data[y][x].hm = &res
			}
		}
	}
	return &res
}
