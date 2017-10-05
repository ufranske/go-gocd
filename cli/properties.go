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
	name := c.String("name")
	if name == "" {
		return nil, nil, errors.New("'--name' is missing")
	}

	value := c.String("value")
	if value == "" {
		return nil, nil, errors.New("'--value' is missing")
	}

	pipeline := c.String("pipeline")
	if pipeline == "" {
		return nil, nil, errors.New("'--pipeline' is missing")
	}
	pipelineCounter := c.Int("pipeline-counter")

	stage := c.String("stage")
	if stage == "" {
		return nil, nil, errors.New("'--stage' is missing")
	}
	stageCounter := c.Int("stage-counter")

	return client.Properties.Create(context.Background(), name, value, &gocd.PropertyRequest{
		Pipeline:        pipeline,
		PipelineCounter: pipelineCounter,
		Stage:           stage,
		StageCounter:    stageCounter,
	})
}

func listPropertiesAction(client *gocd.Client, c *cli.Context) (r interface{}, resp *gocd.APIResponse, err error) {

	pipeline := c.String("pipeline")
	if pipeline == "" {
		return nil, nil, errors.New("'--pipeline' is missing")
	}
	pipelineCounter := c.Int("pipeline-counter")

	stage := c.String("stage")
	if stage == "" {
		return nil, nil, errors.New("'--stage' is missing")
	}
	stageCounter := c.Int("stage-counter")

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
		Action:   listPropertiesAction,
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
		Action:   createPropertyAction,
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
