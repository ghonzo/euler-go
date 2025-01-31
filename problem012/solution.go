// Project Euler
package main

import (
	"fmt"

	"github.com/ghonzo/euler-go/common"
)

// Problem 12: Highly Divisible Triangular Number
// Solution: 76576500
func main() {
	fmt.Println("Problem 12:", solve(500))
}

func solve(n int) int {
	sum := 1
	for i := 2; ; i++ {
		sum += i
		if numFactors(sum) >= n {
			return sum
		}
	}
}

func numFactors(n int) int {
	primeFactors := common.PrimeFactors(n)
	// Now turn that into a map of the form n = a^p*b^q*c^r... with keys (a,b,c...) and values (p,q,r...)
	factorMap := make(map[int]int)
	for _, pf := range primeFactors {
		factorMap[pf]++
	}
	prod := 1
	for _, v := range factorMap {
		prod *= v + 1
	}
	return prod
}
