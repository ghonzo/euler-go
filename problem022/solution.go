// Project Euler
package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

// Problem 22: Names Scores
// Solution: 871198282
func main() {
	fmt.Println("Problem 22:", solve())
}

func solve() int {
	var total int
	b, err := os.ReadFile("0022_names.txt")
	if err != nil {
		panic(err)
	}
	names := strings.Split(strings.ReplaceAll(string(b), "\"", ""), ",")
	slices.Sort(names)
	for i, name := range names {
		total += score(name, i+1)
	}
	return total
}

func score(name string, pos int) int {
	var sum int
	for _, c := range []byte(name) {
		sum += int(c - 'A' + 1)
	}
	return sum * pos
}
