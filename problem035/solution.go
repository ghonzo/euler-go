// Project Euler
package main

import (
	"fmt"
	"strconv"

	"github.com/ghonzo/euler-go/common"
)

// Problem 35: Circular Primes
// Solution: 55
func main() {
	fmt.Println("Problem 35:", solve(1000000))
}

func solve(limit int) int {
	cprimes := make(map[int]bool)
	cprimes[2] = true
outer:
	for n := 3; n < limit; n++ {
		if !cprimes[n] && common.IsPrime(n) {
			perms := []int{n}
			s := strconv.Itoa(n)
			for i := 1; i < len(s); i++ {
				s = s[1:] + s[:1]
				p, _ := strconv.Atoi(s)
				if !common.IsPrime(p) {
					continue outer
				}
				perms = append(perms, p)
			}
			for _, p := range perms {
				cprimes[p] = true
			}
		}
	}
	return len(cprimes)
}
