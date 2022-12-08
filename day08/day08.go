package day08

import (
	"fmt"
	"io"
	"os"

	"aoc2022/aoc"
)

type (
	Forest    [][]Tree
	Direction func(row, col int) (int, int)
	Tree      struct {
		Visible bool
		Height  int
	}
)

func Part1(r io.Reader) int {
	var forest Forest
	for _, l := range aoc.Lines(r) {
		trees := []rune(l)
		row := make([]Tree, len(trees))
		forest = append(forest, row)
		for i, t := range trees {
			row[i].Height = aoc.Atoi(string(t))
		}
	}

	// all trees on the edges are visible
	var visible int
	for row := range forest {
		for col := range forest[row] {
			if isVisibleAtAll(forest, row, col) {
				forest[row][col].Visible = true
				visible += 1
			}
		}
	}

	printForest(forest)
	return visible
}

func up(row, col int) (int, int) {
	return row - 1, col
}

func down(row, col int) (int, int) {
	return row + 1, col
}

func left(row, col int) (int, int) {
	return row, col - 1
}

func right(row, col int) (int, int) {
	return row, col + 1
}

func printForest(f Forest) {
	for row := range f {
		_, _ = fmt.Fprintf(os.Stdout, "\n")
		for col := range f[row] {
			val := "."
			if f[row][col].Visible {
				val = "v"
			}
			_, _ = fmt.Fprint(os.Stdout, val)
		}
	}
}

func isVisibleAtAll(f Forest, row, col int) bool {
	if isAtEdge(f, row, col) {
		return true
	}

	for _, d := range []Direction{up, down, left, right} {
		if isVisible(f, row, col, d) {
			return true
		}
	}

	return false
}

func isAtEdge(f Forest, row int, col int) bool {
	return row == 0 || col == 0 || row == len(f)-1 || col == len(f[row])-1
}

func isVisible(f Forest, row, col int, d Direction) bool {
	t := f[row][col]
	for {
		row, col = d(row, col)
		nextTree := f[row][col]
		if nextTree.Height >= t.Height {
			return false
		}

		if isAtEdge(f, row, col) {
			return true
		}
	}
}

func Part2(r io.Reader) int {
	return 0
}
