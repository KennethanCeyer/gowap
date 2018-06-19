/*
 * (C) 2017-2018 Kenneth Ceyer <kennethan@nhpcw.com>
 * this is distributed under
 * Apache 2.0 <https://www.apache.org/licenses/LICENSE-2.0>
 */

package add

import "github.com/KennethanCeyer/gowap/exception"

var (
	ErrorArgumentMustBeOne  = exception.Pair { Message: "%s command's arguments must be one", Code: 0x1000 }
)

func AlreadyInitializedError() error {
	return exception.Error(ErrorArgumentMustBeOne, CommandName)
}
