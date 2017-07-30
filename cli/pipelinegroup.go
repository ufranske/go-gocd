package cli

import (
	"context"
	"github.com/urfave/cli"
)

const (
	ListPipelineGroupsCommandName  = "list-pipeline-groups"
	ListPipelineGroupsCommandUsage = "List Pipeline Groups"
)

func ListPipelineGroupsAction(c *cli.Context) error {
	pgs, r, err := cliAgent().PipelineGroups.List(context.Background())
	if err != nil {
		return handleOutput(nil, r, "ListPipelineTemplates", err)
	}

	return handleOutput(pgs, r, "ListPipelineTemplates", err)
}

func ListPipelineGroupsCommand() *cli.Command {
	return &cli.Command{
		Name:     ListPipelineGroupsCommandName,
		Usage:    ListPipelineGroupsCommandUsage,
		Action:   ListPipelineGroupsAction,
		Category: "Pipeline Groups",
	}
}
