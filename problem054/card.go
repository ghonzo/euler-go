package main

import (
	"cmp"
	"fmt"
	"sort"
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
	c1, t1 := evaluateHand(h1)
	c2, t2 := evaluateHand(h2)

	if c := cmp.Compare(c1, c2); c != 0 {
		return c
	}

	// same category -> compare tiebreakers lexicographically
	n := len(t1)
	if len(t2) < n {
		n = len(t2)
	}
	for i := 0; i < n; i++ {
		if c := cmp.Compare(t1[i], t2[i]); c != 0 {
			return c
		}
	}
	return 0
}

// evaluateHand returns (category, tiebreakerSlice)
// Higher category value means stronger hand.
func evaluateHand(h Hand) (int, []int) {
	// normalize ranks into ints
	ranks := make([]int, len(h))
	for i, c := range h {
		ranks[i] = int(c.Rank)
	}
	sort.Slice(ranks, func(i, j int) bool { return ranks[i] > ranks[j] })

	// count frequencies
	freq := map[int]int{}
	for _, r := range ranks {
		freq[r]++
	}

	// check flush
	flush := true
	for i := 1; i < len(h); i++ {
		if h[i].Suit != h[0].Suit {
			flush = false
			break
		}
	}

	// check straight (handle wheel A-2-3-4-5)
	straight := false
	highStraight := 0
	uniqueRanks := make([]int, 0, len(freq))
	for r := range freq {
		uniqueRanks = append(uniqueRanks, r)
	}
	sort.Ints(uniqueRanks)
	if len(uniqueRanks) == 5 {
		// normal straight
		min := uniqueRanks[0]
		max := uniqueRanks[4]
		if max-min == 4 {
			straight = true
			highStraight = max
		} else {
			// check wheel: A(14),2,3,4,5 -> treat high as 5
			// uniqueRanks would be [2,3,4,5,14]
			if uniqueRanks[0] == 2 && uniqueRanks[1] == 3 && uniqueRanks[2] == 4 && uniqueRanks[3] == 5 && uniqueRanks[4] == 14 {
				straight = true
				highStraight = 5
			}
		}
	}

	// build list of (count, rank) and sort by count desc then rank desc
	type cr struct{ cnt, rank int }
	crs := make([]cr, 0, len(freq))
	for r, c := range freq {
		crs = append(crs, cr{c, r})
	}
	sort.Slice(crs, func(i, j int) bool {
		if crs[i].cnt != crs[j].cnt {
			return crs[i].cnt > crs[j].cnt
		}
		return crs[i].rank > crs[j].rank
	})

	// Categories (higher is better):
	// 8: Straight Flush, 7: Four of a Kind, 6: Full House,
	// 5: Flush, 4: Straight, 3: Three of a Kind, 2: Two Pair,
	// 1: One Pair, 0: High Card
	// tiebreaker slices are ordered by significance.
	// Straight flush
	if straight && flush {
		return 8, []int{highStraight}
	}

	// Four of a kind
	if crs[0].cnt == 4 {
		fourRank := crs[0].rank
		kicker := crs[1].rank
		return 7, []int{fourRank, kicker}
	}

	// Full house
	if crs[0].cnt == 3 && crs[1].cnt == 2 {
		return 6, []int{crs[0].rank, crs[1].rank}
	}

	// Flush
	if flush {
		// ranks already sorted desc
		return 5, ranks
	}

	// Straight
	if straight {
		return 4, []int{highStraight}
	}

	// Three of a kind
	if crs[0].cnt == 3 {
		threeRank := crs[0].rank
		kickers := make([]int, 0, 2)
		for _, r := range ranks {
			if r != threeRank {
				kickers = append(kickers, r)
			}
		}
		return 3, append([]int{threeRank}, kickers...)
	}

	// Two Pair
	if crs[0].cnt == 2 && crs[1].cnt == 2 {
		highPair := crs[0].rank
		lowPair := crs[1].rank
		kicker := crs[2].rank
		return 2, []int{highPair, lowPair, kicker}
	}

	// One Pair
	if crs[0].cnt == 2 {
		pairRank := crs[0].rank
		kickers := make([]int, 0, 3)
		for _, r := range ranks {
			if r != pairRank {
				kickers = append(kickers, r)
			}
		}
		return 1, append([]int{pairRank}, kickers...)
	}

	// High card
	return 0, ranks
}
