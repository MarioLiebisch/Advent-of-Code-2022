package tetris

func (c *Chamber) DropRock() {
	var x, y int64
	x, y = 2, c.height+3
	new_height := y + Rocks[c.next_rock].h
	for {
		var jet int64
		if c.jets[c.next_jet] == '>' {
			jet = 1
		} else {
			jet = -1
		}

		c.next_jet = (c.next_jet + 1) % len(c.jets)
		// Check to see if the jet is able to pus hthe rock
		if c.Fits(Rocks[c.next_rock], x+jet, y) {
			x += jet
		}
		// Check to see if the rock may still drop
		if c.Fits(Rocks[c.next_rock], x, y-1) {
			y--
			new_height--
			continue
		}
		// Rock should settle, if we arrive here
		c.Place(Rocks[c.next_rock], x, y)
		if c.height < new_height {
			c.height = new_height
		}
		c.next_rock = (c.next_rock + 1) % len(Rocks)
		return
	}
}

