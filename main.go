package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
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
			&cli.StringFlag{
				Name: "notes-order",
				Usage: "order to cycle through teammate array," +
					" options are 'alphabetical', 'reverse-alphabetical'" +
					"`in-place` (exactly as provided)" +
					", and 'random', with 'random' being the default behavior.",
				Destination: &order,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
