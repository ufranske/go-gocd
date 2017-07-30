package main

import (
	"context"
	"encoding/json"
	"github.com/drewsonne/go-gocd/gocd"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
	"io/ioutil"
)

const (
	CreatePipelineConfigCommandName  = "create-pipeline-config"
	CreatePipelineConfigCommandUsage = "Create Pipeline config"
	UpdatePipelineConfigCommandName  = "update-pipeline-config"
	UpdatePipelineConfigCommandUsage = "Update Pipeline config"
)

func CreatePipelineConfigAction(c *cli.Context) error {
	group := c.String("group")
	if group == "" {
		return handleOutput(nil, nil, "CreatePipelineConfig", errors.New("'--group' is missing."))
	}

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

	var pf []byte
	var err error
	if pipeline_file != "" {
		pf, err = ioutil.ReadFile(pipeline_file)
		if err != nil {
			return handeErrOutput("CreatePipelineConfig", err)
		}
	} else {
		pf = []byte(pipeline)
	}
	p := &gocd.Pipeline{}
	err = json.Unmarshal(pf, &p)
	if err != nil {
		return handeErrOutput("CreatePipelineConfig", err)
	}

	pc, r, err := cliAgent().PipelineConfigs.Create(context.Background(), group, p)
	if err != nil {
		return handeErrOutput("CreatePipelineConfig", err)
	}
	return handleOutput(pc, r, "CreatePipelineConfig", err)
}

func UpdatePipelineConfigAction(c *cli.Context) error {
	name := c.String("name")
	if name == "" {
		return handleOutput(nil, nil, "CreatePipelineConfig", errors.New("'--name' is missing."))
	}

	version := c.String("pipeline-version")
	if version == "" {
		return handleOutput(nil, nil, "CreatePipelineConfig", errors.New("'--pipeline-version' is missing."))
	}

	group := c.String("group")
	if group == "" {
		return handleOutput(nil, nil, "CreatePipelineConfig", errors.New("'--group' is missing."))
	}

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

	var pf []byte
	var err error
	if pipeline_file != "" {
		pf, err = ioutil.ReadFile(pipeline_file)
		if err != nil {
			return handeErrOutput("CreatePipelineConfig", err)
		}
	} else {
		pf = []byte(pipeline)
	}
	p := &gocd.Pipeline{}
	err = json.Unmarshal(pf, &p)
	if err != nil {
		return handeErrOutput("CreatePipelineConfig", err)
	}

	pc, r, err := cliAgent().PipelineConfigs.Update(context.Background(), group, name, version, p)
	if err != nil {
		return handeErrOutput("CreatePipelineConfig", err)
	}
	return handleOutput(pc, r, "CreatePipelineConfig", err)

}

func CreatePipelineConfigCommand() *cli.Command {
	return &cli.Command{
		Name:     CreatePipelineConfigCommandName,
		Usage:    CreatePipelineConfigCommandUsage,
		Action:   CreatePipelineConfigAction,
		Category: "Pipeline Configs",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "group"},
			cli.StringFlag{Name: "pipeline"},
			cli.StringFlag{Name: "pipeline-file"},
		},
	}
}

func UpdatePipelineConfigCommand() *cli.Command {
	return &cli.Command{
		Name:     UpdatePipelineConfigCommandName,
		Usage:    UpdatePipelineConfigCommandUsage,
		Action:   UpdatePipelineConfigAction,
		Category: "Pipeline Configs",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "group"},
			cli.StringFlag{Name: "name"},
			cli.StringFlag{Name: "pipeline-version"},
			cli.StringFlag{Name: "pipeline"},
			cli.StringFlag{Name: "pipeline-file"},
		},
	}
}
