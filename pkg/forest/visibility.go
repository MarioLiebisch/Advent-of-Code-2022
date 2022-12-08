package forest

// Constants for debug output
const Red = "\033[31m"
const Green = "\033[32m"

func (f *Forest) CountVisibleTrees() int {
	res := make(map[int]bool, 0)
	// Corners
	res[0] = true
	res[f.width-1] = true
	res[(f.height-1)*1000] = true
	res[(f.height-1)*(f.width-1)] = true

	// From top
	for x := 1; x < f.width-1; x++ {
		last := f.trees[0][x]
		res[x] = true
		for y := 1; y < f.height-1; y++ {
			if f.trees[y][x] > last {
				res[y*1000+x] = true
				last = f.trees[y][x]
			}
		}
	}

	// From bottom
	for x := 1; x < f.width-1; x++ {
		last := f.trees[f.height-1][x]
		res[(f.height-1)*1000+x] = true
		for y := f.height - 1; y > 0; y-- {
			if f.trees[y][x] > last {
				res[y*1000+x] = true
				last = f.trees[y][x]
			}
		}
	}

	// From left
	for y := 1; y < f.height-1; y++ {
		last := f.trees[y][0]
		res[y*1000] = true
		for x := 1; x < f.width-1; x++ {
			if f.trees[y][x] > last {
				res[y*1000+x] = true
				last = f.trees[y][x]
			}
		}
	}

	// From right
	for y := 1; y < f.height-1; y++ {
		last := f.trees[y][f.width-1]
		res[y*1000+f.width-1] = true
		for x := f.width - 1; x > 0; x-- {
			if f.trees[y][x] > last {
				res[y*1000+x] = true
				last = f.trees[y][x]
			}
		}
	}

	// Debug output
	/*
		for y := 0; y < f.height; y++ {
			for x := 0; x < f.width; x++ {
				if res[y*1000+x] {
					print(Green + fmt.Sprint(f.trees[y][x]))
				} else {
					print(Red + fmt.Sprint(f.trees[y][x]))
				}
			}
			println()
		}
	*/

	// Count the trees visible inside, plus account for the border, minus duplicated edges
	return len(res)
}

func (f *Forest) GetHighestScenicScore() int {
	max := 0
	for y := 0; y < f.height; y++ {
		for x := 0; x < f.width; x++ {
			score := f.GetScenicScore(x, y)
			if score > max {
				max = score
			}
		}
	}
	return max
}

func (f *Forest) GetScenicScore(x, y int) int {
	var top, bottom, left, right int
	// to the top
	center := f.trees[y][x]
	highest := 0
	for ty := y - 1; ty >= 0; ty-- {
		top++
		current := f.trees[ty][x]
		if highest <= current {
			highest = current
		}
		if center <= current {
			break
		}
	}
	// to the bottom
	highest = 0
	for ty := y + 1; ty <= f.height-1; ty++ {
		bottom++
		current := f.trees[ty][x]
		if highest <= current {
			highest = current
		}
		if center <= current {
			break
		}
	}
	// to the left
	highest = 0
	for tx := x - 1; tx >= 0; tx-- {
		left++
		current := f.trees[y][tx]
		if highest <= current {
			highest = current
		}
		if center <= current {
			break
		}
	}
	// to the right
	highest = 0
	for tx := x + 1; tx <= f.width-1; tx++ {
		right++
		current := f.trees[y][tx]
		if highest <= current {
			highest = current
		}
		if center <= current {
			break
		}
	}
	return top * bottom * left * right
}
