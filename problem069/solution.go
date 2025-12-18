// Project Euler
package main

import (
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ghonzo/euler-go/common"
)

// Problem 69: Totient Maximum
// Solution:
func main() {
	fmt.Println("Problem 69:", solve(1000000))
}

// This is naive. Do something like the sieve.
func solve(maxN int) int {
	var maxTotient float64
	var maxTotientN int
	// For each value of n, calculate and store the proper divisors, without 1
	// The index will be (n-2)
	var divisors []mapset.Set[int] = []mapset.Set[int]{mapset.NewThreadUnsafeSet(2)}
	for n := 3; n <= maxN; n++ {
		nDivisors := mapset.NewThreadUnsafeSet(common.ProperDivisors(n)[1:]...)
		// Number one is always there
		relativelyPrime := 1
		for p := 2; p < n; p++ {
			if !nDivisors.Contains(p) && !divisors[p-2].ContainsAnyElement(nDivisors) {
				relativelyPrime++
			}
		}
		totient := float64(n) / float64(relativelyPrime)
		if totient > maxTotient {
			maxTotient = totient
			maxTotientN = n
		}
		// Save them
		divisors = append(divisors, nDivisors)
	}
	return maxTotientN
}
