// Project Euler
package common

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

// Find all the prime factors (including repeats) of the given number
func PrimeFactors(n int) []int {
	var factors []int
	// Take care of the 2s
	for n%2 == 0 {
		factors = append(factors, 2)
		n /= 2
	}
	for i := 3; i*i < n; i += 2 {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}
	if n > 1 {
		factors = append(factors, n)
	}
	return factors
}

// Returns true if the given (positive) number is prime
func IsPrime(n int) bool {
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i < n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Returns all divisors of the number n, excluding n itself.
//
// For example, 1, 2, and 3 are proper divisors of 6.
func ProperDivisors(n int) []int {
	d := []int{1}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			d = append(d, i)
			if i*i != n {
				d = append(d, n/i)
			}
		}
	}
	return d
}
