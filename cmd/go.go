package cmd

import (
	"github.com/urfave/cli"
)

// CommandGo is used for swap to gowap profile
// Swap target profile must exists
// Usage: gowap go [profile]
func CommandGo(c *cli.Context) error {
	// TODO:
	// validate current folder is gowap
	// validate target profile exists
	// sync if ENABLE_AUTO_SYNC option is enabled
	// delete current key
	// load target key
	return nil
}
