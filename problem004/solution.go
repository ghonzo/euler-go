// Project Euler
package main

import (
	"fmt"
	"strconv"
)

// Problem 4: Largest Palindrome Product
// Solution: 906609
func main() {
	fmt.Println("Problem 4:", solve(100, 999))
}

func solve(low, high int) int {
	var largest int
outer:
	for a := high; a >= low && largest < a*high; a-- {
		for b := high; b >= low && largest < a*b; b-- {
			prod := a * b
			if isPalindrome(prod) {
				largest = prod
				continue outer
			}
		}
	}
	return largest
}

func isPalindrome(n int) bool {
	nStr := strconv.Itoa(n)
	l := len(nStr)
	for i := 0; i*2 < l; i++ {
		if nStr[i] != nStr[l-i-1] {
			return false
		}
	}
	return true
}
