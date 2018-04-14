package squeak

import (
	"os"
	"github.com/KennethanCeyer/gowap/logger"
)

var loggerInst = logger.New()

func (s *Squeak) checkDirExists(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return PathNotFoundError(path)
		}
		return err
	}
	return nil
}

func (s *Squeak) createProjectIfEmpty() error {
	_, err := os.Stat(s.ProjectPath)
	if err != nil {
		return AlreadyInitializedError()
	}
	err = os.Mkdir(s.ProjectPath, 0755)
	if err != nil {
		return err
	}
	loggerInst.Debugln("%s file is created")
	return nil
}
