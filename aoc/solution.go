package aoc

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
)

type Solution func(io.Reader) (string, error)

func Run(input []byte, s ...Solution) {
	for i, s := range s {
		answer, err := s(bytes.NewReader(input))
		if err != nil {
			panic(err)
		}

		fmt.Println("Part " + strconv.Itoa(i+1) + ": " + answer)
	}
}

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
			panic("cannont convert " + l + "to int: " + err.Error())
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
