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
