// Project Euler
package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ghonzo/euler-go/common"
)

// Problem 67: Maximum Path Sum II
// Solution: 7273
func main() {
	fmt.Println("Problem 67:", solve(common.ReadStringsFromFile("0067_triangle.txt")))
}

// This is actually the exact same algorithm as Problem 18
func solve(entries []string) int {
	rows := readInput(entries)
	for rowNum := len(rows) - 2; rowNum >= 0; rowNum-- {
		row := rows[rowNum]
		rowBelow := rows[rowNum+1]
		for colNum := range row {
			row[colNum] += max(rowBelow[colNum], rowBelow[colNum+1])
		}
	}
	return rows[0][0]
}

func readInput(entries []string) [][]int {
	var rows [][]int
	for _, rowStr := range entries {
		var row []int
		for _, nStr := range strings.Fields(rowStr) {
			n, _ := strconv.Atoi(nStr)
			row = append(row, n)
		}
		rows = append(rows, row)
	}
	return rows
}
