// Project Euler
package main

import (
	"fmt"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ghonzo/euler-go/common"
)

// Problem 74: Digit Factorial Chains
// Solution: 402
func main() {
	start := time.Now()
	fmt.Printf("Problem 74: %d (%s)", solve(1000000), time.Since(start))
}

func solve(maxN int) int {
	var count int
	for n := 1; n < maxN; n++ {
		if chainLength(n) == 60 {
			count++
		}
	}
	return count
}

func chainLength(n int) int {
	termSet := mapset.NewThreadUnsafeSet(n)
	for {
		var sum int
		for _, d := range common.DigitsFromInt(n) {
			sum += factorialForDigit(d)
		}
		if !termSet.Add(sum) {
			return termSet.Cardinality()
		}
		n = sum
	}
}

var factorial [10]int

// This only works for d 0 to 9
func factorialForDigit(d int) int {
	if factorial[0] == 0 {
		// Initialize the factorials. No need for sync.Once in this evnvironment.
		factorial[0] = 1
		for i := 1; i <= 9; i++ {
			factorial[i] = factorial[i-1] * i
		}
	}
	return factorial[d]
}
