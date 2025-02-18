// Project Euler
package main

import (
	"fmt"
)

// Problem 17: Number Letter Counts
// Solution: 21124
func main() {
	fmt.Println("Problem 17:", solve(1000))
}

func solve(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += spellLetters(i)
	}
	return sum
}

func spellLetters(n int) int {
	return len(spell(n))
}

var ones = [10]string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

var teens = [10]string{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}

var tens = [10]string{"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}

// Returns the number spelled out, with no spaces or hyphens
func spell(n int) string {
	// Max is 1000
	switch {
	case n == 0:
		return "zero"
	case n == 1000:
		return "onethousand"
	case n < 10:
		return ones[n]
	case n < 20:
		return teens[n-10]
	case n < 100:
		return concatSpell(tens[n/10], "", n%10)
	default:
		return concatSpell(ones[n/100]+"hundred", "and", n%100)
	}
}

func concatSpell(s, conjunction string, remainder int) string {
	if remainder == 0 {
		return s
	}
	return s + conjunction + spell(remainder)
}
