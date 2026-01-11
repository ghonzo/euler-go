// Project Euler
package main

import (
	"fmt"
	"math"
	"time"
)

// Problem Heegner (Bonus)
// Solution: Wrong so far (but maybe -163)
func main() {
	start := time.Now()
	fmt.Printf("Problem 84: %d (%s)", solve(), time.Since(start))
}

func solve() int {
	minDelta := 1.0
	var minN int
	for n := 2; n <= 1000; n++ {
		if isPefectSquare(n) {
			continue
		}
		h := math.Cos(math.Pi*math.Sqrt(float64(n)))
		delta := math.Min(math.Ceil(h) - h, h - math.Floor(h))
		if delta <= minDelta {
			minDelta = delta
			minN = n
		}
		// What about negative n?
		//v := (math.Exp(float64(n)) + math.Exp(-float64(n))) / 2.0
		v := math.Cosh(math.Pi*math.Sqrt(float64(n)))
		delta = math.Min(math.Ceil(v) - v, v - math.Floor(v))
		if delta <= minDelta {
			minDelta = delta
			minN = -n
		}
	}
	return minN
}

func isPefectSquare(n int) bool {
	sqrt := int(math.Sqrt(float64(n)))
	return sqrt*sqrt == n
}