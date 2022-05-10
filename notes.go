package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/urfave/cli/v2"
)

var teammates = []string{"Your", "team's", "names", "go", "here"}

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

// remove pulls the element at index i from slice of strings s
// https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang
func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

// pop random removes a random element from slice of strings a and
// returns it as a string and the resulting slice of strings
func popRandom(a []string) (string, []string) {
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	randIdx := rand.Intn(len(a))
	popped := a[randIdx]
	rest := remove(a, randIdx)

	return popped, rest
}

// generateNotes accepts user input and builds the notes we want to
// return to the stdout
func generateNotes(s string) (string, error) {
	fmt.Printf("Notes for %s: ", s)
	// Taking input from user
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	in, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return "", err
	}
	return fmt.Sprintf("%s: %s", s, in), nil
}

// start is the default entrypoint to our app
func start(c *cli.Context) error {
	if fromFile != "" {
		// attempt to load from disk
		f, err := loadFromFile(fromFile)
		if err != nil {
			fmt.Printf("could not read from provided file: %v\n", err)
			return err
		}

		// attempt to unmarshal the file contents into a string slice
		if err := marshalFileContents(f); err != nil {
			fmt.Printf("could not marshal file contents to string slice: %v\n", err)
			return err
		}
	}
	if guest.Value() != nil {
		for _, member := range guest.Value() {
			teammates = append(teammates, member)
		}
	}
	fmt.Printf("today's team: %+v\n", teammates)
	teamCount := len(teammates) - 1
	var notes string = ""
	for i := 0; i <= teamCount; i += 1 {
		var member string
		member, teammates = popRandom(teammates)
		note, err := generateNotes(member)
		if err != nil {
			return err
		}
		notes += note
	}
	fmt.Println(strings.Trim(notes, "\n"))
	return nil
}
