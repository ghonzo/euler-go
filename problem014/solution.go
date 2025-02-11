// Project Euler
package main

import (
	"fmt"
)

// Problem 14: Longest Collatz Sequence
// Solution: 837799
func main() {
	fmt.Println("Problem 14:", solve())
}

var chainCache = make(map[int]int)

func solve() int {
	chainCache[1] = 1
	maxChainLength := 1
	maxChainNum := 1
	for n := 2; n < 1000000; n++ {
		chainLength := chain(n)
		if chainLength > maxChainLength {
			maxChainLength = chainLength
			maxChainNum = n
		}
	}
	return maxChainNum
}

func chain(n int) int {
	if v, ok := chainCache[n]; ok {
		return v
	}
	v := chain(next(n)) + 1
	chainCache[n] = v
	return v
}

func next(n int) int {
	if n%2 == 0 {
		return n / 2
	}
	return 3*n + 1
}
