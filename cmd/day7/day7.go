package main

import (
	"aoc2022/pkg/device"
	"fmt"
)

const sample = "./data/sample-7.txt"
const input = "./data/input-7.txt"

func solve1(file string) int {
	fs := device.ReadFileSystem(file)
	return fs.GetSizeSum(100000)
}

func solve2(file string) string {
	fs := device.ReadFileSystem(file)
	c := fs.GetDeletionCandidate(30000000 - (70000000 - fs.GetSize()))
	return c.GetPath() + ": " + fmt.Sprint(c.GetSize())
}
func main() {
	println("Solution 1 - Sample:", solve1(sample))
	println("Solution 1 - Input:", solve1(input))
	println("Solution 2 - Sample:", solve2(sample))
	println("Solution 2 - Input:", solve2(input))
}
