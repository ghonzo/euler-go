// Project Euler
package main

import (
	"fmt"
	"time"
)

// Problem 72: Counting Fractions
// Solution: 303963552391
func main() {
	start := time.Now()
	fmt.Printf("Problem 72: %d (%s)", solve(1000000), time.Since(start))
}

// Insight! The number of reduced proper fractions for n is phi(n), which we solved for in
// Problems 69-70.
func solve(maxN int) int {
	var sum int
	for n := 2; n <= maxN; n++ {
		sum += calculatePhi(n)
	}
	return sum
}

func calculatePhi(n int) int {
	result := n
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			for n%i == 0 {
				n /= i
			}
			result -= result / i
		}
	}
	if n > 1 {
		result -= result / n
	}
	return result
}
