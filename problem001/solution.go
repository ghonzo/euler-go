// Project Euler
package main

import "fmt"

// Problem 1: Multiples of 3 or 5
// Solution: 233168
func main() {
	fmt.Println("Problem 1:", solve(1000))
}

func solve(n int) int {
	var sum int
	for i := 1; i < n; i++ {
		if (i%3 == 0) || (i%5 == 0) {
			sum += i
		}
	}
	return sum
}
