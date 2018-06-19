package cmd

import "github.com/urfave/cli"

func CommandVersion(c *cli.Context) error {
	cli.ShowVersion(c)
	return nil
}
