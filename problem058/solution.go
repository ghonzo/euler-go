// Project Euler
package main

import (
	"fmt"

	"github.com/ghonzo/euler-go/common"
)

// Problem 58: Spiral Primes
// Solution: 26241
func main() {
	fmt.Println("Problem 58:", solve())
}

func solve() int {
	var primes int
	for sideLen := 3; ; sideLen += 2 {
		// Never prime
		bottomRight := sideLen * sideLen
		for i := 1; i < 4; i++ {
			corner := bottomRight - i*(sideLen-1)
			if common.IsPrime(corner) {
				primes++
			}
		}
		if float64(primes)/float64((sideLen-1)*2+1) < 0.1 {
			return sideLen
		}
	}
}
