package main

import (
	"aoc2022/pkg/device"
	"aoc2022/pkg/io"
)

const input = "./data/input-6.txt"

func solve1(signal string) int {
	return device.FindStartOfPacket(signal, 4)
}

func solve2(signal string) int {
	return device.FindStartOfPacket(signal, 14)
}

func main() {
	println("Solution 1 - Sample 1:", solve1("mjqjpqmgbljsphdztnvjfqwrcgsmlb"))
	println("Solution 1 - Sample 2:", solve1("bvwbjplbgvbhsrlpgdmjqwftvncz"))
	println("Solution 1 - Sample 3:", solve1("nppdvjthqldpwncqszvftbrmjlhg"))
	println("Solution 1 - Sample 4:", solve1("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"))
	println("Solution 1 - Sample 5:", solve1("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"))
	println("Solution 1 - Input:", solve1(io.ReadLines(input)[0]))
	println("Solution 2 - Sample 1:", solve2("mjqjpqmgbljsphdztnvjfqwrcgsmlb"))
	println("Solution 2 - Sample 2:", solve2("bvwbjplbgvbhsrlpgdmjqwftvncz"))
	println("Solution 2 - Sample 3:", solve2("nppdvjthqldpwncqszvftbrmjlhg"))
	println("Solution 2 - Sample 4:", solve2("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"))
	println("Solution 2 - Sample 5:", solve2("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"))
	println("Solution 2 - Input:", solve2(io.ReadLines(input)[0]))
}
