package day01

import (
	"io"
	"sort"

	"aoc2022/aoc"
)

func Part1(r io.Reader) int {
	sums := sortedCaloriesPerElf(r)
	return sums[len(sums)-1]
}

func Part2(r io.Reader) int {
	sums := sortedCaloriesPerElf(r)
	return sums[len(sums)-1] + sums[len(sums)-2] + sums[len(sums)-3]
}

func sortedCaloriesPerElf(r io.Reader) []int {
	elves := aoc.NewLineGroupedInts(r)
	sums := make([]int, len(elves))
	for i, e := range elves {
		for _, c := range e {
			sums[i] += int(c)
		}
	}

	sort.Ints(sums)
	return sums
}
