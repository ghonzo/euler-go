// Project Euler
package main

import (
	"fmt"
)

// Problem 30: Digit Fifth Powers
// Solution: 443839
func main() {
	fmt.Println("Problem 30:", solve())
}

func solve() int {
	var sum int
	// Let's try every number between 10 and 999,999
	for i := 10; i <= 999999; i++ {
		if isSolution(i) {
			sum += i
		}
	}
	return sum
}

func isSolution(n int) bool {
	var sum int
	for x := n; x > 0; x /= 10 {
		d := x % 10
		// We could memoize these if we needed more performance
		sum += d * d * d * d * d
	}
	return sum == n
}
