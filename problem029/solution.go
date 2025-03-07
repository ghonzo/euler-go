// Project Euler
package main

import (
	"fmt"
	"math/big"
)

// Problem 29: Distinct Powers
// Solution: 9183
func main() {
	fmt.Println("Problem 29:", solve(100))
}

func solve(top int) int {
	found := make(map[string]bool)
	for a := int64(2); a <= int64(top); a++ {
		for b := int64(2); b <= int64(top); b++ {
			z := new(big.Int)
			found[z.Exp(big.NewInt(a), big.NewInt(b), nil).String()] = true
		}
	}
	return len(found)
}
