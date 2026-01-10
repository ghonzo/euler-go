// Project Euler
package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"time"

	"github.com/ghonzo/euler-go/common"
	"github.com/oleiade/lane/v2"
)

// Problem 82: Path Sum: Three Ways
// Solution: 260324
func main() {
	start := time.Now()
	fmt.Printf("Problem 82: %d (%s)", solve(), time.Since(start))
}

type pos struct {
	row, col int
}

func solve() int {
	b, err := os.ReadFile("0082_matrix.txt")
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
	// Go from each cell in the leftmost column to the rightmost column
	minPathSum := math.MaxInt
	for startRow := range matrix {
		for endRow := range matrix {
			// We could definitely optimize this by pushing in the current minimum and aborting when we cross it
			pathSum := findMinPath(matrix, pos{startRow, 0}, pos{endRow, len(matrix[0]) - 1})
			if pathSum < minPathSum {
				minPathSum = pathSum
			}
		}
	}
	return minPathSum
}

func findMinPath(matrix [][]int, start, end pos) int {
	// Store minimal sum at each position
	minSum := make(map[pos]int)
	// The score is the current sum
	pq := lane.NewMinPriorityQueue[pos, int]()
	pq.Push(start, 0)
	for !pq.Empty() {
		currentPos, currentSum, _ := pq.Pop()
		// Add the current cell's value
		currentSum += matrix[currentPos.row][currentPos.col]
		// Are we at the end?
		if currentPos == end {
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