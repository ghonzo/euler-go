// Project Euler
package main

import (
	"fmt"
	"math/big"
	"time"
)

// Problem 71: Ordered Fractions
// Solution: 428570
func main() {
	start := time.Now()
	fmt.Printf("Problem 71: %d (%s)", solve(1000000), time.Since(start))
}

func solve(maxN int) int {
	target := big.NewRat(3, 7)
	found := big.NewRat(1, 3)
	for denom := int64(4); denom <= int64(maxN); denom++ {
		inc := 2 - (denom % 2)
		for num := denom * found.Num().Int64() / found.Denom().Int64(); num < denom; num += inc {
			r := big.NewRat(num, denom)
			if r.Cmp(target) >= 0 {
				break
			}
			if r.Cmp(found) > 0 {
				found = r
			}
		}
	}
	return int(found.Num().Int64())
}
