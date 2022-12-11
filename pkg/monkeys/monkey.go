package monkeys

import (
	"aoc2022/pkg/io"
	"aoc2022/pkg/strings"
	"sort"
	"strconv"
	bstrings "strings"
)

const (
	Addition       = 0
	Multiplication = 1
)

type Monkey struct {
	inventory   []int
	operation   int
	operand     int
	test        int
	throw_true  int
	throw_false int
	inspected   int
}

type Monkeys [8]Monkey

func ReadMonkeys(file string) Monkeys {
	var res Monkeys
	var mi int
	for _, line := range io.ReadLines(file) {
		if line == "" {
			continue
		}
		if line[0:6] == "Monkey" {
			mi, _ = strconv.Atoi(line[7:8])
		} else {
			switch line[2:6] {
			case "Star":
				res[mi].inventory = strings.ToIntSlice(bstrings.Split(line[18:], ", "))
			case "Oper":
				if line[23] == '+' {
					res[mi].operation = Addition
				} else {
					res[mi].operation = Multiplication
				}
				res[mi].operand, _ = strconv.Atoi(line[25:])
			case "Test":
				res[mi].test, _ = strconv.Atoi(line[21:])
			case "  If":
				if line[7] == 't' {
					res[mi].throw_true, _ = strconv.Atoi(line[29:])
				} else {
					res[mi].throw_false, _ = strconv.Atoi(line[30:])
				}
			}
		}
	}
	return res
}

func (m *Monkeys) GetMonkeyBusiness() int {
	var mb []int
	for _, m := range *m {
		mb = append(mb, m.inspected)
	}
	sort.Ints(mb)
	l := len(mb)
	return mb[l-2] * mb[l-1]
}

func (m *Monkeys) Process(rounds int, decrease bool) *Monkeys {
	// Fake the most common denominator by just multiplying all test values together
	common := 1
	for i := 0; i < 8; i++ {
		t := (*m)[i].test
		if t != 0 {
			if common%t == 0 {
				continue
			}
			common *= t
		}
	}

	for r := 0; r < rounds; r++ {
		for i := 0; i < 8; i++ {
			monkey := &(*m)[i]
			for j := 0; j < len(monkey.inventory); j++ {
				monkey.inspected++
				worry := monkey.inventory[j]
				if monkey.operation == Addition {
					if monkey.operand == 0 {
						worry *= 2
					} else {
						worry += monkey.operand
					}
				} else {
					if monkey.operand == 0 {
						worry *= worry
					} else {
						worry *= monkey.operand
					}
				}
				if decrease {
					worry /= 3
				} else {
					worry %= common
				}
				tmi := monkey.throw_false
				if worry%monkey.test == 0 {
					tmi = monkey.throw_true
				}
				(*m)[tmi].inventory = append((*m)[tmi].inventory, worry)
			}
			monkey.inventory = make([]int, 0)
		}
	}
	return m
}
