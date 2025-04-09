// Project Euler
package main

import (
	"fmt"
	"strconv"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ghonzo/euler-go/common"
)

// Problem 49: Prime Permutations
// Solution: 296962999629
func main() {
	fmt.Println("Problem 49:", solve(1489))
}

const maxLimit = 10000

func solve(lower int) string {
	// First find all the prime numbers
	notPrime := common.SieveNotPrimes(maxLimit)
	// Now work through the numbers
	for i := lower; i < maxLimit; i += 2 {
		if notPrime[i] {
			continue
		}
		digitSet := digits(i)
		for j := 2; i+j*2 < maxLimit; j += 2 {
			if notPrime[i+j] || notPrime[i+j*2] {
				continue
			}
			if digitSet.Equal(digits(i+j)) && digitSet.Equal(digits(i+j*2)) {
				return fmt.Sprintf("%d%d%d", i, i+j, i+j*2)
			}
		}
	}
	panic("no solution")
}

func digits(n int) mapset.Set[rune] {
	digitSet := mapset.NewThreadUnsafeSet[rune]()
	for _, r := range strconv.Itoa(n) {
		digitSet.Add(r)
	}
	return digitSet
}
