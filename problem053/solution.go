// Project Euler
package main

import (
	"fmt"
	"math/big"
)

// Problem 53: Combinatoric Selections
// Solution: 4075
func main() {
	fmt.Println("Problem 53:", solve())
}

func solve() int {
	// Pre-compute all factorials up to 100
	factorials := make([]*big.Int, 101)
	factorials[0] = big.NewInt(1)
	for n := 1; n <= 100; n++ {
		factorials[n] = new(big.Int).Mul(factorials[n-1], big.NewInt(int64(n)))
	}
	oneMillion := big.NewInt(1000000)
	var count int
	for n := 1; n <= 100; n++ {
		for r := 0; r <= n; r++ {
			// Compute nCr = n! / (r! * (n - r)!)
			nCr := new(big.Int).Div(factorials[n], new(big.Int).Mul(factorials[r], factorials[n-r]))
			if nCr.Cmp(oneMillion) == 1 {
				count++
			}
		}
	}
	return count
}
