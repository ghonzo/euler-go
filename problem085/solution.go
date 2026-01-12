// Project Euler
package main

import (
	"fmt"
	"math"
	"time"

	"github.com/ghonzo/euler-go/common"
)

// Problem 85: Counting Rectangles
// Solution: 2772
func main() {
	start := time.Now()
	fmt.Printf("Problem 85: %d (%s)", solve(), time.Since(start))
}

func solve() int {
	const target = 2000000
	// The number of rectangles in a w*h grid is Sum(1..w) * Sum(1..h)
	minDelta := math.MaxInt
	minDeltaArea := 0
	for h := 1; ; h++ {
		for w := h; ; w++ {
			nr := calcNumRectangles(w, h)			
			delta := common.Abs(target - nr)
			if delta < minDelta {
				minDelta = delta
				minDeltaArea = w * h
			} else if nr > target {
				if w == h {
					return minDeltaArea
				}
				break
			}
		}
	}
}

func calcNumRectangles(w, h int) int {
	// We could memoize this if needed
	return (w * (w + 1) / 2) * (h * (h + 1) / 2)
}