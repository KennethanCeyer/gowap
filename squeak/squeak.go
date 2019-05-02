package squeak

import (
	"errors"
	"fmt"
	"github.com/atrox/homedir"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"path/filepath"
)

var (
	homeDir, err = homedir.Dir()
)

// Squeak is managing of ssh working directory
type Squeak struct {
	Workspace  string
	SSHKeyPath string
}

const (
	ProjectIdentifier = ".gowap"
	SSHPath           = ".ssh"
	SSHPrivateFile    = "id_rsa"
	SSHPublicFile     = "id_rsa.pub"
	NutsPath          = "nuts"
	defaultArchive    = "home"
)

var (
	ProjectPath          = filepath.Clean(fmt.Sprintf("%s/%s/%s", homeDir, SSHPath, ProjectIdentifier))
	ProjectReservedPaths = []string{NutsPath}
	SSHKeyPath           = filepath.Clean(fmt.Sprintf("%s/%s", homeDir, SSHPath))
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

func New() *Squeak {
	var squeak = new(Squeak)
	squeak.Workspace = ProjectPath
	squeak.SSHKeyPath = SSHKeyPath
	return squeak
}

func (s *Squeak) Initialize() error {
	err := s.createProjectIfEmpty()

	if err != nil {
		return err
	}

	return nil
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

	for _, reservedPath := range ProjectReservedPaths {
		subPath := path.Join(ProjectPath, reservedPath)
		if err = os.Mkdir(subPath, 0755); err != nil {
			return err
		}
	}

	return nil
}

func (s *Squeak) Store(privateKey string, publicKey string, archive string) error {
	s.Initialize()
	keyPaths := []string{privateKey, publicKey}
	for key, value := range []string{"private", "public"} {
		keyPath := keyPaths[key]
		if _, err := os.Stat(privateKey); os.IsNotExist(err) {
			return errors.New(fmt.Sprintf(ErrorStoreFileNotExists.Message, value, keyPath))
		}
	}

	if archive == "" {
		archive = defaultArchive
	}

	nuts := Nuts{privateKey, publicKey}
	return nuts.Save(archive)
}
