package aoc

import (
	"bufio"
	"io"
	"strconv"
)

func NewLineGroupedInts(r io.Reader) [][]int64 {
	allInts := make([][]int64, 0, 0)
	allInts = append(allInts, make([]int64, 0, 0))
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		l := scanner.Text()
		if l == "" {
			allInts = append(allInts, make([]int64, 0, 0))
			continue
		}

		i, err := strconv.ParseInt(l, 10, 64)
		if err != nil {
			panic("cannot convert " + l + "to int: " + err.Error())
		}

		allInts[len(allInts)-1] = append(allInts[len(allInts)-1], i)
	}

	return allInts
}

func Lines(r io.Reader) []string {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

// Atoi exists for brevity. It panics if the string isn't an int.
func Atoi(s string) int {
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		panic(err)
	}

	return int(i)
}
