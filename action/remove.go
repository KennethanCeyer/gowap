package action

import (
	"github.com/urfave/cli"
)

// CommandRemove is used for remove the gowap profile
// Delete target profile must exists
// Usage: gowap remove [profile]
func CommandRemove(c *cli.Context) error {
	// TODO:
	// validate current folder is gowap
	// validate target profile is exists
	// delete target profile
	return nil
}
