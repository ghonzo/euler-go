// Project Euler
package main

import (
	"fmt"
	"math/big"
	"slices"

	"github.com/ghonzo/euler-go/common"
)

// Problem 55: Lychrel Numbers
// Solution: 249
func main() {
	fmt.Println("Problem 55:", solve())
}

func solve() int {
	var count int
	for n := 1; n < 10000; n++ {
		if isLychrel(n) {
			count++
		}
	}
	return count
}

func isLychrel(n int) bool {
	num := big.NewInt(int64(n))
	for range 50 {
		numDigits := common.DigitsFromBigInt(num)
		slices.Reverse(numDigits)
		num.Add(num, numDigits.BigInt())
		if isPalindrome(num) {
			return false
		}
	}
	return true
}

func isPalindrome(n *big.Int) bool {
	d := common.DigitsFromBigInt(n)
	l := len(d)
	for i := 0; i*2 < l; i++ {
		if d[i] != d[l-i-1] {
			return false
		}
	}
	return true
}
