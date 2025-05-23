// Project Euler
package main

import (
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{5}, 19},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.n); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_spellLetters(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{342}, 23},
		{"example2", args{115}, 20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spellLetters(tt.args.n); got != tt.want {
				t.Errorf("spellLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}
