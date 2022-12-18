package obsidian

type Cube struct {
	X, Y, Z, Surface int
}

type Droplet struct {
	cubes                []Cube
	Width, Height, Depth int
}
