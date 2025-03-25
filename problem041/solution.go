// Project Euler
package main

import (
	"fmt"
	"strconv"

	"github.com/ghonzo/euler-go/common"
)

// Problem 41: Pandigital Prime
// Solution: 7652413
func main() {
	fmt.Println("Problem 41:", solve())
}

func solve() int {
	// This is some real brute force stuff here, but works with modern hardware!
	for n := 987654321; ; n-- {
		if isPandigital(n) && common.IsPrime(n) {
			return n
		}
	}
}

func isPandigital(n int) bool {
	s := strconv.Itoa(n)
	digits := make(map[int]bool)
	for _, r := range s {
		d := int(r - '0')
		if d == 0 || d > len(s) || digits[d] {
			return false
		}
		digits[d] = true
	}
	return true
}
