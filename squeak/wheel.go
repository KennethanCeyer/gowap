package squeak

import (
	"github.com/KennethanCeyer/gowap/logger"
	"os"
	"path/filepath"
)

func checkDirExists(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return PathNotFoundError(path)
		}
		return err
	}
	return nil
}

func createProject(path string) error {
	projectPath := filepath.Join(path, ProjectIdentifier)
	_, err := os.Stat(projectPath)
	if err != nil {
		return AlreadyInitializedError()
	}
	err = os.Mkdir(path, 0755)
	if err != nil {
		return err
	}
	logger := logger.New()
	logger.Infof("%s file is created")
	return nil
}

func CreateProject(path string, i string) error {
	err := checkDirExists(path)
	if err != nil {
		return err
	}
	err = createProject(filepath.Join(path, i))
	if err != nil {
		return err
	}
	return nil
}
