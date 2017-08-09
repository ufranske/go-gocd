package cli

import (
	"context"
	"github.com/urfave/cli"
)

const (
	ListPluginsCommandName  = "list-plugins"
	ListPluginsCommandUsage = "List all the Plugins"
	GetPluginCommandName    = "get-plugin"
	GetPluginCommandUsage   = "Get a Plugin"
)

func GetPluginAction(c *cli.Context) error {
	pgs, r, err := cliAgent(c).Plugins.Get(context.Background(), c.String("name"))
	if err != nil {
		return handleOutput(nil, r, "GetPlugin", err)
	}

	return handleOutput(pgs, r, "ListPipelineTemplates", err)
}

func ListPluginsAction(c *cli.Context) error {
	pgs, r, err := cliAgent(c).Plugins.List(context.Background())
	if err != nil {
		return handleOutput(nil, r, "ListPlugins", err)
	}

	return handleOutput(pgs, r, "ListPlugins", err)
}

func GetPluginCommand() *cli.Command {
	return &cli.Command{
		Name:     GetPluginCommandName,
		Usage:    GetPluginCommandUsage,
		Category: "Plugins",
		Action:   GetPluginAction,
		Flags: []cli.Flag{
			cli.StringFlag{Name: "name"},
		},
	}
}

func ListPluginsCommand() *cli.Command {
	return &cli.Command{
		Name:     ListPluginsCommandName,
		Usage:    ListPluginsCommandUsage,
		Category: "Plugins",
		Action:   ListPluginsAction,
	}
}
