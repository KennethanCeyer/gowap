/*
 * (C) 2017-2018 Kenneth Ceyer <kennethan@nhpcw.com>
 * this is distributed under
 * Apache 2.0 <https://www.apache.org/licenses/LICENSE-2.0>
 */

package cmd

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
