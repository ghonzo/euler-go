// Project Euler
package main

import (
	"fmt"

	"github.com/ghonzo/euler-go/common"
)

// Problem 7: 10 001st Prime
// Solution: 104743
func main() {
	fmt.Println("Problem 7:", solve(10001))
}

func solve(nth int) int {
	return common.Primes(nth)[nth-1]
}
