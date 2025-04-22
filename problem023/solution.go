// Project Euler
package main

import (
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ghonzo/euler-go/common"
)

// Problem 23: Non-Abundant Sums
// Solution: 4179871
func main() {
	fmt.Println("Problem 23:", solve())
}

func solve() int {
	var abundantNumbers []int
	// First, find all abundant numbers
	for n := 12; n < 28123; n++ {
		if isAbundant(n) {
			abundantNumbers = append(abundantNumbers, n)
		}
	}
	// Now mark all numbers that can be the sum of two abundant numbers in a set
	sumSet := mapset.NewThreadUnsafeSet[int]()
	for i, n1 := range abundantNumbers {
		for _, n2 := range abundantNumbers[i:] {
			sum := n1 + n2
			if sum > 28123 {
				break
			}
			sumSet.Add(sum)
		}
	}
	// Finally sum every number not in the set
	var sum int
	for n := 1; n <= 28123; n++ {
		if !sumSet.Contains(n) {
			sum += n
		}
	}
	return sum
}

func isAbundant(n int) bool {
	var sum int
	for _, d := range common.ProperDivisors(n) {
		sum += d
	}
	return sum > n
}
