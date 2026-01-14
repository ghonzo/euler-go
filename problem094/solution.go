// Project Euler
package main

import (
	"fmt"
	"math"
	"time"
)

// Problem 94: Almost Equilateral Triangles
// Solution:
func main() {
	start := time.Now()
	fmt.Printf("Problem 94: %d (%s)", solve(1000000000), time.Since(start))
}

func solve(upperPerimeter int) int {
	var sum int
	// An isosoles triangle has two equal sides a and third side b
	for a := 2; ; a++ {
		b := a - 1
		p := 2*a + b
		if p > upperPerimeter {
			return sum
		}
		f := calcArea(a, b)
		if i, ok := floatToInt(f); ok {
			sum += i
		}
		b = a + 1
		p = 2*a + b
		if p > upperPerimeter {
			continue
		}
		f = calcArea(a, b)
		if i, ok := floatToInt(f); ok {
			sum += i
		}
	}
}

func floatToInt(f float64) (int, bool) {
	rounded := math.Round(f)
	if math.Abs(f - rounded) < 0.000001 {
		return int(rounded), true
	}
	return 0, false
}

func calcArea(a, b int) float64 {
	af, bf := float64(a), float64(b)
	return 0.5*bf*math.Sqrt(af*af-bf*bf/4)
}