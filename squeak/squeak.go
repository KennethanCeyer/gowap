package squeak

import (
	"github.com/atrox/homedir"
)

var (
	homeDir, err = homedir.Dir()
)

// Squeak is managing of ssh working directory
type Squeak struct {
	Workspace  string
	SSHKeyPath string
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

