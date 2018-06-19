/*
 * (C) 2017-2018 Kenneth Ceyer <kennethan@nhpcw.com>
 * this is distributed under
 * Apache 2.0 <https://www.apache.org/licenses/LICENSE-2.0>
 */

package add

import "github.com/urfave/cli"

func Validate(c *cli.Context) error {
	if len(c.Args()) > 1 {
		return AlreadyInitializedError()
	}

	return nil
}
