// Project Euler
package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Problem 18: Maximum Path Sum I
// Solution: 1074
func main() {
	fmt.Println("Problem 18:", solve(input))
}

var input = `75
95 64
17 47 82
18 35 87 10
20 04 82 47 65
19 01 23 75 03 34
88 02 77 73 07 63 67
99 65 04 28 06 16 70 92
41 41 26 56 83 40 80 70 33
41 48 72 33 47 32 37 16 94 29
53 71 44 65 25 43 91 52 97 51 14
70 11 33 28 77 73 17 78 39 68 17 57
91 71 52 38 17 14 91 43 58 50 27 29 48
63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23`

func solve(s string) int {
	rows := readInput(s)
	for rowNum := len(rows) - 2; rowNum >= 0; rowNum-- {
		row := rows[rowNum]
		rowBelow := rows[rowNum+1]
		for colNum := range row {
			row[colNum] += max(rowBelow[colNum], rowBelow[colNum+1])
		}
	}
	return rows[0][0]
}

func readInput(s string) [][]int {
	var rows [][]int
	for _, rowStr := range strings.Split(s, "\n") {
		var row []int
		for _, nStr := range strings.Fields(rowStr) {
			n, _ := strconv.Atoi(nStr)
			row = append(row, n)
		}
		rows = append(rows, row)
	}
	return rows
}
