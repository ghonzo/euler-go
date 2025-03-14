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
		if common.IsPrime(i) {
			var pow = 10
			for {
				right := i / pow
				if right == 0 {
					break
				}
				left := i % pow
				if left == 1 || right == 1 || !common.IsPrime(left) || !common.IsPrime(right) {
					continue outer
				}
				pow *= 10
			}
			found++
			fmt.Println(found, i)
			sum += i
		}
	}
	return sum
}
