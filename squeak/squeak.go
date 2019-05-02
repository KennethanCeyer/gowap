package squeak

import (
	"errors"
	"fmt"
	"github.com/KennethanCeyer/gowap/utils"
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
	DefaultNutsName   = "home"
)

var (
	projectDir           = filepath.Clean(fmt.Sprintf("%s/%s/%s", homeDir, SSHPath, ProjectIdentifier))
	ProjectReservedPaths = []string{NutsPath}
	SSHKeyPath           = filepath.Clean(fmt.Sprintf("%s/%s", homeDir, SSHPath))
)

func New() *Squeak {
	var squeak = new(Squeak)
	squeak.Workspace = projectDir
	squeak.SSHKeyPath = SSHKeyPath
	return squeak
}

func (s *Squeak) GetProjectDir(key string) (string, error) {
	if err := s.CreateProjectIfEmpty(); err != nil {
		return "", err
	}
	projectDir := path.Join(projectDir, key)
	if err, _ = utils.CreateIfEmpty(projectDir); err != nil {
		return "", err
	}
	return projectDir, nil
}

func (s *Squeak) CreateProjectIfEmpty() error {
	err, created := utils.CreateIfEmpty(s.Workspace)
	if err != nil {
		return err
	}

	if created {
		log.Debugln("gowap is initialized")
	}

	return nil
}

func (s *Squeak) Store(privateKey string, publicKey string, nutsName string) error {
	keyPaths := []string{privateKey, publicKey}
	for key, value := range []string{"private", "public"} {
		keyPath := keyPaths[key]
		if _, err := os.Stat(privateKey); os.IsNotExist(err) {
			return errors.New(fmt.Sprintf(ErrorStoreFileNotExists.Message, value, keyPath))
		}
	}

	if nutsName == "" {
		nutsName = DefaultNutsName
	}

	nuts := Nuts{privateKey, publicKey}
	if projectDir, err = s.GetProjectDir(NutsPath); err != nil {
		return err
	}
	return nuts.Save(path.Join(projectDir, nutsName))
}
