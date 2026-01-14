// Project Euler
package main

import (
	"fmt"
	"math"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/etnz/permute"
)

// Problem 93: Arithmetic Expressions
// Solution: 1258
func main() {
	start := time.Now()
	fmt.Printf("Problem 93: %s (%s)", solve(), time.Since(start))
}

func solve() string {
	var maxN int
	var maxNDigits []int
	// Create all combinations of 4 digits from 1 to 9
	d := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for p := range permute.Combinations(4, d) {
		n := countConsecutiveResults(p)
		if n > maxN {
			maxN = n
			maxNDigits = p
		}
	}
	return fmt.Sprintf("%d%d%d%d", maxNDigits[0], maxNDigits[1], maxNDigits[2], maxNDigits[3])
}

func countConsecutiveResults(digits []int) int {
	results := mapset.NewThreadUnsafeSet[float64]()
	d := make([]float64, len(digits))
	for i, digit := range digits {
		d[i] = float64(digit)
	}
	addAllResults(d, results)
	resultIntSet := convertToIntSet(results)
	for n := 1; ; n++ {
		if !resultIntSet.Contains(n) {
			return n - 1
		}
	}
}

func convertToIntSet(floatSet mapset.Set[float64]) mapset.Set[int] {
	intSet := mapset.NewThreadUnsafeSet[int]()
	for _, f := range floatSet.ToSlice() {
		if a, ok := floatToInt(f); ok {
			intSet.Add(a)
		}
	}
	return intSet
}

func floatToInt(f float64) (int, bool) {
	rounded := math.Round(f)
	if math.Abs(f - rounded) < 0.000001 {
		return int(rounded), true
	}
	return 0, false
}

func addAllResults(digits []float64, results mapset.Set[float64]) {
	if len(digits) == 1 {
		results.Add(digits[0])
		return
	}
	// Find pairs of indexes
	for i := 0; i < len(digits) - 1; i++ {
		a := digits[i]
		for j := i + 1; j < len(digits); j++ {
			b := digits[j]
			// Create new digit slice without a and b
			newDigits := make([]float64, 0, len(digits)-2)
			for k, d := range digits {
				if k != i && k != j {
					newDigits = append(newDigits, d)
				}
			}
			// Figure out all the results between a & b
			for _, r := range allResults(a, b) {
				// Recurse with new digit slice
				ds := []float64{r}
				ds = append(ds, newDigits...)
				addAllResults(ds, results)
			}
		}
	}
}

func allResults(a, b float64) []float64 {
	results := []float64{a+b, math.Abs(a-b), a*b}
	if b != 0 {
		results = append(results, a/b)
	}
	if a != 0 {
		results = append(results, b/a)
	}
	return results
}