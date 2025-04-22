// Project Euler
package main

import (
	"fmt"

	"github.com/ghonzo/euler-go/common"
)

// Problem 34: Digit Factorials
// Solution: 40730
func main() {
	fmt.Println("Problem 34:", solve())
}

func solve() int {
	// Cache the factorial values
	var fact [10]int
	fact[0] = 1
	fact[1] = 1
	for f := 2; f <= 9; f++ {
		fact[f] = fact[f-1] * f
	}
	var sum int
	// Can't be more than 9! x 7
	for v := 11; v < 2540160; v++ {
		var s int
		for _, d := range common.DigitsFromInt(v).Ints() {
			s += fact[d]
		}
		if s == v {
			sum += s
		}
	}
	return sum
}
