package cave

func (c *Cave) At(x, y int, floor bool) Tile {
	if floor && y > c.floor+1 {
		return Rock
	}
	if r, ok := c.tiles[y]; ok {
		if t, ok := r[x]; ok {
			return t
		}
	}
	return Air
}

func (c *Cave) CountSand() int {
	return c.sand
}

func (c *Cave) DropSand(floor bool) bool {
	sand_x, sand_y := c.entry_x, c.entry_y
	dropped := false
	for {
		// Did it drop out?
		if !floor && (sand_x < c.left || sand_x > c.right || sand_y >= c.floor) {
			return false
		}
		// Drop straight
		if c.At(sand_x, sand_y+1, floor) == Air {
			sand_y++
			dropped = true
			continue
		}
		// Drop to the left
		if c.At(sand_x-1, sand_y+1, floor) == Air {
			sand_x--
			if sand_x < c.left {
				c.left = sand_x
			}
			sand_y++
			dropped = true
			continue
		}
		// Drop to the right
		if c.At(sand_x+1, sand_y+1, floor) == Air {
			sand_x++
			if sand_x > c.right {
				c.right = sand_x
			}
			sand_y++
			dropped = true
			continue
		}
		break
	}
	if c.tiles[sand_y] == nil {
		c.tiles[sand_y] = make(map[int]Tile)
	}
	c.tiles[sand_y][sand_x] = Sand
	c.sand++
	// Only report the sand dropped if it actually did
	return dropped
}
