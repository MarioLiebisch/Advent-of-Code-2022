package tetris

type Chamber struct {
	height    int64
	structure [7]map[int64]bool
	jets      string
	next_jet  int
	next_rock int
}

type Rock struct {
	w, h  int64
	shape []string
}
