package main

import "aoc2022/pkg/forest"

const sample = "./data/sample-8.txt"
const input = "./data/input-8.txt"

func solve1(file string) int {
	return forest.FromFile(file).CountVisibleTrees()
}

func solve2(file string) int {
	return forest.FromFile(file).GetHighestScenicScore()
}

func main() {
	println("Solution 1 - Sample:", solve1(sample))
	println("Solution 1 - Input:", solve1(input))
	println("Solution 2 - Sample:", solve2(sample))
	println("Solution 2 - Input:", solve2(input))
}
