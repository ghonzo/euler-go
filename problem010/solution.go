// Project Euler
package main

import (
	"fmt"
)

// Problem 10: Summation of Primes
// Solution: 142913828922
func main() {
	fmt.Println("Problem 10:", solve2(2000000))
}

// This version takes a while
func solve(maxLimit int) int {
	primes := []int{2, 3}
	sum := 5
outer:
	for n := 5; n < maxLimit; n += 2 {
		for _, factor := range primes[1:] {
			if n%factor == 0 {
				continue outer
			}
		}
		primes = append(primes, n)
		sum += n
	}
	return sum
}

// This version is much faster
func solve2(maxLimit int) int {
	notPrime := make([]bool, maxLimit)
	var sum int
	for i := 2; i < maxLimit; i++ {
		if !notPrime[i] {
			sum += i
			for j := 2 * i; j < maxLimit; j += i {
				notPrime[j] = true
			}
		}
	}
	return sum
}
