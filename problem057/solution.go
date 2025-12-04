// Project Euler
package main

import (
	"fmt"
	"math/big"

	"github.com/ghonzo/euler-go/common"
)

// Problem 57: Square Root Convergents
// Solution: 153
func main() {
	fmt.Println("Problem 57:", solve())
}

func solve() int {
	var count int
	for n := 2; n <= 1000; n++ {
		exp := expansion(n)
		numDigits := len(common.DigitsFromBigInt(exp.Num()))
		denDigits := len(common.DigitsFromBigInt(exp.Denom()))
		if numDigits > denDigits {
			count++
		}
	}
	return count
}

// Please only call with n >= 2
func expansion(n int) *big.Rat {
	one := big.NewRat(1, 1)
	two := big.NewRat(2, 1)
	result := big.NewRat(5, 2)
	for range n - 2 {
		result.Inv(result)
		result.Add(two, result)
	}
	result.Inv(result)
	result.Add(one, result)
	return result
}
