package day01_test

import (
	"bytes"
	_ "embed"
	"io"
	"testing"

	"aoc2022/day01"
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
			want: 24000,
		},
		{
			name: "input",
			args: args{r: bytes.NewReader(input)},
			want: 66719,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day01.Part1(tt.args.r)
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
			want: 45000,
		},
		{
			name: "input",
			args: args{r: bytes.NewReader(input)},
			want: 198551,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day01.Part2(tt.args.r)
			if got != tt.want {
				t.Errorf("Part2() got = %v, want %v", got, tt.want)
			}
		})
	}
}
