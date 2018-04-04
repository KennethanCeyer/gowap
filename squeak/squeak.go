package squeak

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

const (
	ProjectIdentifier = ".gowap"
	SshIdentifier     = ".ssh"
)

// Squeak is managing of working directory
// It searching the root of application working directory
type Squeak struct {
	Pwd string
}

func isProjectPath(path string) (bool, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return false, err
	}
	for _, file := range files {
		if file.IsDir() && file.Name() == ProjectIdentifier {
			return true, nil
		}
	}
	return false, nil
}

// Trace parent and when the parent doesn't exists or invalid it will return nil
func traceParent(squeak *Squeak) (*Squeak, bool, error) {
	parentPath := filepath.Join(squeak.Pwd, "..")
	parentNode := &Squeak{
		Pwd: parentPath,
	}
	tracePathIsProject, err := isProjectPath(parentPath)
	if err != nil {
		return nil, false, err
	}
	return parentNode, tracePathIsProject, nil
}

// Trace root and the parent is root or invalid it will be end
// This method is based of tail recursion
func traceRoot(squeak *Squeak) (*Squeak, error) {
	tracePathIsProject, err := isProjectPath(squeak.Pwd)
	if err != nil {
		return nil, err
	}
	if tracePathIsProject {
		return squeak, nil
	}
	parentNode, isRoot, err := traceParent(squeak)
	if err != nil {
		return nil, err
	}
	if isRoot {
		return parentNode, nil
	}
	if strings.Count(parentNode.Pwd, string(filepath.Separator)) <= 1 {
		return nil, nil
	}
	return traceRoot(parentNode)
}

// Trace root of working directory
// If root doesn't exists or invalid return will be nil
func (n *Squeak) GetRoot() (*Squeak, error) {
	return traceRoot(n)
}
