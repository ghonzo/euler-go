// Project Euler
package main

import (
	"fmt"
	"math/big"
)

// Problem 26: Reciprocal Cycles
// Solution: 983
func main() {
	fmt.Println("Problem 26:", solve())
}

var maxCycle = 5000

func solve() int {
	ret, cycleLen := 7, 6
	for d := 11; d < 1000; d++ {
		r := big.NewRat(1, int64(d))
		// Use FloatPrec to determin if there are repeating digits
		if n, exact := r.FloatPrec(); !exact {
			// Ok so there are repeating digits. Let's find the cycle length
			str := r.FloatString(maxCycle * 2)
			if cl := findCycleLength(str[2+n : len(str)-1]); cl > cycleLen {
				ret, cycleLen = d, cl
			}
		}
	}
	return ret
}

func findCycleLength(s string) int {
outer:
	for cycleLen := 1; cycleLen < len(s)/2; cycleLen++ {
		for i := cycleLen; i < len(s); i++ {
			if s[i] != s[i%cycleLen] {
				continue outer
			}
		}
		return cycleLen
	}
	return 0
}
