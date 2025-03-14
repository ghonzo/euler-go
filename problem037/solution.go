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
		var pow = 10
		for pow < i {
			right := i % pow
			if right == 1 || !common.IsPrime(right) {
				continue outer
			}
			pow *= 10
		}
		for pow > 1 {
			pow /= 10
			left := i / pow
			if left == 1 || !common.IsPrime(left) {
				continue outer
			}
		}
		found++
		fmt.Println(found, i)
		sum += i
	}
	return sum
}
