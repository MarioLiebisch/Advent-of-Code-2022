package main

import (
	"aoc2022/pkg/io"
	"aoc2022/pkg/numbers"
)

func main() {
	println("Solution 1 - Sample:", numbers.SliceMax(numbers.SummarizeGroups(io.ReadIntegers("data/sample-1.txt", true))))
	println("Solution 1 - Input:", numbers.SliceMax(numbers.SummarizeGroups(io.ReadIntegers("data/input-1.txt", true))))

	println("Solution 2 - Sample:", numbers.SummarizeGroups(numbers.SliceMaxN(numbers.SummarizeGroups(io.ReadIntegers("data/sample-1.txt", true)), 3))[0])
	println("Solution 2 - Input:", numbers.SummarizeGroups(numbers.SliceMaxN(numbers.SummarizeGroups(io.ReadIntegers("data/input-1.txt", true)), 3))[0])
}
