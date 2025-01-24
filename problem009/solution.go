// Project Euler
package main

import "fmt"

// Problem 9: Special Pythagorean Triplet
// Solution: 31875000
func main() {
	fmt.Println("Problem 9:", solve())
}

func solve() int {
	for a := 1; a < 332; a++ {
		for b := a + 1; 1000-a > 2*b; b++ {
			c := 1000 - a - b
			if a*a+b*b == c*c {
				return a * b * c
			}
		}
	}
	panic("no solution")
}
