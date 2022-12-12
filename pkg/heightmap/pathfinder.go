package heightmap

import (
	"aoc2022/pkg/numbers"

	astar "github.com/beefsack/go-astar"
)

func (t *HeightTile) PathNeighbors() []astar.Pather {
	res := make([]astar.Pather, 0)
	// Not crossing the border
	if t.X > 0 {
		o := &t.hm.data[t.Y][t.X-1]
		// Heigh difference is okay
		if o.H <= t.H+1 {
			// Add it
			res = append(res, o)
		}
	}
	if t.X < t.hm.width-1 {
		o := &t.hm.data[t.Y][t.X+1]
		if o.H <= t.H+1 {
			res = append(res, o)
		}
	}
	if t.Y > 0 {
		o := &t.hm.data[t.Y-1][t.X]
		if o.H <= t.H+1 {
			res = append(res, o)
		}
	}
	if t.Y < t.hm.height-1 {
		o := &t.hm.data[t.Y+1][t.X]
		if o.H <= t.H+1 {
			res = append(res, o)
		}
	}
	return res
}

func (t *HeightTile) PathNeighborCost(to astar.Pather) float64 {
	// All tiles have a constant movement cost
	return 1
}

func (t *HeightTile) PathEstimatedCost(to astar.Pather) float64 {
	tto := to.(*HeightTile)
	// Estimate using Manhattan Distance
	return float64(numbers.Abs(tto.X-t.X) + numbers.Abs(tto.Y-t.Y))
}

func (hm *HeightMap) FindPath() ([]astar.Pather, float64, bool) {
	return astar.Path(hm.start, hm.end)
}

func (hm *HeightMap) FindStarts(h int) []astar.Pather {
	res := make([]astar.Pather, 0)
	for i := 0; i < hm.width*hm.height; i++ {
		t := &hm.data[i/hm.width][i%hm.width]
		if t.H == h {
			res = append(res, t)
		}
	}
	return res
}

func (hm *HeightMap) FindPathAlt(start astar.Pather) ([]astar.Pather, float64, bool) {
	return astar.Path(start, hm.end)
}
