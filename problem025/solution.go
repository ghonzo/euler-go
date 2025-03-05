// Project Euler
package main

import (
	"fmt"
	"math/big"
)

// Problem 25: 1000-digit Fibonacci Number
// Solution: 4782
func main() {
	fmt.Println("Problem 25:", solve(1000))
}

func solve(digits int) int {
	a := big.NewInt(1)
	b := big.NewInt(1)
	for i := 3; ; i++ {
		c := new(big.Int)
		c.Add(a, b)
		if len(c.String()) == digits {
			return i
		}
		a, b = b, c
	}
}
