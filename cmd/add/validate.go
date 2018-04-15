package add

import "github.com/urfave/cli"

// add.Validate is used for validate arguments of add command
func Validate(c *cli.Context) error {
	if len(c.Args()) > 1 {
		// TODO: throw an exception properly
	}

	return nil
}
