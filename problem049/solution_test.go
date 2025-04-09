// Project Euler
package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		lower int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"first", args{1001}, "148748178147"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.lower); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
