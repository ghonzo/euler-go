// Project Euler
package main

import (
	"fmt"
)

// Problem 31: Coin Sums
// Solution: 73682
func main() {
	fmt.Println("Problem 14:", solve())
}

type coins struct {
	p1, p2, p5, p10, p20, p50, p100 int
}

func (c coins) value() int {
	return c.p1 + c.p2*2 + c.p5*5 + c.p10*10 + c.p20*20 + c.p50*50 + c.p100*100
}

func solve() int {
	seenStates := make(map[coins]bool)
	solutionStates := 1 // for the 2 pound case
	statesQueue := []coins{{}}
	for len(statesQueue) > 0 {
		c := statesQueue[0]
		statesQueue = statesQueue[1:]
		if seenStates[c] {
			continue
		}
		seenStates[c] = true
		remainder := 200 - c.value()
		if remainder == 0 {
			solutionStates++
			continue
		}
		if remainder >= 100 {
			c2 := c
			c2.p100++
			statesQueue = append(statesQueue, c2)
		}
		if remainder >= 50 {
			c2 := c
			c2.p50++
			statesQueue = append(statesQueue, c2)
		}
		if remainder >= 20 {
			c2 := c
			c2.p20++
			statesQueue = append(statesQueue, c2)
		}
		if remainder >= 10 {
			c2 := c
			c2.p10++
			statesQueue = append(statesQueue, c2)
		}
		if remainder >= 5 {
			c2 := c
			c2.p5++
			statesQueue = append(statesQueue, c2)
		}
		if remainder >= 2 {
			c2 := c
			c2.p2++
			statesQueue = append(statesQueue, c2)
		}
		c2 := c
		c2.p1++
		statesQueue = append(statesQueue, c2)
	}
	return solutionStates
}
