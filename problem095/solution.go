// Project Euler
package main

import (
	"fmt"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ghonzo/euler-go/common"
)

// Problem 95: Amicable Chains
// Solution: 14316
func main() {
	start := time.Now()
	fmt.Printf("Problem 95: %d (%s)", solve(1000000), time.Since(start))
}

func solve(limit int) int {
	// Figure out each sum of proper divisors
	sumDivisors := make([]int, limit+1)
	for n := 6; n <= limit; n++ {
		sumDivisors[n] = sumOfProperDivisors(n)
		if sumDivisors[n] > limit {
			sumDivisors[n] = 0
		}
	}
	// Now find the longest chain
	var longestChain int
	var smallestInLongestChain int
	for n := 6; n <= limit; n++ {
		if clen, smallest := findChain(n, sumDivisors); clen > longestChain {
			longestChain = clen
			smallestInLongestChain = smallest
		}
	}
	return smallestInLongestChain
}

func sumOfProperDivisors(n int) int {
	var sum int
	for _, d := range common.ProperDivisors(n) {
		sum += d
	}
	return sum
}

func findChain(start int, sumDivisors []int) (int, int) {
	// Skip if lower
	next := sumDivisors[start]
	if next <= start {
		return 0,0
	}
	lowest := start
	seen := mapset.NewThreadUnsafeSet(start)
	for {
		if next == 0 {
			return 0,0
		}
		if next < lowest {
			lowest = next
		}
		if !seen.Add(next) {
			if next == start {
				return seen.Cardinality(), lowest
			}
			return 0,0
		}
		next = sumDivisors[next]
	}
}