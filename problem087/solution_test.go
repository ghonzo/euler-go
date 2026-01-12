package main

import "testing"



func Test_solve(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		upperLimit int
		want       int
	}{
		{"example", 50, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := solve(tt.upperLimit)			
			if got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
