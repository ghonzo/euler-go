// Project Euler
package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
)

// Problem 79: Passcode Derivation
// Solution: 73162890
func main() {
	start := time.Now()
	fmt.Printf("Problem 79: %s (%s)", solve(), time.Since(start))
}

type dep struct {
	before mapset.Set[byte]
	after  mapset.Set[byte]	
}

func solve() string {
	b, err := os.ReadFile("0079_keylog.txt")
	if err != nil {
		panic(err)
	}
	logins := strings.Split(string(b), "\n")
	var deps []*dep = make([]*dep, 10)
	for i := range 10 {
		deps[i] = &dep{
			before: mapset.NewSet[byte](),
			after:  mapset.NewSet[byte](),
		}
	}
	for _, login := range logins {
		a, b, c := login[0]-'0', login[1]-'0', login[2]-'0'
		deps[a].after.Append(b, c)
		deps[b].before.Add(a)
		deps[b].after.Add(c)
		deps[c].before.Append(a, b)
	}
	// Going to make a huge assumption here that each digit only shows up once
	// Find the last digit
	var lastDigit byte
	for i := range byte(10) {
		if !deps[i].before.IsEmpty() && deps[i].after.IsEmpty() {
			lastDigit = i
			break
		}
	}

	var result []byte	
	for {
		if next, ok := findNext(deps); ok {
			result = append(result, next+'0')
		} else {
			break
		}		
	}
	result = append(result, lastDigit+'0')
	return string(result)
}

func findNext(deps []*dep) (byte, bool) {
	for i := range byte(10) {
		if deps[i].before.IsEmpty() && !deps[i].after.IsEmpty() {
			// Remove this dependency from all others
			for j := range 10 {
				deps[j].before.Remove(i)
			}
			deps[i].after.Clear()
			return i, true
		}
	}
	return 0, false
}
