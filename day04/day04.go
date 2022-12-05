package day04

import (
	"io"
	"strconv"
	"strings"

	"golang.org/x/tools/container/intsets"

	"aoc2022/aoc"
)

func Part1(r io.Reader) int {
	return sum(
		r,
		func(a, b *intsets.Sparse) int {
			var i intsets.Sparse
			i.Intersection(a, b)
			if i.Equals(a) || i.Equals(b) {
				return 1
			}
			return 0
		},
	)
}

func Part2(r io.Reader) int {
	return sum(
		r,
		func(a, b *intsets.Sparse) int {
			a.IntersectionWith(b)
			if a.Len() > 0 {
				return 1
			}
			return 0
		},
	)
}

func sum(r io.Reader, predicate func(a, b *intsets.Sparse) int) int {
	count := 0
	lines := aoc.Lines(r)
	for _, l := range lines {
		count += predicate(sets(l))
	}

	return count
}

func atoi(s string) int {
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		panic(err)
	}

	return int(i)
}

func sparse(lower, upper string) *intsets.Sparse {
	var s intsets.Sparse
	for i := atoi(lower); i <= atoi(upper); i++ {
		s.Insert(i)
	}
	return &s
}

func sets(l string) (a, b *intsets.Sparse) {
	ranges := strings.FieldsFunc(l, func(r rune) bool {
		return r == '-' || r == ','
	})

	return sparse(ranges[0], ranges[1]), sparse(ranges[2], ranges[3])
}
