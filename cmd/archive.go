/*
 * (C) 2017-2018 Kenneth Ceyer <kennethan@nhpcw.com>
 * this is distributed under
 * Apache 2.0 <https://www.apache.org/licenses/LICENSE-2.0>
 */

package cmd

import (
	"github.com/urfave/cli"
)

// CommandArchive is used for archiving the gowap profile
// This command only can use in gowap working directory
// Usage: gowap archive
func CommandArchive(c *cli.Context) error {
	// TODO:
	// validate current folder is gowap
	// show all search list that belongs current gowap
	return nil
}
