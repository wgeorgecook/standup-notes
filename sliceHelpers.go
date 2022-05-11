package main

import (
	"fmt"
	"math/rand"
	"time"
)

// remove pulls the element at index i from slice of strings s
// https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang
func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

// popRandom removes a random element from slice of strings a and
// returns it as a string and the resulting slice of strings
func popRandom(a []string) (string, []string) {
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	randIdx := rand.Intn(len(a))
	return pop(a, randIdx)
}

// pop removes the element at a given index in the provided slice.
// Returns the element and the remaining slice
func pop(a []string, i int) (string, []string) {
	popped := a[i]
	rest := remove(a, i)
	fmt.Printf("Strat: %s\npopped: %s\nrest: %v\n", order, popped, rest)
	return popped, rest
}
