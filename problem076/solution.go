// Project Euler
package main

import (
	"fmt"
	"time"
)

// Problem 76: Counting Summations
// Solution: 190569291
func main() {
	start := time.Now()
	fmt.Printf("Problem 76: %d (%s)", solve(100), time.Since(start))
}

func solve(n int) int {
	ways := make([]int, n+1)
	ways[0] = 1
	// Count partitions using parts 1..n
	for part := 1; part <= n; part++ {
		for sum := part; sum <= n; sum++ {
			ways[sum] += ways[sum-part]
		}
	}
	// Exclude the partition consisting of the single number
	return ways[n] - 1
}
