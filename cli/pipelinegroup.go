package cli

import (
	"context"
	"github.com/urfave/cli"
)

// List of command name and descriptions
const (
	ListPipelineGroupsCommandName  = "list-pipeline-groups"
	ListPipelineGroupsCommandUsage = "List Pipeline Groups"
)

// ListPipelineGroupsAction handles the interaction between the cli flags and the action handler for
// list-pipeline-groups
func listPipelineGroupsAction(c *cli.Context) cli.ExitCoder {
	pgs, r, err := cliAgent(c).PipelineGroups.List(context.Background(), c.String("group-name"))
	if err != nil {
		return NewCliError("ListPipelineTemplates", r, err)
	}

	return handleOutput(pgs, "ListPipelineTemplates")
}

// ListPipelineGroupsCommand handles the interaction between the cli flags and the action handler for
// list-pipeline-groups
func listPipelineGroupsCommand() *cli.Command {
	return &cli.Command{
		Name:     ListPipelineGroupsCommandName,
		Usage:    ListPipelineGroupsCommandUsage,
		Action:   listPipelineGroupsAction,
		Category: "Pipeline Groups",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "group-name"},
		},
	}
}
