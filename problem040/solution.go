// Project Euler
package main

import (
	"fmt"
	"strconv"
)

// Problem 40: Champernowne's Constant
// Solution: 210
func main() {
	fmt.Println("Problem 40:", solve())
}

func solve() int {
	var pos int
	prod := 1
	for i := 1; ; i++ {
		for _, r := range strconv.Itoa(i) {
			pos++
			switch pos {
			case 1, 10, 100, 1000, 10000, 100000, 1000000:
				prod *= int(r - '0')
			case 1000001:
				return prod
			}
		}
	}
}
