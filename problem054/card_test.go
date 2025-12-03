package main

import "testing"

func TestCardCompareRank(t *testing.T) {
	a, err := NewCard("AH")
	if err != nil {
		t.Fatalf("failed to parse AH: %v", err)
	}
	k, err := NewCard("KS")
	if err != nil {
		t.Fatalf("failed to parse KS: %v", err)
	}
	if !a.WinsAgainst(k) {
		t.Errorf("Ace should beat King")
	}

	ten1, err := NewCard("10D")
	if err != nil {
		t.Fatalf("failed to parse 10D: %v", err)
	}
	ten2, err := NewCard("TD")
	if err != nil {
		t.Fatalf("failed to parse TD: %v", err)
	}
	// equal ranks -> neither winsAgainst the other
	if ten1.WinsAgainst(ten2) || ten2.WinsAgainst(ten1) {
		t.Errorf("equal ranks should not WinAgainst each other")
	}
}
