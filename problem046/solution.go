// Project Euler
package main

import (
	"fmt"
)

// Problem 46: Goldbach's Other Conjecture
// Solution: 5777
func main() {
	fmt.Println("Problem 46:", solve())
}

const maxLimit = 100000

func solve() int {
	// First find all the prime numbers
	notPrime := make([]bool, maxLimit)
	for i := 2; i < maxLimit; i++ {
		if !notPrime[i] {
			for j := 2 * i; j < maxLimit; j += i {
				notPrime[j] = true
			}
		}
	}
	// Let's start with number 35
	for i := 35; i < maxLimit; i += 2 {
		// First, make sure it's composite
		if notPrime[i] {
			// Good. Start adding squares
			for j := 1; ; j++ {
				sq2 := 2 * j * j
				if sq2 >= i {
					// Found it
					return i
				}
				if !notPrime[i-sq2] {
					break
				}
			}
		}
	}
	panic("Didn't find it ... raise maxLimit")
}
