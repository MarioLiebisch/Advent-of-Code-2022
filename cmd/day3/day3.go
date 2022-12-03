package main

import (
	"aoc2022/pkg/io"
	"aoc2022/pkg/strings"
)

const priorities = " abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func solve1(file string) int {
	var sum = 0
	for _, line := range io.ReadLines(file) {
		var c1 = line[:len(line)/2]
		var c2 = line[len(line)/2:]
		var cu = strings.GetCommonChars(c1, c2)
		for _, c := range cu {
			for p, v := range priorities {
				if c == v {
					sum += p
					break
				}
			}
		}
	}
	return sum
}

func solve2(file string) int {
	var sum = 0
	var lines = io.ReadLines(file)
	var group_count = len(lines) / 3
	for i := 0; i < group_count; i++ {
		var c1 = lines[i*3]
		var c2 = lines[i*3+1]
		var c3 = lines[i*3+2]
		var cu = strings.GetCommonChars(strings.GetCommonChars(c1, c2), c3)
		for _, c := range cu {
			for p, v := range priorities {
				if c == v {
					sum += p
					break
				}
			}
		}
	}
	return sum
}

func main() {
	println("Solution 1 - Sample:", solve1("./data/sample-3.txt"))
	println("Solution 1 - Input:", solve1("./data/input-3.txt"))
	println("Solution 2 - Sample:", solve2("./data/sample-3.txt"))
	println("Solution 2 - Input:", solve2("./data/input-3.txt"))
}
