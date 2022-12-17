package tetris

func (c *Chamber) GetHeight() int64 {
	return c.height
}

func (c *Chamber) GetNumJets() int {
	return len(c.jets)
}

func (c *Chamber) Fits(r Rock, x, y int64) bool {
	// Crosses borders?
	if x < 0 || y < 0 || x+r.w > 7 {
		return false
	}
	// Test all parts
	var sx, sy int64
	for sy = 0; sy < r.h; sy++ {
		for sx = 0; sx < r.w; sx++ {
			cx := x + sx           // just add up
			cy := y + r.h - sy - 1 // rock shape is top to bottom, chamber structure bottom to top
			// Is this part solid and is there something in the chamber blocking it?
			if r.shape[sy][sx] == '#' && c.structure[cx][cy] {
				// Then fail
				return false
			}
		}
	}
	// Nothing overlapped, it fits
	return true
}

func (c *Chamber) Place(r Rock, x, y int64) {
	// Test all parts
	var sx, sy int64
	for sy = 0; sy < r.h; sy++ {
		for sx = 0; sx < r.w; sx++ {
			cx := x + sx           // just add up
			cy := y + r.h - sy - 1 // rock shape is top to bottom, chamber structure bottom to top
			// Is this part solid and is there something in the chamber blocking it?
			if r.shape[sy][sx] == '#' {
				c.structure[cx][cy] = true
			}
		}
	}
}
