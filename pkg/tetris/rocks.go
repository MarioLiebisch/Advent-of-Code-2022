package tetris

var Rocks []Rock

func init() {
	Rocks = make([]Rock, 5)
	Rocks[0] = Rock{
		w: 4, h: 1,
		shape: []string{
			"####"},
	}
	Rocks[1] = Rock{
		w: 3, h: 3,
		shape: []string{
			".#.",
			"###",
			".#."},
	}
	Rocks[2] = Rock{
		w: 3, h: 3,
		shape: []string{
			"..#",
			",,#",
			"###"},
	}
	Rocks[3] = Rock{
		w: 1, h: 4,
		shape: []string{
			"#",
			"#",
			"#",
			"#"},
	}
	Rocks[4] = Rock{
		w: 2, h: 2,
		shape: []string{
			"##",
			"##"},
	}
}
