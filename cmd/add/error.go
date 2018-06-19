/*
 * (C) 2017-2018 Kenneth Ceyer <kennethan@nhpcw.com>
 * this is distributed under
 * Apache 2.0 <https://www.apache.org/licenses/LICENSE-2.0>
 */

package add

import (
	"fmt"
)

const (
	ErrorMsgArgumentMustBeOne  = "%s command's arguments must be one"
	ErrorCodeArgumentMustBeOne = 1000
)

func Error(msg string, code int, a ...interface{}) error {
	return fmt.Errorf(fmt.Sprintf("error(%d): %s", code, msg), a)
}

func AlreadyInitializedError() error {
	return Error(ErrorMsgArgumentMustBeOne, ErrorCodeArgumentMustBeOne, CommandName)
}
