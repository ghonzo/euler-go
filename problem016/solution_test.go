// Project Euler
package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		p int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{15}, 26},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.p); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
