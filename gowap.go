package main

import (
	"github.com/KennethanCeyer/gowap/cmd"
	"github.com/urfave/cli"
	"os"
)

var (
	profile = ""
)

// Initialize gowap cli application
// This function has declaration of commands, options
func initApp(app *cli.App) *cli.App {
	app.Name = appName
	app.Usage = appUsage
	app.Version = appVersion
	app.Author = author

	app.Flags = []cli.Flag{}

	app.Commands = []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "initialize ssh path",
			Action:  cmd.CommandInit,
		},
		{
			Name:    "go",
			Aliases: []string{"g"},
			Usage:   "swap ssh profile",
			Action:  cmd.CommandGo,
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
	}

	return app
}

// This is entry point of gowap
// It will get arguments from user commands
func main() {
	app := initApp(cli.NewApp())
	app.Run(os.Args)
}
