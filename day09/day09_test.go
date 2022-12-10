package day09_test

import (
	"bytes"
	_ "embed"
	"io"
	"reflect"
	"testing"

	"aoc2022/day09"
)

//go:embed input.txt
var input []byte

//go:embed example.txt
var example []byte

//go:embed example_large.txt
var exampleLarge []byte

func TestFollow(t *testing.T) {
	type args struct {
		head day09.Knot
		tail day09.Knot
	}

	//    -1
	// -1  0 +1
	//    +1
	tests := []struct {
		name string
		args args
		want args
	}{
		// simple
		{
			name: "tail-above",
			args: args{
				head: day09.Knot{},
				tail: day09.Knot{Row: -2},
			},
			want: args{
				head: day09.Knot{},
				tail: day09.Knot{Row: -1},
			},
		},
		{
			name: "tail-below",
			args: args{
				head: day09.Knot{},
				tail: day09.Knot{Row: 2},
			},
			want: args{
				head: day09.Knot{},
				tail: day09.Knot{Row: 1},
			},
		},
		{
			name: "tail-left-of",
			args: args{
				head: day09.Knot{},
				tail: day09.Knot{Col: -2},
			},
			want: args{
				head: day09.Knot{},
				tail: day09.Knot{Col: -1},
			},
		},
		{
			name: "tail-right-of",
			args: args{
				head: day09.Knot{},
				tail: day09.Knot{Col: 2},
			},
			want: args{
				head: day09.Knot{},
				tail: day09.Knot{Col: 1},
			},
		},
		// diagonal
		{
			name: "tail-up-left-of",
			args: args{
				head: day09.Knot{},
				tail: day09.Knot{Col: -1, Row: -2},
			},
			want: args{
				head: day09.Knot{},
				tail: day09.Knot{Col: 0, Row: -1},
			},
		},
		{
			name: "tail-left-up-of",
			args: args{
				head: day09.Knot{},
				tail: day09.Knot{Col: -2, Row: -1},
			},
			want: args{
				head: day09.Knot{},
				tail: day09.Knot{Col: -1, Row: 0},
			},
		},
		{
			name: "tail-left-down-of",
			args: args{
				head: day09.Knot{},
				tail: day09.Knot{Col: -2, Row: 1},
			},
			want: args{
				head: day09.Knot{},
				tail: day09.Knot{Col: -1, Row: 0},
			},
		},
		{
			name: "tail-down-left-of",
			args: args{
				head: day09.Knot{},
				tail: day09.Knot{Col: -1, Row: 2},
			},
			want: args{
				head: day09.Knot{},
				tail: day09.Knot{Col: 0, Row: 1},
			},
		},
		{
			name: "tail-down-right-of",
			args: args{
				head: day09.Knot{},
				tail: day09.Knot{Col: 1, Row: 2},
			},
			want: args{
				head: day09.Knot{},
				tail: day09.Knot{Col: 0, Row: 1},
			},
		},
		{
			name: "tail-right-down-of",
			args: args{
				head: day09.Knot{},
				tail: day09.Knot{Col: 2, Row: 1},
			},
			want: args{
				head: day09.Knot{},
				tail: day09.Knot{Col: 1, Row: 0},
			},
		},
		{
			name: "tail-right-up-of",
			args: args{
				head: day09.Knot{},
				tail: day09.Knot{Col: 2, Row: -1},
			},
			want: args{
				head: day09.Knot{},
				tail: day09.Knot{Col: 1, Row: 0},
			},
		},
		{
			name: "tail-up-right-of",
			args: args{
				head: day09.Knot{},
				tail: day09.Knot{Col: 1, Row: -2},
			},
			want: args{
				head: day09.Knot{},
				tail: day09.Knot{Col: 0, Row: -1},
			},
		},
		// adjacent
		{
			name: "tail-up-right-adjacent",
			args: args{
				head: day09.Knot{},
				tail: day09.Knot{Col: 1, Row: -1},
			},
			want: args{
				head: day09.Knot{},
				tail: day09.Knot{Col: 1, Row: -1},
			},
		},
		{
			name: "tail-right-adjacent",
			args: args{
				head: day09.Knot{},
				tail: day09.Knot{Col: 1},
			},
			want: args{
				head: day09.Knot{},
				tail: day09.Knot{Col: 1},
			},
		},
		{
			name: "tail-down-right-adjacent",
			args: args{
				head: day09.Knot{},
				tail: day09.Knot{Col: 1, Row: 1},
			},
			want: args{
				head: day09.Knot{},
				tail: day09.Knot{Col: 1, Row: 1},
			},
		},
		{
			name: "tail-down-adjacent",
			args: args{
				head: day09.Knot{},
				tail: day09.Knot{Row: 1},
			},
			want: args{
				head: day09.Knot{},
				tail: day09.Knot{Row: 1},
			},
		},
		{
			name: "tail-down-left-adjacent",
			args: args{
				head: day09.Knot{},
				tail: day09.Knot{Row: 1, Col: -1},
			},
			want: args{
				head: day09.Knot{},
				tail: day09.Knot{Row: 1, Col: -1},
			},
		},
		{
			name: "tail-left-adjacent",
			args: args{
				head: day09.Knot{},
				tail: day09.Knot{Col: -1},
			},
			want: args{
				head: day09.Knot{},
				tail: day09.Knot{Col: -1},
			},
		},
		{
			name: "tail-up-left-adjacent",
			args: args{
				head: day09.Knot{},
				tail: day09.Knot{Col: -1, Row: -1},
			},
			want: args{
				head: day09.Knot{},
				tail: day09.Knot{Col: -1, Row: -1},
			},
		},
		{
			name: "tail-up-adjacent",
			args: args{
				head: day09.Knot{},
				tail: day09.Knot{Row: -1},
			},
			want: args{
				head: day09.Knot{},
				tail: day09.Knot{Row: -1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tail := &tt.args.tail
			tail.Follow(&tt.args.head)
			if !reflect.DeepEqual(tt.args, tt.want) {
				t.Errorf("got = %+v, want %+v", tt.args, tt.want)
			}
		})
	}
}

func TestPart1(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example",
			args: args{r: bytes.NewReader(example)},
			want: 13,
		},
		{
			name: "input",
			args: args{r: bytes.NewReader(input)},
			want: 6011,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day09.Part1(tt.args.r)
			if got != tt.want {
				t.Errorf("Part1() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example",
			args: args{r: bytes.NewReader(example)},
			want: 1,
		},
		{
			name: "example_large",
			args: args{r: bytes.NewReader(exampleLarge)},
			want: 36,
		},
		{
			name: "input",
			args: args{r: bytes.NewReader(input)},
			want: 2419,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day09.Part2(tt.args.r)
			if got != tt.want {
				t.Errorf("Part2() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkPart1(b *testing.B) {
	r := bytes.NewReader(input)
	for i := 0; i < b.N; i++ {
		r.Reset(input)
		day09.Part1(r)
	}
}

func BenchmarkPart2(b *testing.B) {
	r := bytes.NewReader(input)
	for i := 0; i < b.N; i++ {
		r.Reset(input)
		day09.Part2(r)
	}
}
