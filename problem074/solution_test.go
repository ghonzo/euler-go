// Project Euler
package main

import "testing"

func Test_chainLength(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		n    int
		want int
	}{
		{"145", 145, 1},
		{"169", 169, 3},
		{"871", 871, 2},
		{"872", 872, 2},
		{"69", 69, 5},
		{"78", 78, 4},
		{"540", 540, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := chainLength(tt.n)
			if got != tt.want {
				t.Errorf("chainLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
