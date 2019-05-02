package action

import (
	"fmt"
	"github.com/KennethanCeyer/gowap/squeak"
	"github.com/urfave/cli"
	"gopkg.in/AlecAivazis/survey.v1"
	"path"
)

var qs = []*survey.Question{
	{
		Name: "prikey",
		Prompt: &survey.Input{
			Message: "enter file in which to add the private key",
			Default: path.Join(squeak.SSHKeyPath, squeak.SSHPrivateFile)},
		Validate: survey.Required,
	},
	{
		Name: "pubkey",
		Prompt: &survey.Input{
			Message: "enter file in which to add the public key",
			Default: path.Join(squeak.SSHKeyPath, squeak.SSHPublicFile)},
		Validate: survey.Required,
	},
}

// CommandAdd is used for swap to gowap profile
// Swap target profile must exists
// Usage: gowap add [profile]
func CommandAdd(c *cli.Context) error {
	// TODO:
	// validate current folder is gowap
	// validate target profile exists
	// sync if ENABLE_AUTO_SYNC option is enabled
	// delete current key
	// load target key

	var squeakInst = squeak.New()
	err := squeakInst.Initialize()

	if err != nil {
		return err
	}

	answers := struct {
		PubKey string
		PriKey string
	}{}

	if err = survey.Ask(qs, &answers); err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
