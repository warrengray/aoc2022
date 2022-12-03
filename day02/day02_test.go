package day02_test

import (
	"bytes"
	_ "embed"
	"io"
	"testing"

	"aoc2022/day02"
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
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "example",
			args: args{r: bytes.NewReader(example)},
			want: 15,
		},
		{
			name: "input",
			args: args{r: bytes.NewReader(input)},
			want: 15691,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := day02.Part1(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Part1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
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
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "example",
			args: args{r: bytes.NewReader(example)},
			want: 12,
		},
		{
			name: "input",
			args: args{r: bytes.NewReader(input)},
			want: 12989,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := day02.Part2(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Part2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Part2() got = %v, want %v", got, tt.want)
			}
		})
	}
}
