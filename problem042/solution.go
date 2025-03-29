// Project Euler
package main

import (
	"fmt"
	"os"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
)

// Problem 42: Coded Triangle Numbers
// Solution: 162
func main() {
	fmt.Println("Problem 42:", solve())
}

func solve() int {
	b, err := os.ReadFile("0042_words.txt")
	if err != nil {
		panic(err)
	}
	words := strings.Split(strings.ReplaceAll(string(b), "\"", ""), ",")
	// Now find all the triangle numbers
	triSet := mapset.NewThreadUnsafeSet(1)
	for i := 1; i < 500; i++ {
		triSet.Add(i * (i + 1) / 2)
	}
	var count int
	for _, word := range words {
		if triSet.Contains(wordValue(word)) {
			count++
		}
	}
	return count
}

func wordValue(word string) int {
	var sum int
	for _, r := range word {
		sum += int(r - 'A' + 1)
	}
	return sum
}
