package main

import (
	"aoc2022/pkg/tetris"
)

const sample = "./data/sample-17.txt"
const input = "./data/input-17.txt"

func solve1(file string) int64 {
	c := tetris.ReadChamber(file)
	for i := 0; i < 2022; i++ {
		c.DropRock()
	}
	return c.GetHeight()
}

func solve2(file string) int64 {
	return 0
}

func main() {
	println("Solution 1 - Sample:", solve1(sample))
	println("Solution 1 - Input:", solve1(input))
	println("Solution 2 - Sample:", solve2(sample))
	println("Solution 2 - Input:", solve2(input))
}
