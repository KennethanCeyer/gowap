/*
 * (C) 2017-2018 Kenneth Ceyer <kennethan@nhpcw.com>
 * this is distributed under
 * Apache 2.0 <https://www.apache.org/licenses/LICENSE-2.0>
 */

package cmd

import (
	"github.com/urfave/cli"
)

// CommandChange is used for changing gowap profile option
// Change target profile must exists
// Usage: gowap change [profile]
func CommandChange(c *cli.Context) error {
	// TODO:
	// validate current folder is gowap
	// validate target profile
	// change option that profile
	return nil
}
