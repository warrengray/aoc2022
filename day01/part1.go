package main

import (
	"io"
	"sort"
	"strconv"

	"aoc2022/aoc"
)

func Part1(r io.Reader) (string, error) {
	sums := sortedCaloriesPerElf(r)
	return strconv.Itoa(sums[len(sums)-1]), nil
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
