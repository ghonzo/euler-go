package main

import "testing"


func Test_solve(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		sides int
		want  string
	}{
		{"example", 6, "102400"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := solve(tt.sides)
			if got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
