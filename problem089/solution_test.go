package main

import "testing"

func Test_parseRoman(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		want int
	}{
		{"example1", "XIV", 14},
		{"example2", "MCMXCIV", 1994},
		{"example3", "LVIII", 58},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseRoman(tt.s)
			if got != tt.want {
				t.Errorf("parseRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
