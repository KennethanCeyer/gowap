package action

import (
	"fmt"
	"github.com/KennethanCeyer/gowap/exception"
	"github.com/KennethanCeyer/gowap/squeak"
	"github.com/KennethanCeyer/gowap/utils"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"gopkg.in/AlecAivazis/survey.v1"
	"path"
)

var addQs = []*survey.Question{
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
	nutsName := squeak.DefaultNutsName
	if c.NArg() > 0 {
		nutsName = c.Args().Get(0)
	}

	s := squeak.New()

	answers := struct {
		PubKey string
		PriKey string
	}{}

	if err := survey.Ask(addQs, &answers); err != nil {
		log.WithFields(log.Fields{
			"code": exception.ErrorCodeGeneral,
		}).Fatal(err)
		return err
	}

	nutsDir, err := s.GetProjectDir(path.Join(squeak.NutsPath, nutsName))
	if err != nil {
		log.WithFields(log.Fields{
			"code": squeak.ErrorCodeStore,
		}).Fatal(err)
		return err
	}

	if utils.Exists(nutsDir) {
		overwrite := true
		prompt := &survey.Confirm{
			Message: fmt.Sprintf("`%s` is already defined, Do you want to overwrite it?", nutsName),
			Default: overwrite,
		}
		survey.AskOne(prompt, &overwrite, nil)

		if overwrite {
			log.WithFields(log.Fields{
				"overwrite": overwrite,
			}).Info(fmt.Sprintf("try to overwrite to `%s`", nutsName))
		} else {
			log.WithFields(log.Fields{
				"overwrite": overwrite,
			}).Info("gowap add process has been canceled")
			return nil
		}
	}

	if err = s.Store(answers.PriKey, answers.PubKey, nutsName); err != nil {
		log.WithFields(log.Fields{
			"code": squeak.ErrorCodeStore,
		}).Fatal(err)
		return err
	}

	return nil
}
