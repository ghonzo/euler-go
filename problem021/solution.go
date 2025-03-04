// Project Euler
package main

import (
	"fmt"

	"github.com/ghonzo/euler-go/common"
)

// Problem 21: Amicable Numbers
// Solution: 31626
func main() {
	fmt.Println("Problem 21:", solve())
}

func solve() int {
	var sum int
	// number pointing to sum of divisors
	sumMap := make(map[int]int)
	for i := 4; i < 10000; i++ {
		s := sumDivisors(i)
		sumMap[i] = s
		if s != i && sumMap[s] == i {
			sum += i + s
		}
	}
	return sum
}

func sumDivisors(n int) int {
	var sum int
	for _, x := range common.ProperDivisors(n) {
		sum += x
	}
	return sum
}
