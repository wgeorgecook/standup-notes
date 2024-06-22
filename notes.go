package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/urfave/cli/v2"
)

// getUnderscores prints a - for every letter in teamname plus 2
func getUnderscores(teamName string) (underscores string) {
	underscores = "--"
	for i := 0; i <= len(teamName); i++ {
		underscores += "-"
	}
	return
}

// getTeamName accepts user input for the team name of the day
func getTeamName() (string, error) {
	fmt.Printf("Today's name selector: %s\n", getRandomTeamMember(teammates))
	fmt.Print("Team Name: ")
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	in, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return "", err
	}
	// be responsible if you're going up on a Tuesday
	if time.Now().Weekday() == time.Tuesday && stayResponsible {
		in = fmt.Sprintf("%s (Responsible Edition)\n", strings.ReplaceAll(in, "\n", ""))
	}
	return fmt.Sprintf("%s%s\n", in, getUnderscores(in)), nil
}

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
	fmt.Printf("Notes for %s: ", strings.TrimSpace(s))
	// Taking input from user
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	in, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return "", err
	}
	return fmt.Sprintf("%s: %s", strings.TrimSpace(s), in), nil
}

// start is the default entrypoint to our app
func start(c *cli.Context) error {
	if err := initConfig(); err != nil {
		return err
	}
	fmt.Printf("today's team: %s\n", teammates)
	var notes string = "" // starts the comment block for slack
	if setName {
		teamName, err := getTeamName()
		if err != nil {
			return err
		}
		notes += teamName
	}
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
	fmt.Println(notes)
	return nil
}
