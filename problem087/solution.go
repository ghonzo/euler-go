// Project Euler
package main

import (
	"fmt"
	"math"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ghonzo/euler-go/common"
)

// Problem 87: Prime Power Triples
// Solution: 1097343
func main() {
	start := time.Now()
	fmt.Printf("Problem 87: %d (%s)", solve(50000000), time.Since(start))
}

func solve(upperLimit int) int {
	// Get some primes to work with
	primes := common.PrimesUpTo(int(math.Sqrt(float64(upperLimit))))
	found := mapset.NewThreadUnsafeSet[int]()
	for _, c := range primes {
		c4 := c * c * c * c
		if c4 >= upperLimit {
			return found.Cardinality()
		}
		for _, b := range primes {
			b3 := b * b * b
			if c4+b3 >= upperLimit {
				break
			}
			for _, a := range primes {
				a2 := a * a
				sum := a2 + b3 + c4
				if sum >= upperLimit {
					break
				}
				found.Add(sum)
			}
		}
	}
	panic("should not reach here")
}