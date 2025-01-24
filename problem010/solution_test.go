// Project Euler
package main

import (
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		nth int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{10}, 17},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.nth); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solve2(t *testing.T) {
	type args struct {
		maxLimit int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{10}, 17},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve2(tt.args.maxLimit); got != tt.want {
				t.Errorf("solve2() = %v, want %v", got, tt.want)
			}
		})
	}
}
