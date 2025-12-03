package main

import (
	"cmp"
	"fmt"
)

// Card representations for poker hands used in Problem 54.
// Ranks are ordered: 2..10, Jack, Queen, King, Ace.
type Rank int

const (
	Two   Rank = 2
	Three Rank = 3
	Four  Rank = 4
	Five  Rank = 5
	Six   Rank = 6
	Seven Rank = 7
	Eight Rank = 8
	Nine  Rank = 9
	Ten   Rank = 10
	Jack  Rank = 11
	Queen Rank = 12
	King  Rank = 13
	Ace   Rank = 14
)

type Suit byte

type Card struct {
	Rank Rank
	Suit Suit
}

type Hand []Card

// NewCard parses a card code into a Card. Supported code is:
//   - 2-character codes like "AH", "9C", "TD" (T == 10)
func NewCard(code string) (Card, error) {
	if len(code) == 0 {
		return Card{}, fmt.Errorf("empty card code")
	}
	var rch, sch byte
	if len(code) == 2 {
		rch = code[0]
		sch = code[1]
	} else {
		return Card{}, fmt.Errorf("invalid card code: %s", code)
	}
	var rank Rank
	switch rch {
	case '2':
		rank = Two
	case '3':
		rank = Three
	case '4':
		rank = Four
	case '5':
		rank = Five
	case '6':
		rank = Six
	case '7':
		rank = Seven
	case '8':
		rank = Eight
	case '9':
		rank = Nine
	case 'T':
		rank = Ten
	case 'J':
		rank = Jack
	case 'Q':
		rank = Queen
	case 'K':
		rank = King
	case 'A':
		rank = Ace
	default:
		return Card{}, fmt.Errorf("invalid rank: %c", rch)
	}
	suit := Suit(sch)
	return Card{Rank: rank, Suit: suit}, nil
}

func Compare(x, y Card) int {
	return cmp.Compare(x.Rank, y.Rank)
}

// Compare two poker hands h1 and h2.
// Returns 1 if h1 wins, -1 if h2 wins, 0 if tie.
func ComparePokerHands(h1, h2 Hand) int {
	// Implementation goes here
	return 0
}
