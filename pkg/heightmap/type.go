package heightmap

type HeightTile struct {
	X, Y int
	H    int
	hm   *HeightMap
}

type HeightMap struct {
	data          [][]HeightTile
	width, height int
	start, end    *HeightTile
}
