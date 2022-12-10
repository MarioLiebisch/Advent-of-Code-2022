package main

import "aoc2022/pkg/cpu"

const sample = "./data/sample-10.txt"
const input = "./data/input-10.txt"

func solve1(file string) int {
	c := cpu.NewCPU()
	c.Run(cpu.ReadProgram(file), 1)
	return c.GetSignalStrength()
}

func solve2(file string) string {
	c := cpu.NewCPU()
	c.Run(cpu.ReadProgram(file), 1)
	return c.GetScreen()
}

func main() {
	println("Solution 1 - Sample:", solve1(sample))
	println("Solution 1 - Input:", solve1(input))
	println("Solution 2 - Sample:", solve2(sample))
	println("Solution 2 - Input:", solve2(input))
}
