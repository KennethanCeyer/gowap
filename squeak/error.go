package squeak

import (
	"github.com/KennethanCeyer/gowap/exception"
)

var (
	ErrorMsgAlreadyInitialized = exception.Pair { Message: "gowap is already initialized", Code: 0x0001 }
	ErrorMsgPathNotFound       = exception.Pair { Message: "'%s' path doesn't exists", Code: 0x0002 }
)

func AlreadyInitializedError() error {
	return exception.Error(ErrorMsgAlreadyInitialized)
}

func PathNotFoundError(path string) error {
	return exception.Error(ErrorMsgPathNotFound)
}
