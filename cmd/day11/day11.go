package main

import "aoc2022/pkg/monkeys"

const sample = "./data/sample-11.txt"
const input = "./data/input-11.txt"

func solve1(file string) int {
	m := monkeys.ReadMonkeys(file)
	return m.Process(20, true).GetMonkeyBusiness()
}

func solve2(file string) int {
	m := monkeys.ReadMonkeys(file)
	return m.Process(10000, false).GetMonkeyBusiness()
}

func main() {
	println("Solution 1 - Sample:", solve1(sample))
	println("Solution 1 - Input:", solve1(input))
	println("Solution 2 - Sample:", solve2(sample))
	println("Solution 2 - Input:", solve2(input))
}
