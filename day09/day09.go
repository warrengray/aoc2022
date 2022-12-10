package day09

import (
	"io"
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
		k.Row += 1
	case "D":
		k.Row -= 1
	}
}

// Follow aligns k with o based on how far apart they are and returns true if the knot moved.
func (k *Knot) Follow(o *Knot) bool {
	// if we're adjacent in both the horizontal and vertical, there's nothing to do
	if aoc.IntAbs(k.Row-o.Row) < 2 && aoc.IntAbs(k.Col-o.Col) < 2 {
		return false
	}

	if k.Col < o.Col {
		k.Col += 1
	}

	if k.Row < o.Row {
		k.Row += 1
	}

	if k.Row > o.Row {
		k.Row -= 1
	}

	if k.Col > o.Col {
		k.Col -= 1
	}

	return true
}

func Part1(r io.Reader) int {
	return len(moveRope(2, aoc.Lines(r)))
}

func Part2(r io.Reader) int {
	return len(moveRope(10, aoc.Lines(r)))
}

func moveRope(length int, moves []string) map[Knot]bool {
	var rope Rope
	for i := 0; i < length; i++ {
		rope = append(rope, &Knot{})
	}

	visited := make(map[Knot]bool)
	for _, l := range moves {
		s := strings.Fields(l)
		dir, dist := s[0], aoc.Atoi(s[1])
		for i := 0; i < dist; i++ {
			rope[0].Move(dir)
			visited[align(rope[0], rope[1:])] = true
		}
	}
	return visited
}

// align makes Rope r align Knot k, returning the Knot at the end of the rope (the tail).
func align(k *Knot, r Rope) Knot {
	if len(r) == 0 {
		return *k
	}

	if !r[0].Follow(k) {
		return *r[len(r)-1]
	}

	return align(r[0], r[1:])
}
