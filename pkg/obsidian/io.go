package obsidian

import (
	"aoc2022/pkg/io"
	"aoc2022/pkg/numbers"
	"aoc2022/pkg/strings"
	bstrings "strings"
)

func ReadDroplet(file string) *Droplet {
	var res Droplet
	for _, line := range io.ReadLines(file) {
		data := strings.ToIntSlice(bstrings.Split(line, ","))
		var c Cube
		c.X = data[0]
		c.Y = data[1]
		c.Z = data[2]

		if res.Width <= c.X {
			res.Width = c.X + 1
		}
		if res.Height <= c.Y {
			res.Height = c.Y + 1
		}
		if res.Depth <= c.Z {
			res.Depth = c.Z + 1
		}

		c.Surface = 6
		for i, oc := range res.cubes {
			// No surface area left?
			if oc.Surface == 0 {
				continue
			}

			// Check if adjacent
			if numbers.Abs(c.X-oc.X)+numbers.Abs(c.Y-oc.Y)+numbers.Abs(c.Z-oc.Z) == 1 {
				c.Surface--
				res.cubes[i].Surface--
			}

			// Found all surface areas?
			if c.Surface == 0 {
				break
			}
		}
		res.cubes = append(res.cubes, c)
	}
	return &res
}
