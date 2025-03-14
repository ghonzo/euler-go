// Project Euler
package main

import (
	"fmt"

	"github.com/ghonzo/euler-go/common"
)

// Problem 37: Truncatable Primes
// Solution:
func main() {
	fmt.Println("Problem 37:", solve())
}

func solve() int {
	var found int
	var sum int
outer:
	for i := 37; found < 11; i += 2 {
		// If it ends in 5 or 1, not gonna work
		if i%10 == 5 || i%10 == 1 {
			continue
		}
		// Now look every other digit
		var pow = 10
		for pow < i {
			n := i / pow
			switch n % 10 {
			case 0, 2, 4, 6, 8:
				i = ((i+pow)/10)*10 + 1
				continue outer
			case 5:
				i = ((i+pow*2)/10)*10 + 1
				continue outer
			}
			if n == 1 || !common.IsPrime(n) {
				continue outer
			}
			pow *= 10
		}
		// Now go the other way
		for pow > 10 {
			pow /= 10
			if !common.IsPrime(i % pow) {
				continue outer
			}
		}
		if common.IsPrime(i) {
			found++
			fmt.Println(found, i)
			sum += i
		}
	}
	return sum
}
