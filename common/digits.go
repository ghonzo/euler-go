// Project Euler
package common

type Digits []int

// DigitsFromInt returns a Digits type from a positive integer
func DigitsFromInt(n int) Digits {
	if n == 0 {
		return Digits{0}
	}
	digits := make(Digits, 0)
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}
	return digits
}
