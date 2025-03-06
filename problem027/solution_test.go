// Project Euler
package main

import "testing"

func Test_findNumPrimes(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{1, 41}, 40},
		{"example2", args{-79, 1601}, 80},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNumPrimes(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("findNumPrimes() = %v, want %v", got, tt.want)
			}
		})
	}
}
