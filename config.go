package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/urfave/cli/v2"
)

// flags
var (
	// string slice of teammate names
	// yes this is a mutable global don't think of it too long
	teammates = []string{"Your", "team's", "names", "go", "here"}

	// CLI flags

	// guest indicates whether any special guests are joining us today
	guest cli.StringSlice

	// fromFile indicates if we should load our teammate slice from
	// the given file path
	fromFile string

	// order indicates the order we should cycle through the teammate
	// list
	order string

	// setName indicates that we should prompt for a team
	// name when starting notes
	setName bool
)

// order is a set of enums to type which order we cycle through
// teammate note entry
const (
	// Cycle through teammate list alphabetically
	Alphabetical = "alphabetical"
	// Cycle through teammate list in reverse
	ReverseAlphabetical = "reverse-alphabetical"
	// Randomly cycle through teammate list
	Random = "random"
	// Cycle through the teammate list as it is provided
	InPlace = "in-place"
)

// initConfig takes the values from the provided flags and performs
// logic resulting from the callers's configuration
func initConfig() error {
	// load teammates list from file if a file is provided
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

	// add any guests to the teammates array if guests are provided
	if guest.Value() != nil {
		for _, member := range guest.Value() {
			teammates = append(teammates, member)
		}
	}

	// if the order provided is sorted by the alphabet, make sure to
	// put it in alphabetical order so when we traverse it, things
	// are where we expect
	if order == Alphabetical || order == ReverseAlphabetical {
		// trim whitespace
		for i, s := range teammates {
			teammates[i] = strings.TrimSpace(s)
		}
		// sort the teammates list for processing
		sort.Strings(teammates)
	}

	return nil
}
