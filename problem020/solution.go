// Project Euler
package main

import (
	"fmt"
	"math/big"
)

// Problem 20: Factorial Digit Sum
// Solution: 648
func main() {
	fmt.Println("Problem 20:", solve(100))
}

func solve(n int) int {
	z := new(big.Int)
	z = z.MulRange(1, int64(n))
	var sum int
	for _, d := range []byte(z.String()) {
		sum += int(d - '0')
	}
	return sum
}
