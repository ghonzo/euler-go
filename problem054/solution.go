// Project Euler
package main

import (
	"fmt"
	"strings"

	"github.com/ghonzo/euler-go/common"
)

// Problem 54: Poker Hands
// Solution: 376
func main() {
	fmt.Println("Problem 54:", solve(common.ReadStringsFromFile("0054_poker.txt")))
}

func solve(lines []string) int {
	var wins int
	for _, line := range lines {
		if player1Wins(line) {
			wins++
		}
	}
	return wins
}

func player1Wins(entry string) bool {
	var hand1, hand2 Hand
	cards := make([]Card, 0, 10)
	for _, code := range strings.Split(entry, " ") {
		card, err := NewCard(code)
		if err != nil {
			panic(err)
		}
		cards = append(cards, card)
	}
	hand1 = cards[:5]
	hand2 = cards[5:]
	return ComparePokerHands(hand1, hand2) > 0
}
