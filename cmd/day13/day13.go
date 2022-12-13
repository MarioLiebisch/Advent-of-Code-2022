package main

import "aoc2022/pkg/distress"

const sample = "./data/sample-13.txt"
const input = "./data/input-13.txt"

func solve1(file string) int {
	s := distress.ReadSignal(file)
	return s.VerifyOrder()
}

func solve2(file string) int {
	s := distress.ReadSignal(file)
	s.AddDivider()
	s.Sort()
	return s.GetDecoderKey()
}

func main() {
	println("Solution 1 - Sample:", solve1(sample))
	println("Solution 1 - Input:", solve1(input))
	println("Solution 2 - Sample:", solve2(sample))
	println("Solution 2 - Input:", solve2(input))
}
