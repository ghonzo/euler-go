// Project Euler
package main

import (
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ghonzo/euler-go/common"
)

// Problem 47: Distinct Primes Factors
// Solution: 134043
func main() {
	fmt.Println("Problem 47:", solve(4))
}

// n is the number of consecutive numbers, and the number of distinct prime factors
func solve(n int) int {
	buffer := make([]map[int]int, n)
	for i := 8; ; i++ {
		buffer = buffer[1:]
		buffer = append(buffer, distinctPrimeFactors(i))
		if allDistinct(buffer, n) {
			return i - n + 1
		}
	}
}

func distinctPrimeFactors(num int) map[int]int {
	dpf := make(map[int]int)
	for _, p := range common.PrimeFactors(num) {
		dpf[p]++
	}
	return dpf
}

func allDistinct(buffer []map[int]int, n int) bool {
	type pf struct {
		prime, count int
	}
	superSet := mapset.NewThreadUnsafeSet[pf]()
	for _, dpf := range buffer {
		// Reject if not n prime factors
		if len(dpf) != n {
			return false
		}
		// Now add to the super set
		for p, c := range dpf {
			if !superSet.Add(pf{p, c}) {
				return false
			}
		}
	}
	return true
}
