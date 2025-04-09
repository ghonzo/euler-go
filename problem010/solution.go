// Project Euler
package main

import (
	"fmt"

	"github.com/ghonzo/euler-go/common"
)

// Problem 10: Summation of Primes
// Solution: 142913828922
func main() {
	fmt.Println("Problem 10:", solve(2000000))
}

func solve(maxLimit int) int {
	notPrime := common.SieveNotPrimes(maxLimit)
	var sum int
	for k, v := range notPrime {
		if k > 1 && !v {
			sum += k
		}
	}
	return sum
}
