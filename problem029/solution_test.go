// Project Euler
package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		top int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{5}, 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.top); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
