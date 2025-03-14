// Project Euler
package main

import (
	"fmt"
	"strconv"
)

// Problem 36: Double-base Palindromes
// Solution: 872187
func main() {
	fmt.Println("Problem 36:", solve())
}

func solve() int {
	var sum int
	for i := 1; i < 1000000; i++ {
		if isPalindrome(strconv.Itoa(i)) && isPalindrome(strconv.FormatInt(int64(i), 2)) {
			sum += i
		}
	}
	return sum
}

func isPalindrome(s string) bool {
	n := len(s)
	for i := range n / 2 {
		if s[i] != s[n-i-1] {
			return false
		}
	}
	return true
}
