// Project Euler
package common

// Returns true if the given (positive) number is prime. This won't work very well for
// big numbers ... look at math.big#Int PossiblyPrime() for that.
func IsPrime(n int) bool {
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Returns a slice of booleans indicating which numbers are NOT prime up to maxLimit.
func SieveNotPrimes(maxLimit int) []bool {
	notPrime := make([]bool, maxLimit)
	for i := 2; i < maxLimit; i++ {
		if !notPrime[i] {
			for j := 2 * i; j < maxLimit; j += i {
				notPrime[j] = true
			}
		}
	}
	return notPrime
}

// Returns the first n prime numbers.
func Primes(n int) []int {
	primes := []int{2, 3}
outer:
	for i := 5; len(primes) < n; i += 2 {
		for _, factor := range primes[1:] {
			if i%factor == 0 {
				continue outer
			}
		}
		primes = append(primes, i)
	}
	return primes
}
