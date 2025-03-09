// Project Euler
package main

import (
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
)

// Problem 32: Pandigital Products
// Solution: 45228
func main() {
	fmt.Println("Problem 32:", solve())
}

type num struct {
	int
	digits mapset.Set[int]
}

func (n num) numDigits() int {
	return n.digits.Cardinality()
}

// Only creates numbers that have no repeating digits. Returns true if successful
func createNum(v int) (num, bool) {
	n := num{v, nil}
	s := mapset.NewThreadUnsafeSet[int]()
	for v > 0 {
		d := v % 10
		if d == 0 || !s.Add(d) {
			return n, false
		}
		v = v / 10
	}
	n.digits = s
	return n, true
}

func solve() int {
	var sum int
	found := mapset.NewThreadUnsafeSet[int]()
	for a := 2; a <= 98; a++ {
		aNum, ok := createNum(a)
		if !ok {
			continue
		}
		for b := a + 1; ; b++ {
			cNum, ok := createNum(a * b)
			if cNum.int > 9876 {
				break
			}
			if !ok || found.Contains(cNum.int) {
				continue
			}
			bNum, ok := createNum(b)
			if !ok {
				continue
			}
			if aNum.numDigits()+bNum.numDigits()+cNum.numDigits() != 9 {
				continue
			}
			if aNum.digits.Union(bNum.digits).Union(cNum.digits).Cardinality() == 9 {
				//fmt.Printf("%d x %d = %d\n", aNum.int, bNum.int, cNum.int)
				sum += cNum.int
				found.Add(cNum.int)
			}
		}
	}
	return sum
}
