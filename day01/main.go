package main

import (
	_ "embed"

	"aoc2022/aoc"
)

//go:embed input.txt
var input []byte

//go:embed example.txt
var example []byte

func main() {
	aoc.Run(example, Part1, Part2)
	aoc.Run(input, Part1, Part2)
}
