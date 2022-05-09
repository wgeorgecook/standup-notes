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
