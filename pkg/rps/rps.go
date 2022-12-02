package rps

import "aoc2022/pkg/io"

type Play int
type Plays []Play

const (
	Rock     Play = 1
	Paper         = 2
	Scissors      = 3
)

type Score int

const (
	Lose = 0
	Draw = 3
	Win  = 6
)

func ReadPlays(file string) (Plays, Plays) {
	var p1, p2 Plays
	for _, line := range io.ReadLines(file) {
		switch line[0] {
		case 'A':
			p1 = append(p1, Rock)
		case 'B':
			p1 = append(p1, Paper)
		case 'C':
			p1 = append(p1, Scissors)
		}
		switch line[2] {
		case 'X':
			p2 = append(p2, Rock)
		case 'Y':
			p2 = append(p2, Paper)
		case 'Z':
			p2 = append(p2, Scissors)
		}
	}
	return p1, p2
}

func ReadPlays2(file string) (Plays, Plays) {
	var p1, p2 Plays
	for _, line := range io.ReadLines(file) {
		var lp1 Play
		switch line[0] {
		case 'A':
			lp1 = Rock
		case 'B':
			lp1 = Paper
		case 'C':
			lp1 = Scissors
		}

		p1 = append(p1, lp1)

		switch line[2] {
		case 'X': // p2 loses
			switch lp1 {
			case Rock:
				p2 = append(p2, Scissors)
			case Paper:
				p2 = append(p2, Rock)
			case Scissors:
				p2 = append(p2, Paper)
			}
		case 'Y': // draw
			p2 = append(p2, lp1)
		case 'Z': // p2 win
			switch lp1 {
			case Rock:
				p2 = append(p2, Paper)
			case Paper:
				p2 = append(p2, Scissors)
			case Scissors:
				p2 = append(p2, Rock)
			}
		}
	}
	return p1, p2
}

func PlayGames(p1, p2 Plays) int {
	var score int = 0
	for i, play1 := range p1 {
		score += int(PlayGame(play1, p2[i])) + int(play1)
	}
	return score
}

func PlayGame(p1, p2 Play) Score {
	if p1 == p2 {
		return Draw
	}
	if p1 == Rock && p2 == Paper || p1 == Paper && p2 == Scissors || p1 == Scissors && p2 == Rock {
		return Lose
	}
	return Win
}
