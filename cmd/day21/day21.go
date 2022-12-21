package main

import "aoc2022/pkg/monkeys2"

const sample = "./data/sample-21.txt"
const input = "./data/input-21.txt"

func solve1(file string) int {
	ms := monkeys2.ReadMonkeys(file)
	return ms.Resolve("root")
}

func solve2(file string) int {
	ms := monkeys2.ReadMonkeys(file)
	return ms.ResolveReversed("humn")
}

func main() {
	println("Solution 1 - Sample:", solve1(sample))
	println("Solution 1 - Input:", solve1(input))
	println("Solution 2 - Sample:", solve2(sample))
	println("Solution 2 - Input:", solve2(input))
}
