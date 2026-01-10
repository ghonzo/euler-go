// Project Euler
package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/cockroachdb/apd"
)

// Problem 80: Square Root Digital Expansion
// Solution: 40886
func main() {
	start := time.Now()
	fmt.Printf("Problem 80: %d (%s)", solve(), time.Since(start))
}

func solve() int {
	var result int
	for n := 2; n <= 99; n++ {
		// Skip non-irrational square roots
		if n == 4 || n == 9 || n == 16 || n == 25 || n == 36 || n == 49 || n == 64 || n == 81 {
			continue
		}
		result += digitalSum(n)
	}
	return result
}

func digitalSum(n int) int {
	x := apd.New(int64(n), 0)
	ctx := apd.BaseContext.WithPrecision(110)
	d := new(apd.Decimal)
	ctx.Sqrt(d, x)
	s := d.Text('f')
	s = strings.Replace(s, ".", "", 1)
	var sum int
	for i := range 100 {
		sum += int(s[i] - '0')
	}
	return sum
}