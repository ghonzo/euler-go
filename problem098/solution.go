// Project Euler
package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/etnz/permute"
	"github.com/ghonzo/euler-go/common"
)

// Problem 98: Anagramic Squares
// Solution: 18769 (BOARD, BROAD)
func main() {
	start := time.Now()
	fmt.Printf("Problem 98: %d (%s)", solve(), time.Since(start))
}

func solve() int {
	b, err := os.ReadFile("0098_words.txt")
	if err != nil {
		panic(err)
	}
	// A "root" anagram (sorted by letter) pointing to all the words sharing those letters
	anagrams := make(map[string][]string)
	for w := range strings.SplitSeq(string(b), ",") {
		w = strings.Trim(w, "\"")
		root := sortString(w)
		anagrams[root] = append(anagrams[root], w)
	}
	// Sort by anagram length descending
	var anagramRootByLength []string
	for root, words := range anagrams {
		if len(words) < 2 {
			continue
		}
		// I'm going to make a huge assumption that no repeated letters
		if common.HasRepeatedElements([]byte(root)) || len(root) > 10 {
			continue
		}
		anagramRootByLength = append(anagramRootByLength, root)
	}
	// Now reverse sort by length
	slices.SortFunc(anagramRootByLength, func(a, b string) int {
		return len(b) - len(a)
	})
	// Now find all the squares by length
	initializeSquaresByLength()
	var maxSquare, maxSquareLength int
	for _, root := range anagramRootByLength {
		length := len(root)
		if length < maxSquareLength {
			// No need to check smaller lengths
			break
		}
		// Check all combinations of anagrams
		for pair := range permute.Combinations(2, anagrams[root]) {
			ms := findMaxSquareForAnagramPair(pair[0], pair[1])
			if ms > maxSquare {
				maxSquare = ms
				maxSquareLength = length
			}
		}
	}	
	return maxSquare
}

func sortString(s string) string {
	b := []byte(s)
	slices.Sort(b)
	return string(b)
}

var squaresByLength map[int]mapset.Set[int]

func initializeSquaresByLength() {
	squaresByLength = make(map[int]mapset.Set[int])
	for n := 4; ; n++ {
		nsq := n*n
		nsqDigits := common.DigitsFromInt(nsq)
		length := len(nsqDigits)
		if length > 10 {
			break
		}
		if common.HasRepeatedElements(nsqDigits) {
			continue
		}
		set, ok := squaresByLength[length]
		if !ok {
			set = mapset.NewSet[int]()
			squaresByLength[length] = set
		}
		set.Add(nsq)
	}
}

func findMaxSquareForAnagramPair(a, b string) int {
	var maxSquare int
	length := len(a)
	candidateSquares, ok := squaresByLength[length]
	if !ok || candidateSquares.Cardinality() < 2 {
		return 0
	}
	// Try each sequence of digits
	candidateSquares.Each(func(sq int) bool {
		d := common.DigitsFromInt(sq)
		mapping := make(map[rune]int)
		for i, r := range a {
			mapping[r] = d[i]
		}
		// Now find the value for b
		var bDigits common.Digits
		for _, r := range b {
			bDigits = append(bDigits, mapping[r])
		}
		// Leading zero check
		if bDigits[0] == 0 {
			// keep going
			return false
		}
		bInt := bDigits.Int()
		if candidateSquares.Contains(bInt) {
			maxSquare = max(maxSquare, sq, bInt)
		}
		return false
	})
	return maxSquare
}
