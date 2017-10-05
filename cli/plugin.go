package cli

import (
	"context"
	"github.com/urfave/cli"
)

// List of command name and descriptions
const (
	ListPluginsCommandName  = "list-plugins"
	ListPluginsCommandUsage = "List all the Plugins"
	GetPluginCommandName    = "get-plugin"
	GetPluginCommandUsage   = "Get a Plugin"
)

// GetPluginAction retrieves a single plugin by name
func getPluginAction(c *cli.Context) cli.ExitCoder {
	pgs, r, err := cliAgent(c).Plugins.Get(context.Background(), c.String("name"))
	if err != nil {
		return NewCliError("GetPlugin", r, err)
	}
	return handleOutput(pgs, "GetPlugin")
}

// ListPluginsAction retrieves all plugin configurations
func listPluginsAction(c *cli.Context) cli.ExitCoder {
	pgs, r, err := cliAgent(c).Plugins.List(context.Background())
	if err != nil {
		return NewCliError("ListPlugins", r, err)
	}
	return handleOutput(pgs, "ListPlugins")
}

// GetPluginCommand Describes the cli interface for the GetPluginAction
func getPluginCommand() *cli.Command {
	return &cli.Command{
		Name:     GetPluginCommandName,
		Usage:    GetPluginCommandUsage,
		Category: "Plugins",
		Action:   getPluginAction,
		Flags: []cli.Flag{
			cli.StringFlag{Name: "name"},
		},
	}
}

// ListPluginsCommand Describes the cli interface for the ListPluginsCommand
func listPluginsCommand() *cli.Command {
	return &cli.Command{
		Name:     ListPluginsCommandName,
		Usage:    ListPluginsCommandUsage,
		Category: "Plugins",
		Action:   listPluginsAction,
	}
}
