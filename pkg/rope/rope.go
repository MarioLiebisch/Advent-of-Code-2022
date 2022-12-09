package rope

import (
	"aoc2022/pkg/directions"
	"aoc2022/pkg/io"
	"aoc2022/pkg/numbers"
	"strconv"
)

type Knot numbers.Vector2
type Rope []Knot

type RopeInstruction struct {
	dir   rune
	steps int
}

type RopeInstructions []RopeInstruction

func ReadRopeInstruction(file string) RopeInstructions {
	var ris RopeInstructions
	for _, line := range io.ReadLines(file) {
		s, _ := strconv.Atoi(line[2:])
		ris = append(ris, RopeInstruction{dir: rune(line[0]), steps: s})
	}
	return ris
}

func (r *Rope) AdjustTail() *Rope {
	length := len(*r)
	for i := 0; i < length-1; i++ {
		first, second := &(*r)[i], &(*r)[i+1]
		dx, dy := first.X-second.X, first.Y-second.Y
		if dx > 1 {
			second.X++
			if dy > 0 {
				second.Y++
			} else if dy < 0 {
				second.Y--
			}
		} else if dx < -1 {
			second.X--
			if dy > 0 {
				second.Y++
			} else if dy < 0 {
				second.Y--
			}
		} else if dy > 1 {
			second.Y++
			if dx > 0 {
				second.X++
			} else if dx < 0 {
				second.X--
			}
		} else if dy < -1 {
			second.Y--
			if dx > 0 {
				second.X++
			} else if dx < 0 {
				second.X--
			}
		}
	}
	return r
}

func (r *Rope) ApplyInstructions(ris RopeInstructions) int {
	positions := make(map[Knot]bool)
	head := &(*r)[0]
	tail := &(*r)[len(*r)-1]
	for _, ri := range ris {
		for i := 0; i < ri.steps; i++ {
			switch ri.dir {
			case directions.UP:
				head.Y--
			case directions.DOWN:
				head.Y++
			case directions.LEFT:
				head.X--
			case directions.RIGHT:
				head.X++
			}
			r.AdjustTail()
			positions[*tail] = true
		}
	}
	return len(positions)
}
