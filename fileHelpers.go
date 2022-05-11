package main

import (
	"os"
	"strings"
)

// loadFromFile opens the file at the given location and returns its
// contents. This reads the entire file into memory so make sure not
// to shoot yourself in the foot here.
func loadFromFile(path string) ([]byte, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return f, nil
}

// marshalFileContents takes in the bytes from a file read and
// attempts to marshal them into the teammates string slice
func marshalFileContents(b []byte) error {
	teammates = strings.Split(string(b), ",")
	return nil
}
