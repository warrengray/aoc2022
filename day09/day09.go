package day09

import (
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"

	"aoc2022/aoc"
)

func Part1(r io.Reader) int {
	visited := make(map[string]bool)
	var trow, tcol, hrow, hcol int
	for _, l := range aoc.Lines(r) {
		_, _ = hrow, hcol
		s := strings.Fields(l)
		direction, length := s[0], aoc.Atoi(s[1])
		var drow, dcol int
		switch direction {
		case "R":
			dcol = 1
		case "L":
			dcol = -1
		case "U":
			drow = -1
		case "D":
			drow = 1
		}
		for i := 0; i < length; i++ {
			hrow += drow
			hcol += dcol
			switch direction {
			case "R":
				if farEnoughAway(hrow, hcol, trow, tcol) {
					tcol = hcol - 1
					trow = hrow
				}
			case "L":
				if farEnoughAway(hrow, hcol, trow, tcol) {
					tcol = hcol + 1
					trow = hrow
				}
			case "U":
				if farEnoughAway(hrow, hcol, trow, tcol) {
					tcol = hcol
					trow = hrow + 1
				}
			case "D":
				if farEnoughAway(hrow, hcol, trow, tcol) {
					tcol = hcol
					trow = hrow - 1
				}
			}
			head := strconv.Itoa(hrow) + "," + strconv.Itoa(hcol)
			tail := strconv.Itoa(trow) + "," + strconv.Itoa(tcol)
			fmt.Printf("H=(%s) T=(%s)\n", head, tail)
			visited[tail] = true
		}
	}
	return len(visited)
}

func farEnoughAway(hrow, hcol, trow, tcol int) bool {
	return math.Abs(float64(hrow-trow)) >= 2.0 || math.Abs(float64(hcol-tcol)) >= 2.0
}

func Part2(r io.Reader) int {
	return 0
}
