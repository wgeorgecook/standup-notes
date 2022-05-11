package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

// selectTeammate uses the given notes-order strategy to determine
// which teammate we are building notes for in this instance
func selectTeammate(a []string, idx int) (string, []string) {
	switch order {
	case Alphabetical:
		return a[idx], a
	case ReverseAlphabetical:
		return pop(a, len(a)-1)
	case InPlace:
		return a[idx], a
	default:
		return popRandom(a)
	}

}

// buildTeammateNotes accepts user input and builds the notes we want to
// return to the stdout
func buildTeammateNotes(s string) (string, error) {
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
	if err := initConfig(); err != nil {
		return err
	}
	fmt.Printf("today's team: %+v\n", teammates)

	var notes string = ""
	fullCount := len(teammates) - 1
	for i := 0; i <= fullCount; i += 1 {
		var member string
		member, teammates = selectTeammate(teammates, i)
		note, err := buildTeammateNotes(member)
		if err != nil {
			return err
		}
		notes += note
	}
	fmt.Println(strings.Trim(notes, "\n"))
	return nil
}
