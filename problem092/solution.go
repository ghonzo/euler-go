// Project Euler
package main

import (
	"fmt"
	"time"

	"github.com/ghonzo/euler-go/common"
)

// Problem 92: Square Digit Chains
// Solution: 8581146
func main() {
	start := time.Now()
	fmt.Printf("Problem 92: %d (%s)", solve(10000000), time.Since(start))
}

func solve(upperLimit int) int {
	var found int
	for n := 1; n < upperLimit; n++ {
		v := n
		for {
			if v == 1 {
				break
			}
			if v == 89 {
				found++
				break
			}
			v = sumOfDigitSquares(v)
		}
	}
	return found
}

func sumOfDigitSquares(n int) int {
	var sum int
	for _, x := range common.DigitsFromInt(n) {
		sum += x * x
	}
	return sum
}