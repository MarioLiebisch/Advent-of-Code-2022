package cargo

import (
	"aoc2022/pkg/containers"
	"aoc2022/pkg/io"
	"strconv"
	"strings"
)

type Cargo struct {
	stacks       [10]containers.RuneStack
	count        int
	instructions CargoInstructions
}

type CargoInstruction struct {
	count, from, to int
}

type CargoInstructions []CargoInstruction

func ReadCargo(file string) Cargo {
	var c Cargo
	var lines = io.ReadLines(file)
	var end_of_stacks_line = 0
	// Find the line with stack indexes
	for i, l := range lines {
		// Bit naive, but will do the trick
		if l[0] == ' ' && l[1] != ' ' {
			end_of_stacks_line = i
			c.count = (len(l) + 1) / 4
			break
		}
	}

	// Now go through all rows in reverse and "stack" them
	for i := end_of_stacks_line - 1; i >= 0; i-- {
		var line = lines[i]
		// Iterate over all stacks inside
		for j := 0; j < c.count; j++ {
			var r = rune(line[1+j*4])
			if r != ' ' {
				c.stacks[j].Push(r)
			}
		}
	}

	// Go through all instructions
	for i := end_of_stacks_line + 2; i < len(lines); i++ {
		tmp := lines[i][5:]
		tmp2 := strings.Split(tmp, " from ")
		tmp3 := strings.Split(tmp2[1], " to ")
		var ci CargoInstruction
		ci.count, _ = strconv.Atoi(tmp2[0])
		ci.from, _ = strconv.Atoi(tmp3[0])
		ci.to, _ = strconv.Atoi(tmp3[1])
		c.instructions = append(c.instructions, ci)
	}
	return c
}

func (c *Cargo) ApplyInstructions() *Cargo {
	for _, ci := range c.instructions {
		for i := 0; i < ci.count; i++ {
			c.stacks[ci.to-1].Push(c.stacks[ci.from-1].Pop())
		}
	}
	return c
}

func (c *Cargo) ApplyInstructions2() *Cargo {
	for _, ci := range c.instructions {
		c.stacks[ci.to-1].PushN(c.stacks[ci.from-1].PopN(ci.count))
	}
	return c
}

func (c *Cargo) ReadTops() string {
	var ret []rune
	for i := 0; i < c.count; i++ {
		ret = append(ret, c.stacks[i].Top())
	}
	return string(ret)
}
