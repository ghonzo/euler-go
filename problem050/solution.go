// Project Euler
package main

import (
	"fmt"

	"github.com/ghonzo/euler-go/common"
)

// Problem 50: Consecutive Prime Sum
// Solution: 997651
func main() {
	fmt.Println("Problem 50:", solve(1000000))
}

func solve(limit int) int {
	notPrime := common.SieveNotPrimes(limit)
	var primes []int
	for k, v := range notPrime {
		if k >= 2 && !v {
			primes = append(primes, k)
		}
	}
	var maxLength, maxLengthPrime int
	for startIndex := range primes {
		var sum int
		for i := 0; startIndex+i < len(primes); i++ {
			sum += primes[startIndex+i]
			if sum > limit {
				break
			}
			if i > maxLength && !notPrime[sum] {
				maxLength = i
				maxLengthPrime = sum
			}
		}
	}
	return maxLengthPrime
}
