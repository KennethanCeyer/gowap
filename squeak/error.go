package squeak

import (
	"fmt"
)

const (
	ErrorMsgAlreadyInitialized  = "gowap is already initialized"
	ErrorCodeAlreadyInitialized = 001
	ErrorMsgPathNotFound        = "'%s' path doesn't exists"
	ErrorCodePathNotFound       = 100
)

func Error(msg string, code int, a ...interface{}) error {
	return fmt.Errorf(fmt.Sprintf("error(%d): %s", code, msg), a)
}

func AlreadyInitializedError() error {
	return Error(ErrorMsgAlreadyInitialized, ErrorCodeAlreadyInitialized)
}

func PathNotFoundError(path string) error {
	return Error(ErrorMsgPathNotFound, ErrorCodePathNotFound, path)
}
