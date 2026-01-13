// Project Euler
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// Problem 89: Roman Numerals
// Solution: 743
func main() {
	start := time.Now()
	fmt.Printf("Problem 89: %d (%s)", solve(), time.Since(start))
}

func solve() int {
	b, err := os.ReadFile("0089_roman.txt")
	if err != nil {
		panic(err)
	}
	var saved int
	for _, line := range strings.Split(string(b), "\n") {	
		n := parseRoman(line)
		r := toRoman(n)
		saved += len(line) - len(r)
	}
	return saved
}

var romanValues = map[byte]int {
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func parseRoman(s string) int {
	var val int
	var lastVal int
	for _, b := range []byte(s) {
		v := romanValues[b]
		if lastVal != 0 && v > lastVal {
			val -= 2*lastVal
		}
		val += v
		lastVal = v
	}
	return val
}

func toRoman(n int) string {
	var result strings.Builder	
	for {
		// Could write less code here with a small struct
		switch {			
		case n >= 1000:
			result.WriteByte('M')
			n -= 1000
		case n >= 900:
			result.WriteString("CM")
			n -= 900
		case n >= 500:
			result.WriteByte('D')
			n -= 500
		case n >= 400:
			result.WriteString("CD")
			n -= 400
		case n >= 100:
			result.WriteByte('C')
			n -= 100
		case n >= 90:
			result.WriteString("XC")
			n -= 90
		case n >= 50:
			result.WriteByte('L')
			n -= 50
		case n >= 40:
			result.WriteString("XL")
			n -= 40
		case n >= 10:
			result.WriteByte('X')
			n -= 10
		case n >= 9:
			result.WriteString("IX")
			n -= 9
		case n >= 5:
			result.WriteByte('V')
			n -= 5
		case n >= 4:
			result.WriteString("IV")
			n -= 4
		case n >= 1:
			result.WriteByte('I')
			n -= 1
		default:
			return result.String()
		}
	}
}