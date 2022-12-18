package main

import "aoc2022/pkg/obsidian"

const sample = "./data/sample-18.txt"
const input = "./data/input-18.txt"

func solve1(file string) int {
	d := obsidian.ReadDroplet(file)
	return d.GetSurface()
}

func solve2(file string) int {
	d := obsidian.ReadDroplet(file)
	return d.GetOutsideSurface()
}

func main() {
	println("Solution 1 - Sample:", solve1(sample))
	println("Solution 1 - Input:", solve1(input))
	println("Solution 2 - Sample:", solve2(sample))
	println("Solution 2 - Input:", solve2(input))
}
