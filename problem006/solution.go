// Project Euler
package main

import "fmt"

// Problem 6: Sum Square Difference
// Solution: 25164150
func main() {
	fmt.Println("Problem 6:", solve(1, 100))
}

func solve(low, high int) int {
	var sum, sumSq int
	for i := low; i <= high; i++ {
		sum += i
		sumSq += i * i
	}
	return sum*sum - sumSq
}
