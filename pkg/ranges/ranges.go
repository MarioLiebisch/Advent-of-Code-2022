package ranges

import (
	"aoc2022/pkg/numbers"

	"golang.org/x/exp/slices"
)

type Range struct {
	Min, Max int
}

type Ranges []Range

func (rs *Ranges) Merge() {
	rs.Sort()
	count := len(*rs)
	for i := 0; i < count; i++ {
		for j := i + 1; j < count; {
			r1 := &(*rs)[i]
			r2 := &(*rs)[j]
			// Full overlap
			if r2.Min <= r1.Max+1 {
				if r2.Max > r1.Max {
					r1.Max = r2.Max
				}
				*rs = append((*rs)[:j], (*rs)[j+1:]...)
				count--
				continue
			}
			j++
		}
	}
}

func (rs *Ranges) Sort() *Ranges {
	slices.SortFunc(*rs, func(a, b Range) bool { return a.Min < b.Min })
	return rs
}

func (rs *Ranges) Invert(from, to int) *Ranges {
	rs.Sort()
	var res Ranges
	pos := 0
	tto := to
	for _, r := range *rs {
		tto = to
		if pos < r.Min {
			tto = numbers.Min(tto, r.Min-1)
			res = append(res, Range{pos, tto})
		} else {
			pos = r.Max + 1
		}
	}
	if pos < tto {
		res = append(res, Range{Min: pos, Max: tto})
	}
	return &res
}

func (r *Range) Count() int {
	return r.Max - r.Min
}

func (rs *Ranges) Count() int {
	sum := 0
	for _, r := range *rs {
		sum += r.Count()
	}
	return sum
}

func (rs *Ranges) Includes(value int) bool {
	for _, r := range *rs {
		if r.Min <= value && r.Max >= value {
			return true
		}
	}
	return false
}
