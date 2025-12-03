package main

import "fmt"

// Card representations for poker hands used in Problem 54.
// Ranks are ordered: 2..10, Jack, Queen, King, Ace.
type Rank int

const (
    Two Rank = 2
    Three Rank = 3
    Four Rank = 4
    Five Rank = 5
    Six Rank = 6
    Seven Rank = 7
    Eight Rank = 8
    Nine Rank = 9
    Ten Rank = 10
    Jack Rank = 11
    Queen Rank = 12
    King Rank = 13
    Ace Rank = 14
)

type Suit rune

type Card struct {
    Rank Rank
    Suit Suit
}

// NewCard parses a card code into a Card. Supported codes include:
//  - 2-character codes like "AH", "9C", "TD" (T == 10)
//  - 3-character codes with explicit 10: "10H"
// Returns an error for invalid inputs.
func NewCard(code string) (Card, error) {
    if len(code) == 0 {
        return Card{}, fmt.Errorf("empty card code")
    }
    var rch, sch string
    if len(code) == 3 { // e.g. "10H"
        rch = code[:2]
        sch = code[2:]
    } else if len(code) == 2 {
        rch = code[:1]
        sch = code[1:]
    } else {
        return Card{}, fmt.Errorf("invalid card code: %s", code)
    }
    var rank Rank
    switch rch {
    case "2":
        rank = Two
    case "3":
        rank = Three
    case "4":
        rank = Four
    case "5":
        rank = Five
    case "6":
        rank = Six
    case "7":
        rank = Seven
    case "8":
        rank = Eight
    case "9":
        rank = Nine
    case "10", "T", "t":
        rank = Ten
    case "J", "j":
        rank = Jack
    case "Q", "q":
        rank = Queen
    case "K", "k":
        rank = King
    case "A", "a":
        rank = Ace
    default:
        return Card{}, fmt.Errorf("invalid rank: %s", rch)
    }
    if len(sch) != 1 {
        return Card{}, fmt.Errorf("invalid suit: %s", sch)
    }
    suit := Suit(sch[0])
    return Card{Rank: rank, Suit: suit}, nil
}

// WinsAgainst returns true if card c has a higher rank than other.
// Suit is ignored for rank comparison.
func (c Card) WinsAgainst(other Card) bool {
    return c.Rank > other.Rank
}
