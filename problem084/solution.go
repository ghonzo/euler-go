// Project Euler
package main

import (
	"fmt"
	"maps"
	"math/rand/v2"
	"slices"
	"time"
)

// Problem 84: Monopoly Odds
// Solution: 101524
func main() {
	start := time.Now()
	fmt.Printf("Problem 84: %s (%s)", solve(4), time.Since(start))
}

const iter = 1000000

func solve(sides int) string {
	finished := make(map[int]int)
	pos := 0
	doubles := 0
	for range iter {
		d1 := rand.IntN(sides) + 1
		d2 := rand.IntN(sides) + 1
		if d1 == d2 {
			doubles++
			if doubles == 3 {
				// Go to jail
				pos = 10
				doubles = 0
				finished[pos]++
				continue
			}
		} else {
			doubles = 0
		}
		pos = (pos + d1 + d2) % 40
		switch pos {
		case 2, 17, 33: // Community Chest
			// Draw a card
			switch rand.IntN(16) {
			case 0: // Advance to GO
				pos = 0
			case 1: // Go to Jail
				pos = 10
			}
		case 7, 22, 36: // Chance
			switch rand.IntN(16) {
			case 0: // Advance to GO
				pos = 0
			case 1: // Go to Jail
				pos = 10
			case 2: // Go to C1
				pos = 11
			case 3: // Go to E3
				pos = 24
			case 4: // Go to H2
				pos = 39
			case 5: // Go to R1
				pos = 5
			case 6, 7: // Go to next R
				if pos == 7 {
					pos = 15
				} else if pos == 22 {
					pos = 25
				} else {
					pos = 5
				}
			case 8: // Go to next U
				if pos == 7 {
					pos = 12
				} else if pos == 22 {
					pos = 28
				} else {
					pos = 12				
				}
			case 9: // Go back 3 squares
				pos -= 3
			}
		case 30: // Go to Jail
			pos = 10
		}
		finished[pos]++
	}
	// Now find the top 3
	spaces := slices.SortedFunc(maps.Keys(finished), func(a, b int) int {
		return finished[b] - finished[a]
	})
	return fmt.Sprintf("%02d%02d%02d", spaces[0], spaces[1], spaces[2])
}
