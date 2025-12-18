// Project Euler
package main

import (
	"fmt"
	"math"
	"slices"
	"time"

	"github.com/ghonzo/euler-go/common"
)

// Problem 70: Totient Permutation
// Solution: 8319823
func main() {
	start := time.Now()
	fmt.Printf("Problem 70: %d (%s)", solve(10000000), time.Since(start))
}

func solve(maxN int) int {
	var minTotient float64 = math.MaxFloat64
	var minTotientN int
	for n := 2; n <= maxN; n++ {
		phi := calculatePhi(n)
		totient := float64(n) / float64(phi)
		if totient < minTotient {
			// Now check to see if they are permutations
			if arePermutations(n, phi) {
				minTotient = totient
				minTotientN = n
			}
		}
	}
	return minTotientN
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

func arePermutations(a, b int) bool {
	aDigits := common.DigitsFromInt(a)
	slices.Sort(aDigits)
	bDigits := common.DigitsFromInt(b)
	slices.Sort(bDigits)
	return slices.Equal(aDigits, bDigits)
}
