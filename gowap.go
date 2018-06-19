package main

import (
	"os"
	"github.com/urfave/cli"
	"github.com/KennethanCeyer/gowap/cmd"
	gowapApp "github.com/KennethanCeyer/gowap/app"
)

var (
	app *cli.App
)

// Initialize gowap cli application
// This function has declaration of commands, options
func initApp(app *cli.App) *cli.App {
	app.Name = gowapApp.AppName
	app.Usage = gowapApp.AppUsage
	app.Version = gowapApp.AppVersion
	app.Author = gowapApp.Author

	app.Flags = []cli.Flag{}

	app.Commands = []cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "add ssh profile",
			Action:  cmd.CommandAdd,
		},
		{
			Name:    "remove",
			Aliases: []string{"r"},
			Usage:   "remove ssh profile",
			Action:  cmd.CommandRemove,
		},
		{
			Name:    "change",
			Aliases: []string{"c"},
			Usage:   "change ssh profile",
			Action:  cmd.CommandChange,
		},
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "show list ssh profiles",
			Action:  cmd.CommandList,
		},
		{
			Name:    "search",
			Aliases: []string{"s"},
			Usage:   "search ssh profile",
			Action:  cmd.CommandSearch,
		},
		{
			Name:    "archive",
			Aliases: []string{"h"},
			Usage:   "archive ssh profile",
			Action:  cmd.CommandArchive,
		},
		{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "show current gowap version",
			Action:  cmd.CommandVersion,
		},
	}

	return app
}

// This is entry point of gowap
// It will get arguments from user commands
func main() {
	app = initApp(cli.NewApp())
	app.Run(os.Args)
}
