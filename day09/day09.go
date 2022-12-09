package day09

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"aoc2022/aoc"
)

type Rope []*Knot
type Knot struct {
	Row int
	Col int
}

// Move moves the knot by one space in the provided direction.
func (k *Knot) Move(direction string) {
	switch direction {
	case "R":
		k.Col += 1
	case "L":
		k.Col -= 1
	case "U":
		k.Row -= 1
	case "D":
		k.Row += 1
	}
}

// Follow aligns k with o based on how far apart they are and returns true if the knot moved.
func (k *Knot) Follow(o *Knot) bool {
	switch {
	case o.Row > k.Row+1: // need to move down
		k.Col = o.Col
		k.Row += 1
		return true
	case o.Row < k.Row-1: // need to move up
		k.Col = o.Col
		k.Row -= 1
		return true
	case o.Col > k.Col+1: // need to move right
		k.Col += 1
		k.Row = o.Row
		return true
	case o.Col < k.Col-1: // need to move left
		k.Col -= 1
		k.Row = o.Row
		return true
	}

	return false
}

func Part1(r io.Reader) int {
	var rope = Rope{{}, {}}
	return len(moveRope(aoc.Lines(r), rope))
}

func moveRope(lines []string, rope []*Knot) map[string]bool {
	visited := make(map[string]bool)
	for _, l := range lines {
		s := strings.Fields(l)
		direction, length := s[0], aoc.Atoi(s[1])
		for i := 0; i < length; i++ {
			fmt.Println(direction)
			rope[0].Move(direction)
			for j := 1; j < len(rope); j++ {
				rope[j].Follow(rope[j-1])
			}
			h := strconv.Itoa(rope[0].Row) + "," + strconv.Itoa(rope[0].Col)
			t := strconv.Itoa(rope[len(rope)-1].Row) + "," + strconv.Itoa(rope[len(rope)-1].Col)
			fmt.Printf("H=(%s) T=(%s)\n", h, t)
			visited[t] = true
		}
	}
	return visited
}

func Part2(r io.Reader) int {
	var rope Rope
	for i := 0; i < 10; i++ {
		rope = append(rope, &Knot{})
	}
	return len(moveRope(aoc.Lines(r), rope))
}
