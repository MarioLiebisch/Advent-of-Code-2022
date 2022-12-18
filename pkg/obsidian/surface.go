package obsidian

func (d *Droplet) GetSurface() int {
	res := 0
	for _, c := range *&d.cubes {
		res += c.Surface
	}
	return res
}

type model []bool

func (v *model) floodfill(x, y, z, w, h, d int) {
	// Is this cell outside borders?
	if x < 0 || y < 0 || z < 0 || x >= w || y >= h || z >= d {
		return
	}

	i := (x*h+y)*d + z
	// Is this cell filled?
	if (*v)[i] {
		return
	}
	// Fill it
	(*v)[i] = true
	// Fill neighbours
	v.floodfill(x-1, y, z, w, h, d)
	v.floodfill(x+1, y, z, w, h, d)
	v.floodfill(x, y-1, z, w, h, d)
	v.floodfill(x, y+1, z, w, h, d)
	v.floodfill(x, y, z-1, w, h, d)
	v.floodfill(x, y, z+1, w, h, d)
}

// Bit naive approach, but going with it…
func (d *Droplet) GetOutsideSurface() int {
	// Create a model where all cells that aren't inside are true
	v := make(model, d.Width*d.Height*d.Depth)
	for _, c := range d.cubes {
		v[(c.X*d.Height+c.Y)*d.Depth+c.Z] = true
	}
	// Go over all surfaces and floodfill from there
	// Left and right
	for z := 0; z < d.Depth; z++ {
		for y := 0; y < d.Height; y++ {
			v.floodfill(0, y, z, d.Width, d.Height, d.Depth)
			v.floodfill(d.Width-1, y, z, d.Width, d.Height, d.Depth)
		}
	}
	// Front and back
	for x := 0; x < d.Width; x++ {
		for y := 0; y < d.Height; y++ {
			v.floodfill(x, y, 0, d.Width, d.Height, d.Depth)
			v.floodfill(x, y, d.Depth-1, d.Width, d.Height, d.Depth)
		}
	}
	// Top and bottom
	for z := 0; z < d.Depth; z++ {
		for x := 0; x < d.Width; x++ {
			v.floodfill(x, 0, z, d.Width, d.Height, d.Depth)
			v.floodfill(x, d.Height-1, z, d.Width, d.Height, d.Depth)
		}
	}

	// Now count the surface area of the enclosed areas
	insideSurface := 0
	for z := 0; z < d.Depth; z++ {
		for y := 0; y < d.Height; y++ {
			for x := 0; x < d.Width; x++ {
				if !v[(x*d.Height+y)*d.Depth+z] {
					// Get this cube's surface area
					if x > 0 && v[((x-1)*d.Height+y)*d.Depth+z] {
						insideSurface++
					}
					if x < d.Width-1 && v[((x+1)*d.Height+y)*d.Depth+z] {
						insideSurface++
					}
					if y > 0 && v[(x*d.Height+y-1)*d.Depth+z] {
						insideSurface++
					}
					if y < d.Height-1 && v[(x*d.Height+y+1)*d.Depth+z] {
						insideSurface++
					}
					if z > 0 && v[(x*d.Height+y)*d.Depth+z-1] {
						insideSurface++
					}
					if z < d.Depth-1 && v[(x*d.Height+y)*d.Depth+z+1] {
						insideSurface++
					}
				}
			}
		}
	}
	// And just subtract it from the total surface area
	return d.GetSurface() - insideSurface
}
