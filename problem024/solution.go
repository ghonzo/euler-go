// Project Euler
package main

import (
	"fmt"
)

// Problem 24: Lexicographic Permutations
// Solution: 2783915460
func main() {
	fmt.Println("Problem 24:", solve())
}

func solve() []int {
	var count int
	return permute([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, &count, 1000000)
}

func permute(s []int, count *int, limit int) []int {
	var p []int
	if len(s) == 1 {
		(*count)++
		return s
	}
	for i := range s {
		rest := make([]int, 0)
		rest = append(rest, s[:i]...)
		rest = append(rest, s[i+1:]...)
		p = permute(rest, count, limit)
		if *count == limit {
			ret := make([]int, 0)
			ret = append(ret, s[i])
			ret = append(ret, p...)
			return ret
		}
	}
	return p
}
