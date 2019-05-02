package squeak

import (
	"fmt"
	"os"
	"github.com/KennethanCeyer/gowap/logger"
	"path/filepath"
)

const (
	ProjectIdentifier = ".gowap"
	SSHPath        = ".ssh"
	SSHPrivateFile = "id_rsa"
	SSHPublicFile  = "id_rsa.pub"
)

var (
	log         = logger.New()
	ProjectPath = filepath.Clean(fmt.Sprintf("%s/%s/%s", homeDir, SSHPath, ProjectIdentifier))
	SSHKeyPath  = filepath.Clean(fmt.Sprintf("%s/%s", homeDir, SSHPath))
)

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
	if dirExists(s.Workspace) {
		return nil
	}

	err = os.Mkdir(s.Workspace, 0755)
	if err != nil {
		return err
	}

	log.Debugln("gowap is initialized")
	return nil
}
