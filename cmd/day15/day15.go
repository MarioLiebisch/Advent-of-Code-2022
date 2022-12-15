package main

import "aoc2022/pkg/sensors"

const sample = "./data/sample-15.txt"
const input = "./data/input-15.txt"

func solve1(file string, row int) int {
	t := sensors.Read(file)
	return t.GetCoverageAtRow(row)
}

func solve2(file string, max int) int {
	t := sensors.Read(file)
	return t.GetTuningFrequency(max)
}

func main() {
	println("Solution 1 - Sample:", solve1(sample, 10))
	println("Solution 1 - Input:", solve1(input, 2000000))
	println("Solution 2 - Sample:", solve2(sample, 20))
	println("Solution 2 - Input:", solve2(input, 4000000))
}
