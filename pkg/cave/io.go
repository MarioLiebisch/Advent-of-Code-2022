package cave

import (
	"aoc2022/pkg/io"
	"aoc2022/pkg/numbers"
	"strconv"
	"strings"
)

func Read(file string) *Cave {
	var ret Cave
	var last_x, last_y int

	ret.entry_x, ret.entry_y = 500, 0
	ret.tiles = make(map[int]map[int]Tile)
	ret.tiles[ret.entry_y] = make(map[int]Tile)
	ret.tiles[ret.entry_y][ret.entry_x] = Opening
	ret.ceiling = 0
	ret.floor = 0
	ret.left = 500
	ret.right = 500

	for _, line := range io.ReadLines(file) {
		first := true
		for _, data := range strings.Split(line, " -> ") {
			vdata := strings.Split(data, ",")
			x, _ := strconv.Atoi(vdata[0])
			y, _ := strconv.Atoi(vdata[1])
			if first {
				first = false
			} else {
				for yi := numbers.Min(y, last_y); yi <= numbers.Max(y, last_y); yi++ {
					if ret.tiles[yi] == nil {
						ret.tiles[yi] = make(map[int]Tile)
					}
					for xi := numbers.Min(x, last_x); xi <= numbers.Max(x, last_x); xi++ {
						ret.tiles[yi][xi] = Rock
					}
				}
			}
			last_x = x
			last_y = y
			ret.floor = numbers.Max(ret.floor, y)
			ret.left = numbers.Min(ret.left, x)
			ret.right = numbers.Max(ret.right, x)
		}
	}
	return &ret
}

func (c *Cave) Print() {
	for y := c.ceiling; y <= c.floor; y++ {
		if r, ok := c.tiles[y]; ok {
			for x := c.left; x <= c.right; x++ {
				if v, ok := r[x]; ok {
					print(string(v))
				} else {
					print(string(Air))
				}
			}
			println()
		} else { // Print a line of air not in the data
			println(strings.Repeat(string(Air), c.right-c.left+1))
		}
	}
}
