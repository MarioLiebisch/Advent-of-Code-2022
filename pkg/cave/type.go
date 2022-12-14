package cave

const (
	Air     = '.'
	Rock    = '#'
	Sand    = 'o'
	Opening = '+'
)

type Tile rune

type Cave struct {
	tiles                       map[int]map[int]Tile
	floor, ceiling, left, right int
	entry_x, entry_y            int
	sand                        int
}
