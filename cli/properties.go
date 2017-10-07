package cli

import (
	"context"
	"errors"
	"github.com/drewsonne/go-gocd/gocd"
	"github.com/urfave/cli"
)

const (
	listPropertiesCommandName  = "list-properties"
	listPropertiesCommandUsage = "List the properties for a given job."
	createPropertyCommandName  = "create-property"
	createPropertyCommandUsage = "Create a property for a given job."
	propertiesGroup            = "Properties"
)

func createPropertyAction(client *gocd.Client, c *cli.Context) (r interface{}, resp *gocd.APIResponse, err error) {
	var name, value, pipeline, stage string

	if name = c.String("name"); name == "" {
		return nil, nil, NewFlagError("name")
	}

	if value = c.String("value"); value == "" {
		return nil, nil, NewFlagError("value")
	}

	if pipeline = c.String("pipeline"); pipeline == "" {
		return nil, nil, NewFlagError("pipeline")
	}

	if stage = c.String("stage"); stage == "" {
		return nil, nil, NewFlagError("stage")
	}

	pipelineCounter := c.Int("pipeline-counter")
	stageCounter := c.Int("stage-counter")

	return client.Properties.Create(context.Background(), name, value, &gocd.PropertyRequest{
		Pipeline:        pipeline,
		PipelineCounter: pipelineCounter,
		Stage:           stage,
		StageCounter:    stageCounter,
	})
}

func listPropertiesAction(client *gocd.Client, c *cli.Context) (r interface{}, resp *gocd.APIResponse, err error) {
	var pipeline, stage string

	if pipeline = c.String("pipeline"); pipeline == "" {
		return nil, nil, NewFlagError("pipeline")
	}

	if stage = c.String("stage"); stage == "" {
		return nil, nil, NewFlagError("stage")
	}

	stageCounter := c.Int("stage-counter")
	pipelineCounter := c.Int("pipeline-counter")

	return client.Properties.List(context.Background(), &gocd.PropertyRequest{
		Pipeline:        pipeline,
		PipelineCounter: pipelineCounter,
		Stage:           stage,
		StageCounter:    stageCounter,
	})
}

func listPropertiesCommand() *cli.Command {
	return &cli.Command{
		Name:     listPropertiesCommandName,
		Usage:    listPropertiesCommandUsage,
		Action:   actionWrapper(listPropertiesAction),
		Category: propertiesGroup,
		Flags: []cli.Flag{
			cli.StringFlag{Name: "pipeline", EnvVar: "GO_PIPELINE_NAME"},
			cli.IntFlag{Name: "pipeline-counter", EnvVar: "GO_PIPELINE_COUNTER"},
			cli.StringFlag{Name: "stage", EnvVar: "GO_STAGE_NAME"},
			cli.IntFlag{Name: "stage-counter", EnvVar: "GO_STAGE_COUNTER"},
			cli.StringFlag{Name: "job", EnvVar: "GO_JOB_NAME"},
		},
	}
}

func createPropertyCommand() *cli.Command {
	return &cli.Command{
		Name:     createPropertyCommandName,
		Usage:    createPropertyCommandUsage,
		Action:   actionWrapper(createPropertyAction),
		Category: propertiesGroup,
		Flags: []cli.Flag{
			cli.StringFlag{Name: "name"},
			cli.StringFlag{Name: "value"},
			cli.StringFlag{Name: "pipeline", EnvVar: "GO_PIPELINE_NAME"},
			cli.IntFlag{Name: "pipeline-counter", EnvVar: "GO_PIPELINE_COUNTER"},
			cli.StringFlag{Name: "stage", EnvVar: "GO_STAGE_NAME"},
			cli.IntFlag{Name: "stage-counter", EnvVar: "GO_STAGE_COUNTER"},
			cli.StringFlag{Name: "job", EnvVar: "GO_JOB_NAME"},
		},
	}
}
