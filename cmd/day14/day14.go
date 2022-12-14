package main

import "aoc2022/pkg/cave"

const sample = "./data/sample-14.txt"
const input = "./data/input-14.txt"

func solve1(file string) int {
	c := cave.Read(file)
	for c.DropSand(false) {
	}
	c.Print()
	return c.CountSand()
}

func solve2(file string) int {
	c := cave.Read(file)
	for c.DropSand(true) {
	}
	c.Print()
	return c.CountSand()
}

func main() {
	println("Solution 1 - Sample:", solve1(sample))
	println("Solution 1 - Input:", solve1(input))
	println("Solution 2 - Sample:", solve2(sample))
	println("Solution 2 - Input:", solve2(input))
}
