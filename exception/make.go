/*
 * (C) 2017-2018 Kenneth Ceyer <kennethan@nhpcw.com>
 * this is distributed under
 * Apache 2.0 <https://www.apache.org/licenses/LICENSE-2.0>
 */

package exception

import "fmt"

// Error is generate new Error to using Pair struct
func Error(p Pair, a ...interface{}) error {
	return fmt.Errorf(fmt.Sprintf("exception(%d): %s", p.Code, p.Message), a...)
}

