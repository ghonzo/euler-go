// Project Euler
package main

import (
	"fmt"
)

// Problem 28: Number Spiral Diagonals
// Solution: 669171001
func main() {
	fmt.Println("Problem 28:", solve(1001))
}

func solve(dim int) int {
	sum := 1
	for i := 3; i <= dim; i += 2 {
		// Calculated that for each dim, add 4n^2 - 6n + 6
		sum += 4*i*i - 6*i + 6
	}
	return sum
}
