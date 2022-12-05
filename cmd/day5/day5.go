package main

import "aoc2022/pkg/cargo"

const sample = "./data/sample-5.txt"
const input = "./data/input-5.txt"

func solve1(file string) string {
	var c = cargo.ReadCargo(file)
	return c.ApplyInstructions().ReadTops()
}

func solve2(file string) string {
	var c = cargo.ReadCargo(file)
	return c.ApplyInstructions2().ReadTops()
}

func main() {
	println("Solution 1 - Sample:", solve1(sample))
	println("Solution 1 - Input:", solve1(input))
	println("Solution 2 - Sample:", solve2(sample))
	println("Solution 2 - Input:", solve2(input))
}
