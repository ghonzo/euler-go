// Project Euler
package main

import (
	"fmt"
	"time"
)

// Problem 100: Arranged Probability
// Solution: 756872327473
func main() {
	start := time.Now()
	fmt.Printf("Problem 100: %d (%s)", solve(), time.Since(start))
}

func solve() uint64 {
	// Two variables, b is number of blue, t is total
	// Equation relating them is 2b^2 - t^2 - 2b + t = 0
	// Put this in https://www.alpertron.com.ar/QUAD.HTM to get a generative function
	b := uint64(85)
	t := uint64(85+35)
	for t < uint64(1000000000000) {
		b, t = 3*b + 2*t - 2, 4*b + 3*t - 3
	}
	return b
}