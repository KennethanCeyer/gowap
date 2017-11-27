package cmd

import (
	"github.com/urfave/cli"
)

// CommandList is used for show the list of the gowap profiles
// This command only can use in gowap working directory
// Usage: gowap list
func CommandList(c *cli.Context) error {
	// TODO:
	// validate current folder is gowap
	// show all profile list that belongs current gowap
	return nil
}
