package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

// guest indicates whether any special guests are joining us today
var guest cli.StringSlice

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
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
