// Project Euler
package main

import "testing"

func Test_calculatePhi(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		n    int
		want int
	}{
		{"example1", 9, 6},
		{"example2", 87109, 79180},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculatePhi(tt.n)
			if got != tt.want {
				t.Errorf("calculatePhi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_arePermutations(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		a    int
		b    int
		want bool
	}{
		{"example", 87109, 79180, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := arePermutations(tt.a, tt.b)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("arePermutations() = %v, want %v", got, tt.want)
			}
		})
	}
}
