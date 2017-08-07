package cli

import (
	"context"
	"errors"
	"github.com/urfave/cli"
)

// List of command name and descriptions
const (
	GetPipelineHistoryCommandName  = "get-pipeline-history"
	GetPipelineCommandName         = "get-pipeline"
	GetPipelineHistoryCommandUsage = "Get Pipeline History"
	GetPipelineCommandUsage        = "Get Pipeline"
)

// GetPipelineAction handles the business logic between the command objects and the go-gocd library.
func GetPipelineAction(c *cli.Context) error {
	if c.String("name") == "" {
		return handleOutput(nil, nil, "GetPipeline", errors.New("'--name' is missing"))
	}
	pgs, r, err := cliAgent().Pipelines.Get(context.Background(), c.String("name"), -1)
	if err != nil {
		return handleOutput(nil, r, "GetPipeline", err)
	}

	return handleOutput(pgs, r, "GetPipeline", err)

}

// GetPipelineHistoryAction handles the interaction between the cli flags and the action handler for
// get-pipeline-history-action
func GetPipelineHistoryAction(c *cli.Context) error {
	if c.String("name") == "" {
		return handleOutput(nil, nil, "GetPipelineHistory", errors.New("'--name' is missing"))
	}

	pgs, r, err := cliAgent().Pipelines.GetHistory(context.Background(), c.String("name"), -1)
	if err != nil {
		return handleOutput(nil, r, "GetPipelineHistory", err)
	}

	return handleOutput(pgs, r, "GetPipelineHistory", err)
}

// GetPipelineCommand handles the interaction between the cli flags and the action handler for
// get-pipeline
func GetPipelineCommand() *cli.Command {
	return &cli.Command{
		Name:     GetPipelineCommandName,
		Usage:    GetPipelineCommandUsage,
		Category: "Pipelines",
		Flags: [] cli.Flag{
			cli.StringFlag{Name: "name"},
		},
		Action: GetPipelineAction,
	}
}

// GetPipelineHistoryCommand handles the interaction between the cli flags and the action handler for
// get-pipeline-history-action
func GetPipelineHistoryCommand() *cli.Command {
	return &cli.Command{
		Name:     GetPipelineHistoryCommandName,
		Usage:    GetPipelineHistoryCommandUsage,
		Category: "Pipelines",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "name"},
		},
		Action: GetPipelineHistoryAction,
	}
}
