package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

// flags
var (
	// guest indicates whether any special guests are joining us today
	guest cli.StringSlice

	// fromFile indicates if we should load our teammate slice from
	// the given file path
	fromFile string
)

func main() {
	app := &cli.App{
		Name:   "Standup Notes",
		Usage:  "Generate notes for daily standup",
		Action: start,
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:        "guest",
				Usage:       "name of any extra special guests joining standup today",
				Destination: &guest,
			},
			&cli.StringFlag{
				Name:        "from-file",
				Usage:       "path to load teammate string slice from",
				Destination: &fromFile,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
