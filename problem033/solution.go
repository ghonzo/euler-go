// Project Euler
package main

import (
	"fmt"
	"math/big"
)

// Problem 33: Digit Cancelling Fractions
// Solution: 100
func main() {
	fmt.Println("Problem 33:", solve())
}

func solve() int {
	product := big.NewRat(1, 1)
	for a := 12; a <= 97; a++ {
		a10, a1 := a/10, a%10
		if a1 == 0 || a10 == a1 {
			continue
		}
		for b := a; b <= 98; b++ {
			b10, b1 := b/10, b%10
			if b1 == 0 || b10 == b1 {
				continue
			}
			r1 := big.NewRat(int64(a), int64(b))
			if (a10 == b1 && r1.Cmp(big.NewRat(int64(a1), int64(b10))) == 0) ||
				(a1 == b10 && r1.Cmp(big.NewRat(int64(a10), int64(b1))) == 0) {
				product.Mul(product, r1)
			}
		}
	}
	return int(product.Denom().Int64())
}
