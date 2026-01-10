package main

import "testing"


func Test_digitalSum(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		n    int
		want int
	}{
		{"example", 2, 475},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := digitalSum(tt.n)
			if got != tt.want {
				t.Errorf("digitalSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
