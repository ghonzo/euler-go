// Project Euler
package main

import (
	"fmt"
)

// Problem 45: Triangular, Pentagonal, and Hexagonal
// Solution: 1533776805
func main() {
	fmt.Println("Problem 45:", solve())
}

func solve() int {
	t := 285
	p := 165
	h := 143
	for {
		// Increment h
		h++
		num := h * (2*h - 1)
		// Now increment p until we equal or exceed
		for {
			p++
			pNum := p * (3*p - 1) / 2
			if pNum > num {
				break
			}
			if pNum == num {
				// Now increment t until we equal or exceed
				for {
					t++
					tNum := t * (t + 1) / 2
					if tNum == num {
						return num
					}
					if tNum > num {
						break
					}
				}
			}
		}
	}
}
