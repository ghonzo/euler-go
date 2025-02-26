// Project Euler
package main

import "testing"

func Test_sumDivisors(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{220}, 284},
		{"example2", args{284}, 220},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumDivisors(tt.args.n); got != tt.want {
				t.Errorf("sumDivisors() = %v, want %v", got, tt.want)
			}
		})
	}
}
