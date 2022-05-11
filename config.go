package main

import (
	"fmt"

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
)

// initConfig takes the values from the provided flags and performs
// logic resulting from the callers's configuration
func initConfig() error {
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

	return nil
}
