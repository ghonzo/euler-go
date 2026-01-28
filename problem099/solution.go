// Project Euler
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"time"

	"github.com/ghonzo/euler-go/common"
)

// Problem 99: Largest Exponential
// Solution: 709
func main() {
	start := time.Now()
	fmt.Printf("Problem 99: %d (%s)", solve(), time.Since(start))
}

func solve() int {
	file, err := os.Open("0099_base_exp.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	// Now go line by line
	var lineNum int
	var biggestNum float64
	var lineNumOfBiggest int
	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		parts := strings.Split(line, ",")
		base := float64(common.Atoi(parts[0]))
		exp := float64(common.Atoi(parts[1]))
		// ln(a^b) = b*ln(a)
		num := exp*math.Log(base)
		if num > biggestNum {
			biggestNum = num
			lineNumOfBiggest = lineNum
		}
	}
	return lineNumOfBiggest
}
