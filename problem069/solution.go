// Project Euler
package main

import (
	"fmt"
	"time"

	"github.com/ghonzo/euler-go/common"
)

// Problem 69: Totient Maximum
// Solution: 510510
func main() {
	start := time.Now()
	fmt.Printf("Problem 69: %d (%s)", solve(1000000), time.Since(start))
}

// This takes a long time to run, but it does complete.
func solve(maxN int) int {
	// For each n, store how many lower numbers are NOT relatively prime
	nonPrime := make([]int, maxN+1)
	// A temp slice that we zero out each time that keeps track of all multiple of prime factors
	isMultiple := make([]int, maxN+1)
	for n := 2; n <= maxN; n++ {
		clear(isMultiple)
		for _, pd := range common.ProperDivisors(n) {
			// Replace "1" with the actual number
			if pd == 1 {
				pd = n
			}
			for m := n; m <= maxN; m += pd {
				isMultiple[m] = 1
			}
		}
		for i := n; i <= maxN; i++ {
			nonPrime[i] += isMultiple[i]
		}
	}
	var maxTotient float64
	var maxTotientN int
	// Now calculate totients and find max
	for n := 2; n <= maxN; n++ {
		relativelyPrime := n - nonPrime[n]
		totient := float64(n) / float64(relativelyPrime)
		if totient > maxTotient {
			maxTotient = totient
			maxTotientN = n
		}
	}
	return maxTotientN
}
