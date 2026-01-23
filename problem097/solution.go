// Project Euler
package main

import (
	"fmt"
	"math/big"
	"time"
)

// Problem 97: Large Non-Mersenne Prime
// Solution: 8739992577
func main() {
	start := time.Now()
	fmt.Printf("Problem 97: %s (%s)", solve(), time.Since(start))
}

func solve() string {
	x := big.NewInt(2)
	x.Exp(x, big.NewInt(7830457), nil)
	x.Mul(x, big.NewInt(28433))
	x.Add(x, big.NewInt(1))
	x.Mod(x, big.NewInt(10000000000))
	return fmt.Sprintf("%010d", x.Uint64())
}