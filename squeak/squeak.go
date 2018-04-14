package squeak

import (
	"fmt"
)

const (
	ProjectIdentifier = ".gowap"
	SshIdentifier     = ".ssh"
)

// Squeak is managing of ssh working directory
type Squeak struct {
	ProjectPath string
	KeyPath string
}

func New() *Squeak {
	var squeak = new(Squeak)
	squeak.ProjectPath = fmt.Sprintf("~/%s", ProjectIdentifier)
	squeak.KeyPath = fmt.Sprintf("~/%s", SshIdentifier)
	return squeak
}

func (s *Squeak) initialize() error {
	s.createProjectIfEmpty()
	return nil
}
