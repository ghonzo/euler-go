// Project Euler
package main

import (
	"fmt"
	"strconv"

	"github.com/ghonzo/euler-go/common"
)

// Problem 38: Pandigital Multiples
// Solution: 932718654
func main() {
	fmt.Println("Problem 38:", solve())
}

func solve() int {
	var result int
	for n := 9; n < 9876; n++ {
		var s string
		for i := 1; i < 10; i++ {
			s += strconv.Itoa(n * i)
			if len(s) > 9 || !isPandigital(s) {
				break
			}
			if len(s) == 9 {
				result = max(result, common.Atoi(s))
				break
			}
		}
	}
	return result
}

// Returns true if it just consists of digits 1-9 with no repeats
func isPandigital(s string) bool {
	digits := make(map[rune]bool)
	for _, r := range s {
		if r == '0' || digits[r] {
			return false
		}
		digits[r] = true
	}
	return true
}
