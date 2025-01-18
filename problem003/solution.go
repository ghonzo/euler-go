// Project Euler
package main

import (
	"fmt"
	"slices"
)

// Problem 3: Largest Prime Factor
// Solution: 6857
func main() {
	fmt.Println("Problem 3:", solve(600851475143))
}

func solve(n int) int {
	return slices.Max(primeFactors(n))
}

// Maybe a library function someday
func primeFactors(n int) []int {
	var factors []int
	// Take care of the 2s
	for n%2 == 0 {
		factors = append(factors, 2)
		n /= 2
	}
	for i := 3; i*i < n; i += 2 {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}
	if n > 1 {
		factors = append(factors, n)
	}
	return factors
}
