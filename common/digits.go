// Project Euler
package common

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

func (d Digits) Int() int {
	var v int
	for _, x := range d {
		v = v*10 + x
	}
	return v
}
