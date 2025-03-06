// Project Euler
package main

import (
	"fmt"

	"github.com/ghonzo/euler-go/common"
)

// Problem 27: Quadratic Primes
// Solution: -59231
func main() {
	fmt.Println("Problem 27:", solve())
}

var primeMap = make(map[int]bool)

func solve() int {
	var ret, numPrimes int
	for a := -999; a < 1000; a++ {
		for b := -1000; b <= 1000; b++ {
			np := findNumPrimes(a, b)
			if np > numPrimes {
				numPrimes = np
				ret = a * b
			}
		}
	}
	return ret
}

func findNumPrimes(a, b int) int {
	var prime, ok bool
	for n := 0; ; n++ {
		y := n*n + a*n + b
		if y < 2 {
			return n
		}
		if prime, ok = primeMap[y]; !ok {
			prime = common.IsPrime(y)
			primeMap[y] = prime
		}
		if !prime {
			return n
		}
	}
}
