/*
 * (C) 2017-2018 Kenneth Ceyer <kennethan@nhpcw.com>
 * this is distributed under
 * Apache 2.0 <https://www.apache.org/licenses/LICENSE-2.0>
 */

package cmd

import (
	"github.com/urfave/cli"
	"github.com/KennethanCeyer/gowap/squeak"
	"github.com/KennethanCeyer/gowap/cmd/add"
	"github.com/KennethanCeyer/gowap/app"
	"github.com/KennethanCeyer/gowap/peep"
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

	var log = peep.New(app.LogLevel)

	err := add.Validate(c)
	if err != nil {
		log.Error(err)
		return err
	}

	name := c.Args()[0]
	squeakInst := squeak.New()
	exists := squeakInst.Exists(name)

	if err != nil {
		return err
	}

	return nil
}
