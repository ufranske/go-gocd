package cli

import (
	"github.com/urfave/cli"
	"context"
)

const (
	ListEnvironmentsCommandName  = "list-environments"
	ListEnvironmentsCommandUsage = "List all environments"
)

func ListEnvironmentsAction(c *cli.Context) error {
	es, r, err := cliAgent(c).Environments.List(context.Background())
	if err != nil {
		return handleOutput(nil, r, "ListEnvironments", err)
	}

	es.RemoveLinks()

	return handleOutput(es, r, "ListEnvironments", err)
}

func ListEnvironmentsCommand() *cli.Command {
	return &cli.Command{
		Name:   ListEnvironmentsCommandName,
		Usage:  ListEnvironmentsCommandUsage,
		Action: ListEnvironmentsAction,
	}
}