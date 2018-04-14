package squeak

import (
	"os"
	"github.com/KennethanCeyer/gowap/logger"
)

var loggerInst = logger.New()

func dirExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (s *Squeak) createProjectIfEmpty() error {
	if dirExists(s.ProjectPath) {
		return nil
	}

	err = os.Mkdir(s.ProjectPath, 0755)
	if err != nil {
		return err
	}

	loggerInst.Debugln("gowap is initialized")
	return nil
}
