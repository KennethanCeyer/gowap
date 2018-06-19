/*
 * (C) 2017-2018 Kenneth Ceyer <kennethan@nhpcw.com>
 * this is distributed under
 * Apache 2.0 <https://www.apache.org/licenses/LICENSE-2.0>
 */

package squeak

import (
	"github.com/KennethanCeyer/gowap/app"
	"github.com/KennethanCeyer/gowap/exception"
)

var (
	ErrorAlreadyInitialized = exception.Pair { Message: "%s is already initialized", Code: 0x0001 }
	ErrorPathNotFound       = exception.Pair { Message: "'%s' path doesn't exists", Code: 0x0002 }
)

func AlreadyInitializedError() error {
	return exception.Error(ErrorAlreadyInitialized, app.AppName)
}

func PathNotFoundError(path string) error {
	return exception.Error(ErrorPathNotFound, path)
}
