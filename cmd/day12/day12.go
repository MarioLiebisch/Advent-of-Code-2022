package main

import (
	"aoc2022/pkg/heightmap"
)

const sample = "./data/sample-12.txt"
const input = "./data/input-12.txt"

func solve1(file string) int {
	hm := heightmap.Read(file)
	path, _, _ := hm.FindPath()
	return len(path) - 1
}

func solve2(file string) int {
	hm := heightmap.Read(file)
	best := 0
	for _, s := range hm.FindStarts(0) {
		path, _, found := hm.FindPathAlt(s)
		if found && (best == 0 || len(path) < best) {
			best = len(path)
		}
	}
	return best - 1
}

func main() {
	println("Solution 1 - Sample:", solve1(sample))
	println("Solution 1 - Input:", solve1(input))
	println("Solution 2 - Sample:", solve2(sample))
	println("Solution 2 - Input:", solve2(input))
}
