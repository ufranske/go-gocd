package cli

import (
	"context"
	"errors"
	"github.com/drewsonne/go-gocd/gocd"
	"github.com/urfave/cli"
	"strings"
)

// List of command name and descriptions
const (
	ListEnvironmentsCommandName                = "list-environments"
	ListEnvironmentsCommandUsage               = "List all environments"
	GetEnvironmentCommandName                  = "get-environment"
	GetEnvironmentCommandUsage                 = "Get an environment by name"
	AddPipelinesToEnvironmentCommandName       = "add-pipelines-to-environment"
	AddPipelinesToEnvironmentCommandUsage      = "Add one or more pipelines to an environment"
	RemovePipelinesFromEnvironmentCommandName  = "remove-pipelines-from-environment"
	RemovePipelinesFromEnvironmentCommandUsage = "Remove one or more pipelines from an environment"
)

// ListEnvironmentsAction handles the listing of environments
func listEnvironmentsAction(c *cli.Context) cli.ExitCoder {
	es, r, err := cliAgent(c).Environments.List(context.Background())
	if err == nil {
		es.RemoveLinks()
	} else {
		return NewCliError("ListEnvironments", r, err)
	}

	return handleOutput(es, "ListEnvironments")
}

// GetEnvironmentAction handles the retrieval of environments
func getEnvironmentAction(c *cli.Context) cli.ExitCoder {
	var name string
	if name = c.String("name"); name == "" {
		return NewCliError("GetEnvironment", nil, errors.New("'--name' is missing"))
	}
	e, r, err := cliAgent(c).Environments.Get(context.Background(), name)
	if err == nil {
		e.RemoveLinks()
	} else {
		return NewCliError("GetEnvironment", r, err)
	}
	return handleOutput(e, "GetEnvironment")
}

// AddPipelinesToEnvironmentAction handles the adding of a pipeline to an environment
func addPipelinesToEnvironmentAction(c *cli.Context) cli.ExitCoder {
	var environment, pipelines string

	if environment = c.String("environment-name"); environment == "" {
		return NewCliError("AddPipelinesToEnvironment", nil, errors.New("'--environment-name' is missing"))
	}
	if pipelines = c.String("pipeline-names"); pipelines == "" {
		return NewCliError("AddPipelinesToEnvironment", nil, errors.New("'--pipeline-names' is missing"))
	}

	e, r, err := cliAgent(c).Environments.Patch(context.Background(), environment, &gocd.EnvironmentPatchRequest{
		Pipelines: &gocd.PatchStringAction{
			Add: strings.Split(pipelines, ","),
		},
	})
	if err == nil {
		e.RemoveLinks()
	} else {
		return NewCliError("AddPipelinesToEnvironment", r, err)
	}
	return handleOutput(e, "AddPipelinesToEnvironment")
}

// RemovePipelinesFromEnvironmentAction handles the removing of a pipeline from an environment
func removePipelinesFromEnvironmentAction(c *cli.Context) cli.ExitCoder {
	var environment, pipelines string

	if environment = c.String("environment-name"); environment == "" {
		return NewCliError("RemovePipelinesFromEnvironment", nil, errors.New("'--environment-name' is missing"))
	}
	if pipelines = c.String("pipeline-names"); pipelines == "" {
		return NewCliError("RemovePipelinesFromEnvironment", nil, errors.New("'--pipeline-names' is missing"))
	}

	e, r, err := cliAgent(c).Environments.Patch(context.Background(), environment, &gocd.EnvironmentPatchRequest{
		Pipelines: &gocd.PatchStringAction{
			Remove: strings.Split(pipelines, ","),
		},
	})
	if err == nil {
		e.RemoveLinks()
	} else {
		return NewCliError("RemovePipelinesFromEnvironment", r, err)
	}
	return handleOutput(e, "RemovePipelinesFromEnvironment")
}

// ListEnvironmentsCommand handles definition of cli command
func listEnvironmentsCommand() *cli.Command {
	return &cli.Command{
		Name:     ListEnvironmentsCommandName,
		Usage:    ListEnvironmentsCommandUsage,
		Action:   listEnvironmentsAction,
		Category: "Environments",
	}
}

// GetEnvironmentCommand handles definition of cli command
func getEnvironmentCommand() *cli.Command {
	return &cli.Command{
		Name:     GetEnvironmentCommandName,
		Usage:    GetEnvironmentCommandUsage,
		Action:   getEnvironmentAction,
		Category: "Environments",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "name"},
		},
	}
}

// AddPipelinesToEnvironmentCommand handles definition of cli command
func addPipelinesToEnvironmentCommand() *cli.Command {
	return &cli.Command{
		Name:     AddPipelinesToEnvironmentCommandName,
		Usage:    AddPipelinesToEnvironmentCommandUsage,
		Action:   addPipelinesToEnvironmentAction,
		Category: "Environments",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "environment-name"},
			cli.StringFlag{
				Name:  "pipeline-names",
				Usage: "Comma seperated list of pipeline names to add.",
			},
		},
	}
}

// RemovePipelinesFromEnvironmentCommand handles definition of cli command
func removePipelinesFromEnvironmentCommand() *cli.Command {
	return &cli.Command{
		Name:     RemovePipelinesFromEnvironmentCommandName,
		Usage:    RemovePipelinesFromEnvironmentCommandUsage,
		Action:   removePipelinesFromEnvironmentAction,
		Category: "Environments",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name: "environment-name",
			},
			cli.StringFlag{
				Name:  "pipeline-names",
				Usage: "Comma seperated list of pipeline names to remove.",
			},
		},
	}
}
