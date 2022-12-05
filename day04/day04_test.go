package day04_test

import (
	"bytes"
	_ "embed"
	"io"
	"testing"

	"aoc2022/day04"
)

//go:embed input.txt
var input []byte

//go:embed example.txt
var example []byte

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
			want: 2,
		},
		{
			name: "input",
			args: args{r: bytes.NewReader(input)},
			want: 599,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day04.Part1(tt.args.r)
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
			want: 4,
		},
		{
			name: "input",
			args: args{r: bytes.NewReader(input)},
			want: 928,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day04.Part2(tt.args.r)
			if got != tt.want {
				t.Errorf("Part2() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkPart1(b *testing.B) {
	r := bytes.NewReader(input)
	for i := 0; i < b.N; i++ {
		day04.Part1(r)
	}
}

func BenchmarkPart2(b *testing.B) {
	r := bytes.NewReader(input)
	for i := 0; i < b.N; i++ {
		day04.Part2(r)
	}
}