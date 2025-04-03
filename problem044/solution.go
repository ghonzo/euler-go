// Project Euler
package main

import (
	"fmt"
	"math"

	mapset "github.com/deckarep/golang-set/v2"
)

// Problem 44: Pentagon Numbers
// Solution: 5482660
func main() {
	fmt.Println("Problem 44:", solve())
}

func solve() int {
	// Let's just brute force this and see what we get
	pNumSet := mapset.NewThreadUnsafeSet[int]()
	for n := 1; n < 10000; n++ {
		pNumSet.Add(n * (3*n - 1) / 2)
	}
	pNumSlice := mapset.Sorted(pNumSet)
	minDiff := math.MaxInt
	for k, pk := range pNumSlice {
		for _, pj := range pNumSlice[k+1:] {
			if pj-pk > minDiff {
				break
			}
			if pNumSet.Contains(pk+pj) && pNumSet.Contains(pj-pk) {
				minDiff = min(minDiff, pj-pk)
			}
		}
	}
	return minDiff
}
