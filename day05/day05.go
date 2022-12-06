package day05

import (
	"io"
	"strings"

	"aoc2022/aoc"
)

type Stacks map[int][]rune

func Part1(r io.Reader) string {
	return operate(r, crateMover9000)
}

func Part2(r io.Reader) string {
	return operate(r, crateMover9001)
}

func operate(r io.Reader, crane func(string, Stacks)) string {
	stacks := make(Stacks)
	for _, l := range aoc.Lines(r) {
		switch {
		case l == "":
		case strings.Fields(l)[0] == "1":
		case strings.HasPrefix(l, "move"):
			crane(l, stacks)
		default:
			parseState(l, stacks)
		}
	}

	var order []rune
	for i := 1; i <= len(stacks); i++ {
		order = append(order, stacks[i][0])
	}

	return string(order)
}

func crateMover9001(instruction string, stacks Stacks) {
	f := strings.Fields(instruction)
	count, from, to := aoc.Atoi(f[1]), aoc.Atoi(f[3]), aoc.Atoi(f[5])
	var c []rune
	c = append(c, stacks[from][:count]...)
	stacks[to] = append(c, stacks[to]...)
	stacks[from] = stacks[from][count:]
}

func crateMover9000(instruction string, stacks Stacks) {
	f := strings.Fields(instruction)
	count, from, to := aoc.Atoi(f[1]), aoc.Atoi(f[3]), aoc.Atoi(f[5])
	for i := 0; i < count; i++ {
		var c rune
		stacks[from], c = stacks[from][1:], stacks[from][0]
		stacks[to] = append([]rune{c}, stacks[to]...)
	}
}

func parseState(state string, into Stacks) {
	s := []rune(state)
	for i := 0; i < len(s); i += 4 {
		if s[i+1] == ' ' {
			continue
		}

		into[(i/4)+1] = append(into[(i/4)+1], s[i+1])
	}
}
