// Project Euler
package main

import (
	"fmt"
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
	sum := 1
	for _, x := range divisors(n) {
		sum += x
	}
	return sum
}

func divisors(n int) []int {
	var d []int
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			d = append(d, i)
			if i*i != n {
				d = append(d, n/i)
			}
		}
	}
	return d
}
