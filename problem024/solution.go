// Project Euler
package main

import (
	"fmt"
	"slices"
	"strings"
)

// Problem 24: Lexicographic Permutations
// Solution: 2783915460
func main() {
	fmt.Println("Problem 24:", solve())
}

func solve() string {
	seq := permute([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, new(int), 1000000)
	var sb strings.Builder
	for _, d := range seq {
		sb.WriteByte('0' + byte(d))
	}
	return sb.String()
}

func permute(s []int, count *int, limit int) []int {
	var p []int
	if len(s) == 1 {
		(*count)++
		return s
	}
	for i := range s {
		p = permute(slices.Concat(s[:i], s[i+1:]), count, limit)
		if *count == limit {
			return slices.Insert(p, 0, s[i])
		}
	}
	return p
}
