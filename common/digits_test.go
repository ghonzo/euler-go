package common_test

import (
	"slices"
	"testing"

	"github.com/ghonzo/euler-go/common"
)

func TestDigitsFromInt(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		n    int
		want common.Digits
	}{
		{"zero", 0, common.Digits{0}},
		{"single digit", 7, common.Digits{7}},
		{"multiple digits", 12345, common.Digits{1, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := common.DigitsFromInt(tt.n)
			if !slices.Equal(got, tt.want) {
				t.Errorf("DigitsFromInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDigits_Int(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		n    common.Digits
		want int
	}{
		{"zero", common.Digits{0}, 0},
		{"single digit", common.Digits{7}, 7},
		{"multiple digits", common.Digits{1, 2, 3, 4, 5}, 12345},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.n.Int()
			if got != tt.want {
				t.Errorf("Int() = %v, want %v", got, tt.want)
			}
		})
	}
}
