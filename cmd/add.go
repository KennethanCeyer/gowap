package cmd

import (
	"github.com/urfave/cli"
	"github.com/KennethanCeyer/gowap/squeak"
)

// CommandAdd is used for swap to gowap profile
// Swap target profile must exists
// Usage: gowap add [profile]
func CommandAdd(c *cli.Context) error {
	// TODO:
	// validate current folder is gowap
	// validate target profile exists
	// sync if ENABLE_AUTO_SYNC option is enabled
	// delete current key
	// load target key

	var squeakInst = squeak.New()
	squeakInst.initialize()

	return nil
}
