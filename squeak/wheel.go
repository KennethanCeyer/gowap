package squeak

import (
	"os"
	"github.com/KennethanCeyer/gowap/peep"
	"github.com/KennethanCeyer/gowap/app"
	"fmt"
)

var log = peep.New(app.LogLevel)

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

	log.Debug(fmt.Sprintf("%v is initialized", app.AppName))
	return nil
}
