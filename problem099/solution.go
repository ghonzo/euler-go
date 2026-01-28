// Project Euler
package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ghonzo/euler-go/common"
)

// Problem 99: Largest Exponential
// Solution: Gives 849 which isn't right
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
	biggestNum := new(big.Int)
	var lineNumOfBiggest int
	var biggestBase, biggestExp int
	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		parts := strings.Split(line, ",")
		base := common.Atoi(parts[0])
		exp := common.Atoi(parts[0])
		if base < biggestBase && exp < biggestExp {
			continue
		}
		x := big.NewInt(int64(base))
		x.Exp(x, big.NewInt(int64(exp)), nil)
		if x.Cmp(biggestNum) > 0 {
			biggestNum = x
			lineNumOfBiggest = lineNum
			biggestBase = base
			biggestExp = exp
		}
	}
	return lineNumOfBiggest
}
