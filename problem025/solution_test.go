// Project Euler
package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		digits int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{3}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.digits); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
