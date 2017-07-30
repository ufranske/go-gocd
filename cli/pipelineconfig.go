package main

import (
	"github.com/urfave/cli"
	"context"
	"github.com/drewsonne/gocdsdk"
	"encoding/json"
	"github.com/pkg/errors"
)

const (
	CreatePipelineConfigCommandName  = "create-pipeline-config"
	CreatePipelineConfigCommandUsage = "Create Pipeline config"
)

func CreatePipelineConfigAction(c *cli.Context) error {
	group := c.String("group")
	if group == "" {
		return handleOutput(nil, nil, "CreatePipelineConfig", errors.New("'--group' is missing."))
	}

	p := &gocd.Pipeline{}
	pipeline := c.String("pipeline")
	pipeline_file := c.String("pipeline-file")
	if pipeline == "" && pipeline_file == "" {
		return handeErrOutput(
			"CreatePipelineConfig",
			errors.New("One of '--pipeline-file' or '--pipeline' must be specified."),
		)
	}

	if pipeline != "" && pipeline_file != "" {
		return handeErrOutput(
			"CreatePipelineConfig",
			errors.New("Only one of '--pipeline-file' or '--pipeline' can be specified."),
		)
	}

	if pipeline_file != "" {

	}

	json.Unmarshal([]byte(pipeline), &p)

	pc, r, err := cliAgent().PipelineConfigs.Create(context.Background(), group, p)
	if err != nil {
		return handeErrOutput("CreatePipelineConfig", err)
	}
	return handleOutput(pc, r, "CreatePipelineConfig", err)
}

func CreatePipelineConfigCommand() *cli.Command {
	return &cli.Command{
		Name:   CreatePipelineConfigCommandName,
		Usage:  CreatePipelineConfigCommandUsage,
		Action: CreatePipelineConfigAction,
		Flags: []cli.Flag{
			cli.StringFlag{Name: "group"},
			cli.StringFlag{Name: "pipeline"},
			cli.StringFlag{Name: "pipeline-file"},
		},
	}
}
