// Project Euler
package main

import (
	"fmt"
	"time"
)

// Problem 19: Counting Sundays
// Solution: 171
func main() {
	fmt.Println("Problem 19:", solve())
}

func solve() int {
	var count int
	for t := time.Date(1901, time.January, 1, 0, 0, 0, 0, time.UTC); t.Year() < 2001; t = t.AddDate(0, 1, 0) {
		if t.Weekday() == time.Sunday {
			count++
		}
	}
	return count
}
