// Project Euler
package main

import (
	"fmt"
	"time"

	"github.com/ghonzo/euler-go/common"
)

// Problem 73: Counting Fractions in a Range
// Solution: 7295372
func main() {
	start := time.Now()
	fmt.Printf("Problem 73: %d (%s)", solve(12000), time.Since(start))
}

func solve(maxN int) int {
	var count int
	for denom := 5; denom <= maxN; denom++ {
		inc := denom % 2
		for num := (denom / 3) + 1; num < (denom/2)+inc; num++ {
			if common.GCD(num, denom) == 1 {
				count++
			}
		}
	}
	return count
}
