package squeak

import (
	"github.com/KennethanCeyer/gowap/exception"
)

var (
	ErrorMsgAlreadyInitialized = exception.Pair { "gowap is already initialized", 0x0001 }
	ErrorMsgPathNotFound       = exception.Pair { "'%s' path doesn't exists", 0x0002 }
)

func AlreadyInitializedError() error {
	return exception.Error(ErrorMsgAlreadyInitialized)
}

func PathNotFoundError(path string) error {
	return exception.Error(ErrorMsgPathNotFound)
}
