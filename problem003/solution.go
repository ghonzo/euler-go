// Project Euler
package main

import (
	"fmt"
	"slices"

	"github.com/ghonzo/euler-go/common"
)

// Problem 3: Largest Prime Factor
// Solution: 6857
func main() {
	fmt.Println("Problem 3:", solve(600851475143))
}

func solve(n int) int {
	return slices.Max(common.PrimeFactors(n))
}
