// Project Euler
package main

import (
	"fmt"
)

// Problem 7: 10 001st Prime
// Solution: 104743
func main() {
	fmt.Println("Problem 7:", solve(10001))
}

func solve(nth int) int {
	primes := []int{2, 3}
outer:
	for n := 5; len(primes) < nth; n += 2 {
		for _, factor := range primes[1:] {
			if n%factor == 0 {
				continue outer
			}
		}
		primes = append(primes, n)
	}
	return primes[len(primes)-1]
}
