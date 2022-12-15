package sensors

import (
	"aoc2022/pkg/numbers"
	"aoc2022/pkg/ranges"
)

func (t *Tunnels) GetCoverageRangesAtRow(row int) ranges.Ranges {
	var coverage ranges.Ranges
	for _, s := range t.sensors {
		r := s.GetCoverageAtRow(row)
		if r != nil {
			coverage = append(coverage, *r)
		}
	}
	coverage.Merge()
	return coverage
}

func (t *Tunnels) GetCoverageRangesAtColumn(col int) ranges.Ranges {
	var coverage ranges.Ranges
	for _, s := range t.sensors {
		r := s.GetCoverageAtColumn(col)
		if r != nil {
			coverage = append(coverage, *r)
		}
	}
	coverage.Merge()
	return coverage
}

func (t *Tunnels) GetCoverageAtRow(row int) int {
	c := t.GetCoverageRangesAtRow(row)
	return c.Count()
}

func (t *Tunnels) GetCoverageAtColumn(col int) int {
	c := t.GetCoverageRangesAtColumn(col)
	return c.Count()
}

func (s *Sensor) GetCoverageAtRow(row int) *ranges.Range {
	var res ranges.Range
	y_dist := numbers.Abs(row - s.Y)
	x_dist := s.R - y_dist
	if x_dist <= 0 {
		return nil
	}
	res.Min = s.X - x_dist
	res.Max = s.X + x_dist
	return &res
}

func (s *Sensor) GetCoverageAtColumn(col int) *ranges.Range {
	var res ranges.Range
	x_dist := numbers.Abs(col - s.X)
	y_dist := s.R - x_dist
	if y_dist <= 0 {
		return nil
	}
	res.Min = s.Y - y_dist
	res.Max = s.Y + y_dist
	return &res
}

func (t *Tunnels) GetTuningFrequency(max int) int {
	for i := 0; i <= max; i++ {
		// Determine the coverage for this row
		xrs := t.GetCoverageRangesAtRow(i)
		// Invert it
		xrs = *xrs.Invert(0, max)
		for _, xr := range xrs {
			// Now only try all columns not covered
			for x := xr.Min; x <= xr.Max; x++ {
				// Determine the coverage for this column
				yrs := t.GetCoverageRangesAtColumn(x)
				// Invert it
				yrs = *yrs.Invert(0, max)
				// Does it include our original value (i.e. assumed y)?
				if yrs.Includes(i) {
					// This is a hit
					return x*4000000 + i
				}
			}
		}
	}
	// Shouldn't happen
	return 0
}
