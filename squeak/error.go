/*
 * (C) 2017-2018 Kenneth Ceyer <kennethan@nhpcw.com>
 * this is distributed under
 * Apache 2.0 <https://www.apache.org/licenses/LICENSE-2.0>
 */

package squeak

import (
	"github.com/KennethanCeyer/gowap/exception"
)

const (
	ErrorCodeAleradyInitialized = 1
	ErrorCodeStore              = 6000
)

var (
	ErrorAlreadyInitialized = exception.Pair{Message: "%s is already initialized", Code: ErrorCodeAleradyInitialized}
	ErrorStoreFileNotExists = exception.Pair{Message: "%s ssh key path `` does not exists", Code: ErrorCodeStore}
)
