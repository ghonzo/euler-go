// Project Euler
package main

import (
	"fmt"

	"github.com/etnz/permute"
)

// Problem 43: Sub-string Divisibility
// Solution: 16695334890
func main() {
	fmt.Println("Problem 43:", solve())
}

func solve() int {
	var sum int
	d := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, p := range permute.Permutations(d) {
		if isSpecial(p) {
			sum += toInt(p)
		}
	}
	return sum
}

func isSpecial(p []int) bool {
	return p[0] != 0 &&
		isDiv(p, 1, 2) &&
		isDiv(p, 2, 3) &&
		isDiv(p, 3, 5) &&
		isDiv(p, 4, 7) &&
		isDiv(p, 5, 11) &&
		isDiv(p, 6, 13) &&
		isDiv(p, 7, 17)
}

func isDiv(p []int, start, div int) bool {
	return (p[start]*100+p[start+1]*10+p[start+2])%div == 0
}

func toInt(p []int) int {
	n := 0
	for _, d := range p {
		n = n*10 + d
	}
	return n
}
