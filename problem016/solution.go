// Project Euler
package main

import (
	"fmt"
	"math/big"
)

// Problem 16: Power Digit Sum
// Solution: 1366
func main() {
	fmt.Println("Problem 16:", solve(1000))
}

func solve(p int) int {
	var sum int
	z := new(big.Int)
	z = z.Exp(big.NewInt(2), big.NewInt(int64(p)), nil)
	for _, c := range []byte(z.String()) {
		sum += int(c - '0')
	}
	return sum
}
