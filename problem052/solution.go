// Project Euler
package main

import (
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ghonzo/euler-go/common"
)

// Problem 52: Permuted Multiples
// Solution: 142857
func main() {
	fmt.Println("Problem 52:", solve())
}

func solve() int {
outer:
	for x := 1; ; x++ {
		xDigits := common.DigitsFromInt(x)
		xDigitsSet := mapset.NewThreadUnsafeSet(xDigits...)
		// Make sure they are all unique digits
		if xDigitsSet.Cardinality() != len(xDigits) {
			continue
		}
		for multiple := 2; multiple <= 6; multiple++ {
			multipleDigits := common.DigitsFromInt(x * multiple)
			multipleDigitsSet := mapset.NewThreadUnsafeSet(multipleDigits...)
			if !xDigitsSet.Equal(multipleDigitsSet) {
				continue outer
			}
		}
		return x
	}
}
