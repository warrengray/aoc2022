package day06

import (
	"io"

	"aoc2022/aoc"
)

func Part1(r io.Reader) int {
	return processInput(r, 4)
}

func Part2(r io.Reader) int {
	return processInput(r, 14)
}

func processInput(r io.Reader, uniqueLength int) int {
	input := []rune(aoc.Lines(r)[0])
	for i := range input {
		if !containsDuplicates(input[i : i+uniqueLength]) {
			return i + uniqueLength
		}
	}

	return -1
}

func containsDuplicates(rs []rune) bool {
	m := make(map[rune]bool)
	for _, c := range rs {
		m[c] = true
	}

	return len(m) < len(rs)
}
