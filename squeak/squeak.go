package squeak

import (
	"github.com/mitchellh/go-homedir"
	"fmt"
	"path/filepath"
)

const (
	ProjectIdentifier = ".gowap"
	SshIdentifier     = ".ssh"
)

var (
	homeDir, err = homedir.Dir()
)

// Squeak is managing of ssh working directory
type Squeak struct {
	ProjectPath string
	KeyPath string
}

func New() *Squeak {
	var squeak = new(Squeak)
	squeak.ProjectPath = filepath.Clean(fmt.Sprintf("%s/%s/%s", homeDir, SshIdentifier, ProjectIdentifier))
	squeak.KeyPath = filepath.Clean(fmt.Sprintf("%s/%s", homeDir, SshIdentifier))
	return squeak
}

func (s *Squeak) Initialize() error {
	err := s.createProjectIfEmpty()

	if err != nil {
		return err
	}

	return nil
}

