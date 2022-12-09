package day09_test

import (
	"bytes"
	_ "embed"
	"io"
	"testing"

	"aoc2022/day09"
)

//go:embed input.txt
var input []byte

//go:embed example.txt
var example []byte

//go:embed example_large.txt
var exampleLarge []byte

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
			want: 1,
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

// between 2414 and 2456

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
