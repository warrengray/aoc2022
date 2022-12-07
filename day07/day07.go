package day07

import (
	"io"
	"math"
	"strings"

	"aoc2022/aoc"
)

func Part1(r io.Reader) int {
	var sum int
	for _, size := range sizes(r) {
		if size <= 100000 {
			sum += size
		}
	}

	return sum
}

func Part2(r io.Reader) int {
	dirSizes := sizes(r)
	capacity := 70000000
	required := 30000000 - (capacity - dirSizes[""])
	smallest := capacity

	for _, size := range dirSizes {
		if size < required {
			continue
		}

		smallest = int(math.Min(float64(size), float64(smallest)))
	}

	return smallest
}

func sizes(r io.Reader) map[string]int {
	dirSizes := make(map[string]int)
	stack := []string{""}
	for _, l := range aoc.Lines(r) {
		f := strings.Fields(l)
		switch {
		case f[0] == "dir":
		case f[1] == "ls":
		case f[1] == "cd" && f[2] == "/": // go to root
			stack = []string{""}
		case f[1] == "cd" && f[2] == "..": // come out of dir
			stack = stack[:len(stack)-1]
		case f[1] == "cd": // go into dir
			stack = append(stack, f[2])
		default: // file
			size := aoc.Atoi(f[0])
			for i := range stack {
				dirSizes[strings.Join(stack[:i+1], "/")] += size
			}
		}
	}

	return dirSizes
}
