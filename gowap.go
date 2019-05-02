package main

import (
	"github.com/KennethanCeyer/gowap/action"
	gowap "github.com/KennethanCeyer/gowap/app"
	"github.com/urfave/cli"
	"os"
)

var (
	app *cli.App
)

// Initialize gowap cli application
// This function has declaration of commands, options
func initApp(app *cli.App) *cli.App {
	app.Name = gowap.AppName
	app.Usage = gowap.AppUsage
	app.Version = gowap.AppVersion
	app.Author = gowap.Author

	app.Flags = []cli.Flag{}

	app.Commands = []cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "add ssh profile",
			Action:  action.CommandAdd,
		},
		{
			Name:    "remove",
			Aliases: []string{"r"},
			Usage:   "remove ssh profile",
			Action:  action.CommandRemove,
		},
		{
			Name:    "change",
			Aliases: []string{"c"},
			Usage:   "change ssh profile",
			Action:  action.CommandChange,
		},
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "show list ssh profiles",
			Action:  action.CommandList,
		},
		{
			Name:    "search",
			Aliases: []string{"s"},
			Usage:   "search ssh profile",
			Action:  action.CommandSearch,
		},
		{
			Name:    "archive",
			Aliases: []string{"h"},
			Usage:   "archive ssh profile",
			Action:  action.CommandArchive,
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
