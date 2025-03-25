// Project Euler
package main

import (
	"fmt"
	"math"
)

// Problem 39: Integer Right Triangles
// Solution: 840
func main() {
	fmt.Println("Problem 39:", solve())
}

func solve() int {
	found := make(map[int]int)
	for a := 1; a < 1000; a++ {
		for b := a; b < 1000; b++ {
			c := math.Sqrt(float64(a*a + b*b))
			if math.Trunc(c) == c {
				found[a+b+int(c)]++
			}
		}
	}
	var maxP, numP int
	for p, count := range found {
		if p <= 1000 && count > numP {
			maxP, numP = p, count
		}
	}
	return maxP
}
