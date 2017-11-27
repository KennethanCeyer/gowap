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

// Node is managing of working directory
// It searching the root of application working directory
type Node struct {
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
func traceParent(node *Node) (*Node, bool, error) {
	parentPath := filepath.Join(node.Pwd, "..")
	parentNode := &Node{
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
func traceRoot(node *Node) (*Node, error) {
	tracePathIsProject, err := isProjectPath(node.Pwd)
	if err != nil {
		return nil, err
	}
	if tracePathIsProject {
		return node, nil
	}
	parentNode, isRoot, err := traceParent(node)
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
func (n *Node) GetRoot() (*Node, error) {
	return traceRoot(n)
}
