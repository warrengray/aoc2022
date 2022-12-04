package day03

import (
	"io"
	"unicode"

	"golang.org/x/tools/container/intsets"

	"aoc2022/aoc"
)

func Part1(r io.Reader) int {
	sum := 0
	for _, s := range aoc.Lines(r) {
		runes := []rune(s)
		sum += priority(findCommon(runes[:len(runes)/2], runes[len(runes)/2:]))
	}

	return sum
}

func Part2(r io.Reader) int {
	sum := 0
	lines := aoc.Lines(r)
	for i := 0; i < len(lines); i += 3 {
		sum += priority(findCommon([]rune(lines[i]), []rune(lines[i+1]), []rune(lines[i+2])))
	}

	return sum
}

func priority(d rune) int {
	if unicode.IsLower(d) {
		return int(d - 96)
	}

	return int(d - 38)
}

func findCommon(sets ...[]rune) rune {
	s := sparse(sets[0])
	for i := 1; i < len(sets); i++ {
		s.IntersectionWith(sparse(sets[i]))
	}

	if s.Len() > 1 {
		panic("more than one common result")
	}

	return rune(s.Min())
}

func sparse(rs []rune) *intsets.Sparse {
	var s intsets.Sparse
	for _, r := range rs {
		s.Insert(int(r))
	}
	return &s
}
