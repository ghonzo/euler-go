// Project Euler
package main

import (
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ghonzo/euler-go/common"
)

// Problem 32: Pandigital Products
// Solution: 45228
func main() {
	fmt.Println("Problem 32:", solve())
}

func solve() int {
	var sum int
	found := mapset.NewThreadUnsafeSet[int]()
	for a := 2; a <= 98; a++ {
		aDigits := common.DigitsFromInt(a)
		if hasRepeatingDigits(aDigits) {
			continue
		}
		for b := a + 1; ; b++ {
			cDigits := common.DigitsFromInt(a * b)
			if cDigits.Int() > 9876 {
				break
			}
			if found.Contains(cDigits.Int()) || hasRepeatingDigits(cDigits) {
				continue
			}
			bDigits := common.DigitsFromInt(b)
			if hasRepeatingDigits(bDigits) {
				continue
			}
			if len(aDigits)+len(bDigits)+len(cDigits) != 9 {
				continue
			}
			digitSet := mapset.NewThreadUnsafeSet(aDigits...)
			digitSet.Append(bDigits...)
			digitSet.Append(cDigits...)
			if digitSet.Cardinality() == 9 {
				sum += cDigits.Int()
				found.Add(cDigits.Int())
			}
		}
	}
	return sum
}

// Also returns true if one of the digits is 0
func hasRepeatingDigits(d common.Digits) bool {
	s := mapset.NewThreadUnsafeSet[int]()
	for _, v := range d {
		if v == 0 || !s.Add(v) {
			return true
		}
	}
	return false
}
