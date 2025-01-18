// Project Euler
package main

import "fmt"

// Problem 2: Even Fibonacci Numbers
// Solution: 4613732
func main() {
	fmt.Println("Problem 2:", solve())
}

func solve() int {
	a, b, sum := 1, 2, 0
	for b < 4000000 {
		if b%2 == 0 {
			sum += b
		}
		a, b = b, a+b
	}
	return sum
}
