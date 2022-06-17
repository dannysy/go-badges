package cmd

import (
	"github.com/urfave/cli/v2"

	"github.com/dannysy/go-badges/internal"
)

var CoverageCmd = cli.Command{
	Name:    "coverage-badge",
	Aliases: []string{"cb"},
	Usage:   "generates link to coverage badge from Golang test coverage report and updates README.md file",
	Action:  generateCoverageBadge,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "coverage-file",
			Usage:    "path to file with coverage test report",
			Required: true,
			Aliases:  []string{"cf"},
		},
		&cli.StringFlag{
			Name:    "readme-file",
			Usage:   "path to README.md",
			Value:   "README.md",
			Aliases: []string{"rf"},
		},
	},
}

func generateCoverageBadge(ctx *cli.Context) error {
	return internal.UpsertCoverageBadge(ctx.String("coverage-file"), ctx.String("readme-file"))
}
