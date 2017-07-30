package main

import (
	"github.com/urfave/cli"
	"context"
	"errors"
)

const (
	GetPipelineHistoryCommandName  = "get-pipeline-history"
	GetPipelineHistoryCommandUsage = "Get Pipeline History"
)

func GetPipelineHistoryAction(c *cli.Context) error {
	if c.String("name") == "" {
		return handleOutput(nil, nil, "GetPipelineHistory", errors.New("'--name' is missing."))
	}

	pgs, r, err := cliAgent().Pipelines.GetHistory(context.Background(), c.String("name"), -1)
	if err != nil {
		return handleOutput(nil, r, "GetPipelineHistory", err)
	}

	return handleOutput(pgs, r, "GetPipelineHistory", err)
}

func GetPipelineHistoryCommand() *cli.Command {
	return &cli.Command{
		Name:     GetPipelineHistoryCommandName,
		Usage:    GetPipelineHistoryCommandUsage,
		Action:   GetPipelineHistoryAction,
		Category: "Pipelines",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "name"},
		},
	}
}
