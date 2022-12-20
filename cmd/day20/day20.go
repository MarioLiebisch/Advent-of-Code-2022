package main

import "aoc2022/pkg/encryption"

const sample = "./data/sample-20.txt"
const input = "./data/input-20.txt"

func solve1(file string) int {
	f := encryption.ReadFile(file)
	f.Mix()
	return f.GetGroveCoordinates()
}

func solve2(file string) int {
	f := encryption.ReadFile(file)
	f.ApplyKey(811589153)
	for i := 0; i < 10; i++ {
		f.Mix()
	}
	return f.GetGroveCoordinates()
}

func main() {
	println("Solution 1 - Sample:", solve1(sample))
	println("Solution 1 - Input:", solve1(input))
	println("Solution 2 - Sample:", solve2(sample))
	println("Solution 2 - Input:", solve2(input))
}
