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

func createPropertyAction(c *cli.Context) cli.ExitCoder {
	name := c.String("name")
	if name == "" {
		return NewCliError(propertiesGroup, nil, errors.New("'--name' is missing"))
	}

	value := c.String("value")
	if value == "" {
		return NewCliError(propertiesGroup, nil, errors.New("'--value' is missing"))
	}

	pipeline := c.String("pipeline")
	if pipeline == "" {
		return NewCliError(propertiesGroup, nil, errors.New("'--pipeline' is missing"))
	}
	pipelineCounter := c.Int("pipeline-counter")

	stage := c.String("stage")
	if stage == "" {
		return NewCliError("Properties", nil, errors.New("'--stage' is missing"))
	}
	stageCounter := c.Int("stage-counter")

	created, r, err := cliAgent(c).Properties.Create(context.Background(), name, value, &gocd.PropertyRequest{
		Pipeline:        pipeline,
		PipelineCounter: pipelineCounter,
		Stage:           stage,
		StageCounter:    stageCounter,
	})
	if err != nil {
		return NewCliError(propertiesGroup, r, err)
	}
	return handleOutput(created, "ListProperties")
}

func listPropertiesAction(c *cli.Context) cli.ExitCoder {

	pipeline := c.String("pipeline")
	if pipeline == "" {
		return NewCliError(propertiesGroup, nil, errors.New("'--pipeline' is missing"))
	}
	pipelineCounter := c.Int("pipeline-counter")

	stage := c.String("stage")
	if stage == "" {
		return NewCliError(propertiesGroup, nil, errors.New("'--stage' is missing"))
	}
	stageCounter := c.Int("stage-counter")

	properties, r, err := cliAgent(c).Properties.List(context.Background(), &gocd.PropertyRequest{
		Pipeline:        pipeline,
		PipelineCounter: pipelineCounter,
		Stage:           stage,
		StageCounter:    stageCounter,
	})
	if err != nil {
		return NewCliError("ListProperties", r, err)
	}
	return handleOutput(properties, "ListProperties")
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
