package common

import (
	"bufio"
	"os"
)

// ReadStringsFromFile expects a filename and returns a slice of strings, one per file in that file.
// It will panic is there is an error opening the file.
func ReadStringsFromFile(filename string) []string {
	input, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer input.Close()
	var strings []string
	s := bufio.NewScanner(input)
	for s.Scan() {
		strings = append(strings, s.Text())
	}
	return strings
}
