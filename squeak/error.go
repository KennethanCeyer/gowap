package squeak

import (
	"fmt"
	"github.com/KennethanCeyer/gowap/app"
)

const (
	ErrorMsgAlreadyInitialized  = "%s is already initialized"
	ErrorCodeAlreadyInitialized = 001
	ErrorMsgPathNotFound        = "'%s' path doesn't exists"
	ErrorCodePathNotFound       = 100
)

func Error(msg string, code int, a ...interface{}) error {
	return fmt.Errorf(fmt.Sprintf("error(%d): %s", code, msg), a)
}

func AlreadyInitializedError() error {
	return Error(ErrorMsgAlreadyInitialized, ErrorCodeAlreadyInitialized, app.AppName)
}

func PathNotFoundError(path string) error {
	return Error(ErrorMsgPathNotFound, ErrorCodePathNotFound, path)
}
