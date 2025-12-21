// Project Euler
package main

import (
	"fmt"
	"time"

	"github.com/ghonzo/euler-go/common"
)

// Problem 75: Singular Integer Right Triangles
// Solution: 161667
func main() {
	start := time.Now()
	fmt.Printf("Problem 75: %d (%s)", solve(1500000), time.Since(start))
}

func solve(maxL int) int {
	// Key is length, value is number of solutions
	solutionsForL := make(map[int]int)
	// Use Euclid's formula to generate Pythagorean triples
outer:
	for m := 2; ; m++ {
		mOdd := m%2 == 1
		for n := 1; n < m; n++ {
			// Make sure that m and n are not both odd and are coprime
			if (mOdd && (n%2 == 1)) || common.GCD(m, n) != 1 {
				continue
			}
			for k := 1; ; k++ {
				a := k * (m*m - n*n)
				b := k * 2 * m * n
				c := k * (m*m + n*n)
				l := a + b + c
				if l > maxL {
					if n == 1 && k == 1 {
						break outer
					}
					break
				}
				solutionsForL[l]++
			}
		}
	}
	var count int
	for _, v := range solutionsForL {
		if v == 1 {
			count++
		}
	}
	return count
}
