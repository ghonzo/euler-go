package main

import "testing"

func Test_solveMatrix(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		matrix [][]int
		want   int
	}{
		{"example", [][]int{
			{131, 673, 234, 103, 18},
			{201, 96, 342, 965, 150},
			{630, 803, 746, 422, 111},
			{537, 699, 497, 121, 956},
			{805, 732, 524, 37, 331},
		}, 994},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := solveMatrix(tt.matrix)
			if got != tt.want {
				t.Errorf("solveMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
