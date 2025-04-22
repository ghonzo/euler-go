// Project Euler
package common

type Digits struct {
	value  int
	digits []int
}

// DigitsFromInt returns a Digits type from a positive integer
func DigitsFromInt(n int) Digits {
	if n == 0 {
		return Digits{0, []int{0}}
	}
	d := Digits{n, make([]int, 0)}
	for n > 0 {
		d.digits = append(d.digits, n%10)
		n /= 10
	}
	return d
}

// Ints returns the digits of the number as a slice of integers
func (d Digits) Ints() []int {
	return d.digits
}

func (d Digits) Int() int {
	return d.value
}
