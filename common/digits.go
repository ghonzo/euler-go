// Project Euler
package common

import "math/big"

type Digits []int

// DigitsFromInt returns a Digits type from a positive integer
func DigitsFromInt(n int) Digits {
	if n == 0 {
		return Digits([]int{0})
	}
	var d []int
	for n > 0 {
		d = append([]int{n % 10}, d...)
		n /= 10
	}
	return d
}

// DigitsFromBigInt returns a Digits type from a positive big.Int
func DigitsFromBigInt(n *big.Int) Digits {
	if n.Sign() == 0 {
		return Digits([]int{0})
	}
	var d []int
	ten := big.NewInt(10)
	for n.Sign() > 0 {
		var mod *big.Int
		n.DivMod(n, ten, mod)
		d = append([]int{int(mod.Int64())}, d...)
	}
	return d
}

// Int converts a Digits type back to an integer
func (d Digits) Int() int {
	var v int
	for _, x := range d {
		v = v*10 + x
	}
	return v
}

// BigInt converts a Digits type back to a big.Int
func (d Digits) BigInt() *big.Int {
	v := big.NewInt(0)
	ten := big.NewInt(10)
	for _, x := range d {
		v.Mul(v, ten)
		v.Add(v, big.NewInt(int64(x)))
	}
	return v
}
