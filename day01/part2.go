package main

import (
	"io"
	"strconv"
)

func Part2(r io.Reader) (string, error) {
	sums := sortedCaloriesPerElf(r)
	return strconv.Itoa(sums[len(sums)-1] + sums[len(sums)-2] + sums[len(sums)-3]), nil
}
