// Project Euler
package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ghonzo/euler-go/common"
	"github.com/oleiade/lane/v2"
)

// Problem 83: Path Sum: Four Ways
// Solution: 425185
func main() {
	start := time.Now()
	fmt.Printf("Problem 83: %d (%s)", solve(), time.Since(start))
}

type pos struct {
	row, col int
}

func solve() int {
	b, err := os.ReadFile("0083_matrix.txt")
	if err != nil {
		panic(err)
	}
	var matrix [][]int
	for _, line := range strings.Split(string(b), "\n") {
		var row []int
		for _, numStr := range strings.Split(line, ",") {
			row = append(row, common.Atoi(numStr))
		}
		matrix = append(matrix, row)
	}
	return solveMatrix(matrix)
}

func solveMatrix(matrix [][]int) int {
	// Store minimal sum at each position
	minSum := make(map[pos]int)
	// The score is the current sum
	pq := lane.NewMinPriorityQueue[pos, int]()
	pq.Push(pos{}, 0)
	for !pq.Empty() {
		currentPos, currentSum, _ := pq.Pop()
		// Add the current cell's value
		currentSum += matrix[currentPos.row][currentPos.col]
		// Are we at the end?
		if currentPos.row == len(matrix)-1 && currentPos.col == len(matrix[0])-1 {
			return currentSum
		}
		// Have we seen a lower sum here?
		if prevSum, found := minSum[currentPos]; found && prevSum <= currentSum {
			// We can skip this branch
			continue
		}
		minSum[currentPos] = currentSum
		// Now move to right if we can
		if currentPos.col < len(matrix[0])-1 {
			pq.Push(pos{currentPos.row, currentPos.col + 1}, currentSum)
		}
		// Now move to left if we can
		if currentPos.col > 0 {
			pq.Push(pos{currentPos.row, currentPos.col - 1}, currentSum)
		}
		// Now move down if we can
		if currentPos.row < len(matrix)-1 {
			pq.Push(pos{currentPos.row + 1, currentPos.col}, currentSum)
		}
		// Now move up if we can
		if currentPos.row > 0 {
			pq.Push(pos{currentPos.row - 1, currentPos.col}, currentSum)
		}
	}
	panic("no solution found")
}