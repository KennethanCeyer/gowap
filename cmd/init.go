package cmd

import (
	"fmt"
	"github.com/KennethanCeyer/gowap/squeak"
	"github.com/urfave/cli"
	"os/user"
	"path/filepath"
)

// CommandInit is used for initialize gowap file on current user path(PWD)
// It doesn't allowed duplication using and using in no ssh path
// Usage: gowap init
func CommandInit(c *cli.Context) error {
	// TODO:
	// - initialize current user directory
	// - validate it is .ssh
	// - validate a duplication of init using
	// - create meta files for gowap
	// - check keys (id_rsa, id_rsa.pub)
	// - set default profile automatically if it is exists
	usr, err := user.Current()
	if err != nil {
		return err
	}
	node := squeak.Node{Pwd: usr.HomeDir}
	rootNode, err := node.GetRoot()
	if err != nil {
		return err
	}
	if rootNode == nil {
		err = squeak.CreateProject(usr.HomeDir, squeak.SshIdentifier)
		if err != nil {
			return err
		}
	}
	return nil
}
