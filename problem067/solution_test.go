// Project Euler
package main

import (
	"testing"

	"github.com/ghonzo/euler-go/common"
)

func Test_solve(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		entries []string
		want    int
	}{
		{"example", common.ReadStringsFromFile("example.txt"), 23},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := solve(tt.entries)
			if got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
