package main

import "testing"

func Test_isLychrel(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		n    int
		want bool
	}{
		{"not1", 47, false},      // 47 becomes a palindrome in one iteration
		{"not2", 349, false},     // 349 becomes a palindrome in three iterations
		{"lychrel1", 196, true},  // 196 is believed to be a Lychrel number
		{"lychrel2", 4994, true}, // 4994 is believed to be a Lychrel number
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isLychrel(tt.n)
			if got != tt.want {
				t.Errorf("isLychrel() = %v, want %v", got, tt.want)
			}
		})
	}
}
