package main

import (
	"aoc2022/pkg/rps"
)

func main() {
	var p2, p1 rps.Plays = rps.ReadPlays("data/sample-2.txt")
	println("Solution 1 - Sample:", rps.PlayGames(p1, p2))

	p2, p1 = rps.ReadPlays("data/input-2.txt")
	println("Solution 1 - Input:", rps.PlayGames(p1, p2))

	p2, p1 = rps.ReadPlays2("data/sample-2.txt")
	println("Solution 2 - Sample:", rps.PlayGames(p1, p2))

	p2, p1 = rps.ReadPlays2("data/input-2.txt")
	println("Solution 2 - Input:", rps.PlayGames(p1, p2))
}
