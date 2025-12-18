// Project Euler
package main

import (
	"fmt"
	"time"
)

// Problem 69: Totient Maximum
// Solution: 510510
func main() {
	start := time.Now()
	fmt.Printf("Problem 69: %d (%s)", solve(1000000), time.Since(start))
}

func solve(maxN int) int {
	var maxTotient float64
	var maxTotientN int
	for n := 2; n <= maxN; n++ {
		phi := calculatePhi(n)
		totient := float64(n) / float64(phi)
		if totient > maxTotient {
			maxTotient = totient
			maxTotientN = n
		}
	}
	return maxTotientN
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
