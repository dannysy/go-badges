package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/dannysy/go-badges/cmd"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{&cmd.CoverageCmd},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
