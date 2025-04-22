// Project Euler
package main

import (
	"fmt"

	"github.com/etnz/permute"
)

// Problem 24: Lexicographic Permutations
// Solution: 2783915460
func main() {
	fmt.Println("Problem 24:", solve())
}

func solve() string {
	var i int
	digits := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	// We need the 1,000,000th permutation of the digits 0-9
	for perm := range permute.LexPermutations(digits) {
		i++
		if i == 1000000 {
			return string(perm)
		}
	}
	panic("should never reach here")
}
