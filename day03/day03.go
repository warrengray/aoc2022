package day03

import (
	"io"
	"unicode"

	"aoc2022/aoc"
)

func Part1(r io.Reader) int {
	sum := 0
	for _, s := range aoc.Lines(r) {
		d := findDuplicate([]rune(s))
		if unicode.IsLower(d) {
			sum += int(d - 96)
			continue
		}

		sum += int(d - 38)
	}
	return sum
}

func findDuplicate(s []rune) rune {
	items := make(map[rune]bool)
	var i int
	for _, c := range s {
		i++
		if i <= len(s)/2 {
			items[c] = true
			continue
		}

		if items[c] {
			return c
		}
	}
	return 0
}

func Part2(r io.Reader) int {
	return 0
}
