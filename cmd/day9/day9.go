package main

import "aoc2022/pkg/rope"

const sample = "./data/sample-9.txt"
const input = "./data/input-9.txt"

func solve1(file string) int {
	r := make(rope.Rope, 2)
	return r.ApplyInstructions(rope.ReadRopeInstruction(file))
}

func solve2(file string) int {
	r := make(rope.Rope, 10)
	return r.ApplyInstructions(rope.ReadRopeInstruction(file))
}

func main() {
	println("Solution 1 - Sample:", solve1(sample))
	println("Solution 1 - Input:", solve1(input))
	println("Solution 2 - Sample:", solve2(sample))
	println("Solution 2 - Input:", solve2(input))
}
