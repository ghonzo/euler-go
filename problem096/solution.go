// Project Euler
package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// Problem 96: Su Doku
// Solution: 24702
func main() {
	start := time.Now()
	fmt.Printf("Problem 96: %d (%s)", solve(), time.Since(start))
}

func solve() int {
	var sum int
	f, err := os.Open("0096_sudoku.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	bufReader := bufio.NewReader(f)
	for {
		// Skip that first line, read the next 9 lines for the board
		s, err := bufReader.ReadString('\n')
		if len(s) == 0 || err == io.EOF {
			break
		}
		bv, err := convertToBoardValues(bufReader)
		if err != nil {
			panic(err)
		}
		b, err := createBoard(bv)
		if err != nil {
			panic(err)
		}
		bv, err = b.solve()
		if err != nil {
			panic(err)
		}
		sum += int(bv[0][0])*100 + int(bv[0][1])*10 + int(bv[0][2])
	}
	return sum
}

// The raw values for a board, row and col indexed
type boardValues [9][9]byte

func (bv boardValues) String() string {
	var sb strings.Builder
	for row := range 9 {
		for col := range 9 {
			v := bv[row][col]
			if v == 0 {
				sb.WriteByte('.')
			} else {
				sb.WriteByte(v + '0')

			}
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

// Expect 9x9 characters
func convertToBoardValues(br *bufio.Reader) (boardValues, error) {
	var bv boardValues
	// I expect 9 lines of 9 characters. Otherwise it's an error
	for row := range 9 {
		for col := range 9 {
			b, err := br.ReadByte()
			if err != nil {
				return bv, err
			}
			if b == '\n' {
				return bv, fmt.Errorf("expected 9 characters, but row %d only had %d chars", row+1, col)
			}
			if b >= '1' && b <= '9' {
				bv[row][col] = b - '0'
			} // Otherwise it'll already be 0
		}
		// Now throw out up to newline
		_, err := br.ReadBytes('\n')
		if err != nil {
			return bv, err
		}
	}
	return bv, nil
}

// immutable parts of the board
type cell struct {
	row, col       int
	val            byte
	peers          []*cell
	possibleValues []byte
}

type board [9][9]*cell

func createBoard(bv boardValues) (board, error) {
	var b board
	for row := range 9 {
		for col := range 9 {
			var c cell
			c.row = row
			c.col = col
			c.val = bv[row][col]
			b[row][col] = &c
		}
	}
	// Fill in peers
	for row := range 9 {
		for col := range 9 {
			b[row][col].peers = b.peerCells(b[row][col])
		}
	}
	// Validate the initial board
	if err := b.validate(); err != nil {
		return b, err
	}
	// Now iteratively figure out possible values
outer:
	for row := range 9 {
		for col := range 9 {
			if b[row][col].val == 0 {
				pv := b.findPossibleValues(b[row][col])
				switch len(pv) {
				case 0:
					return b, fmt.Errorf("impossible constraint at [%d][%d]", row+1, col+1)
				case 1:
					// Set this to be the value and start over
					b[row][col].val = pv[0]
					goto outer
				default:
					b[row][col].possibleValues = pv
				}
			}
		}
	}
	return b, nil
}

func (b board) peerCells(c *cell) []*cell {
	var peers []*cell
	// First do every cell in the row
	for col := range 9 {
		pc := b[c.row][col]
		if pc != c {
			peers = append(peers, pc)
		}
	}
	// Now every cell in the col
	for row := range 9 {
		pc := b[row][c.col]
		if pc != c {
			peers = append(peers, pc)
		}
	}
	// Now every cell in the "supercell"
	rowOffset := (c.row / 3) * 3
	colOffset := (c.col / 3) * 3
	for row := range 3 {
		if c.row == row+rowOffset {
			continue
		}
		for col := range 3 {
			if c.col == col+colOffset {
				continue
			}
			peers = append(peers, b[row+rowOffset][col+colOffset])
		}
	}
	return peers
}

func (b board) findPossibleValues(c *cell) []byte {
	unavailableValues := make(map[byte]bool)
	for _, pc := range c.peers {
		unavailableValues[pc.val] = true
	}
	var pv []byte
	for i := byte(1); i <= 9; i++ {
		if !unavailableValues[i] {
			pv = append(pv, i)
		}
	}
	return pv
}

func (b board) validate() error {
	// Make sure that for any fixed values there are no conflicting values. This checks values multiple times so it's inefficient
	for row := range 9 {
		for col := range 9 {
			v := b[row][col].val
			if v > 0 {
				for _, pc := range b[row][col].peers {
					if pc.val == v {
						return fmt.Errorf("conflicting values: [%d][%d] and [%d][%d] both have %d", row+1, col+1, pc.row+1, pc.col+1, v)
					}
				}
			}
		}
	}
	return nil
}

func (b board) String() string {
	return b.asBoardValues().String()
}

func (b board) asBoardValues() boardValues {
	var bv boardValues
	for row := range 9 {
		for col := range 9 {
			bv[row][col] = b[row][col].val
		}
	}
	return bv
}

func (b board) solve() (boardValues, error) {
	bv := b.asBoardValues()
	if !b.solvePosition(&bv, 0) {
		return bv, errors.New("No solution found")
	}
	return bv, nil
}

func (b board) solvePosition(bv *boardValues, pos int) bool {
	if pos == 9*9 {
		// Solved
		return true
	}
	row := pos / 9
	col := pos % 9
	if b[row][col].val > 0 {
		// Fixed value, so move on
		return b.solvePosition(bv, pos+1)
	}
	// Try each of the possible values
outer:
	for _, pv := range b[row][col].possibleValues {
		// Check each of the peer cells
		for _, pc := range b[row][col].peers {
			if bv[pc.row][pc.col] == pv {
				// Nope, can't be that one
				continue outer
			}
		}
		// Might work! Set it to that value and move to the next cell
		bv[row][col] = pv
		if b.solvePosition(bv, pos+1) {
			// Got it!
			return true
		}
	}
	// No luck. Set it back to 0 and return false
	bv[row][col] = 0
	return false
}
