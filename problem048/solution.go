// Project Euler
package main

import (
	"fmt"
	"math/big"
)

// Problem 48: Self Powers
// Solution: 9110846700
func main() {
	fmt.Println("Problem 48:", solve(1000))
}

func solve(n int) string {
	sum := big.NewInt(0)
	for i := 1; i <= n; i++ {
		pow := big.NewInt(int64(i))
		pow.Exp(pow, pow, nil)
		sum.Add(sum, pow)
	}
	str := sum.String()
	return str[len(str)-10:]
}
