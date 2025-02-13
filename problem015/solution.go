// Project Euler
package main

import (
	"fmt"
)

// Problem 15: Lattice Paths
// Solution: 137846528820
func main() {
	fmt.Println("Problem 15:", solve(20))
}

// So, this is a Pascal's Triangle. We're finding the nth element in the 2*nth row
func solve(n int) int {
	halfRow := []int{1, 2}
	for i := 1; i < n; i++ {
		halfRow = cycle(halfRow)
	}
	return halfRow[len(halfRow)-1]
}

// Calc the next two rows, return the second row
func cycle(a []int) []int {
	for i := len(a) - 1; i > 0; i-- {
		a[i] = a[i] + a[i-1]
	}
	b := make([]int, len(a)+1)
	b[0] = 1
	for i := 1; i < len(a); i++ {
		b[i] = a[i-1] + a[i]
	}
	b[len(a)] = a[len(a)-1] * 2
	return b
}
