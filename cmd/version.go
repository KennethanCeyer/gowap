/*
 * (C) 2017-2018 Kenneth Ceyer <kennethan@nhpcw.com>
 * this is distributed under
 * Apache 2.0 <https://www.apache.org/licenses/LICENSE-2.0>
 */

package cmd

import "github.com/urfave/cli"

func CommandVersion(c *cli.Context) error {
	cli.ShowVersion(c)
	return nil
}
