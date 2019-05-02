package action

import (
	"github.com/KennethanCeyer/gowap/exception"
	"github.com/KennethanCeyer/gowap/squeak"
	log "github.com/sirupsen/logrus"
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

	var archive string
	if c.NArg() > 0 {
		archive = c.Args().Get(0)
	}

	s := squeak.New()
	err := s.Initialize()

	if err != nil {
		log.WithFields(log.Fields{
			"code": exception.ErrorCodeIntialze,
		}).Fatal(err)
		return err
	}

	answers := struct {
		PubKey string
		PriKey string
	}{}

	if err = survey.Ask(qs, &answers); err != nil {
		log.WithFields(log.Fields{
			"code": exception.ErrorCodeGeneral,
		}).Fatal(err)
		return err
	}

	if err = s.Store(answers.PriKey, answers.PubKey, archive); err != nil {
		log.WithFields(log.Fields{
			"code": squeak.ErrorCodeStore,
		}).Fatal(err)
		return err
	}

	return nil
}
