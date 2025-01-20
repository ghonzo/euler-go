// Project Euler
package main

import (
	"fmt"

	"github.com/ghonzo/euler-go/common"
)

// Problem 5: Smallest Multiple
// Solution: 232792560
func main() {
	fmt.Println("Problem 5:", solve(1, 20))
}

func solve(low, high int) int {
	n := low * (low + 1)
	for i := low + 2; i <= high; i++ {
		n = common.LCM(n, i)
	}
	return n
}
