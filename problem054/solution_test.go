package main

import "testing"

func Test_player1Wins(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		entry string
		want  bool
	}{
		{"hand1", "5H 5C 6S 7S KD 2C 3S 8S 8D TD", false}, // Player 2 wins with a pair of 8s over player 1's pair of 5s
		{"hand2", "5D 8C 9S JS AC 2C 5C 7D 8S QH", true},  // Player 1 wins with high card Ace over player 2's high card Queen
		{"hand3", "2D 9C AS AH AC 3D 6D 7D TD QD", false}, // Player 2 wins with flush over player 1's three Aces
		{"hand4", "4D 6S 9H QH QC 3D 6D 7H QD QS", true},  // Player 1 wins with pair of Queens and 9 kicker over player 2's pair of Queens and 7 kicker
		{"hand5", "2H 2D 4C 4D 4S 3C 3D 3S 9S 9D", true},  // Player 1 wins with full house (4s over 2s) over player 2's full house (3s over 9s)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := player1Wins(tt.entry)
			if got != tt.want {
				t.Errorf("player1Wins() = %v, want %v", got, tt.want)
			}
		})
	}
}
