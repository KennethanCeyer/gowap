package add

import "github.com/urfave/cli"

func Validate(c *cli.Context) error {
	if len(c.Args()) > 1 {
		// TODO: throw an error properly
	}

	return nil
}
