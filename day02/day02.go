package day02

import (
	"io"
	"strings"

	"aoc2022/aoc"
)

var (
	// A, X = rock
	// B, Y = paper
	// C, Z = scissors
	plays       = map[string]int{"X": 1, "Y": 2, "Z": 3}
	moves       = map[string]string{"A": "X", "B": "Y", "C": "Z"}
	winsAgainst = map[string]string{"X": "Z", "Y": "X", "Z": "Y"}
	losesTo     = make(map[string]string)
)

func init() {
	for k, v := range winsAgainst {
		losesTo[v] = k
	}
}

func Part1(r io.Reader) int {
	return tally(
		r,
		func(us, _ string) string {
			return us
		},
	)
}

func Part2(r io.Reader) int {
	return tally(
		r,
		func(us, them string) string {
			switch us {
			case "X": // need to lose
				return winsAgainst[them]
			case "Y": // need to draw
				return them
			case "Z": // need to win
				return losesTo[them]
			}
			return us
		},
	)
}

func tally(r io.Reader, decider func(us, them string) string) int {
	var score int
	for _, l := range aoc.Lines(r) {
		round := strings.Fields(l)
		them := moves[round[0]]
		us := decider(round[1], them)
		score += plays[us]
		switch {
		case us == them:
			score += 3 // draw
		case winsAgainst[us] == them:
			score += 6 // win
		}
	}

	return score
}
