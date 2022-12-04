package main

import (
	"aoc2022/pkg/io"
	"aoc2022/pkg/numbers"
	"aoc2022/pkg/strings"
	bstrings "strings"
)

func solve1(file string) int {
	var res = 0
	for _, line := range io.ReadLines(file) {
		var assignments = bstrings.Split(line, ",")
		var from1, to1 = strings.ParseRange(assignments[0])
		var from2, to2 = strings.ParseRange(assignments[1])
		if (from2 >= from1 && to2 <= to1) || (from1 >= from2 && to1 <= to2) {
			res++
		}
	}
	return res
}

func solve2(file string) int {
	var res = 0
	for _, line := range io.ReadLines(file) {
		var assignments = bstrings.Split(line, ",")
		var from1, to1 = strings.ParseRange(assignments[0])
		var from2, to2 = strings.ParseRange(assignments[1])
		if numbers.BetweenInclusive(from1, from2, to2) ||
			numbers.BetweenInclusive(to1, from2, to2) ||
			numbers.BetweenInclusive(from2, from1, to1) ||
			numbers.BetweenInclusive(to2, from1, to1) {
			res++
		}
	}
	return res
}

func main() {
	println("Solution 1 - Sample:", solve1("./data/sample-4.txt"))
	println("Solution 1 - Input:", solve1("./data/input-4.txt"))
	println("Solution 2 - Sample:", solve2("./data/sample-4.txt"))
	println("Solution 2 - Input:", solve2("./data/input-4.txt"))
}
