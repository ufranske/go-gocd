package cli

import (
	"context"
	"encoding/json"
	"github.com/drewsonne/go-gocd/gocd"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
	"io/ioutil"
)

// List of command name and descriptions
const (
	CreatePipelineConfigCommandName  = "create-pipeline-config"
	CreatePipelineConfigCommandUsage = "Create Pipeline config"
	UpdatePipelineConfigCommandName  = "update-pipeline-config"
	UpdatePipelineConfigCommandUsage = "Update Pipeline config"
	DeletePipelineConfigCommandName  = "delete-pipeline-config"
	DeletePipelineConfigCommandUsage = "Remove Pipeline. This will not delete the pipeline history, which will be stored in the database"
	GetPipelineConfigCommandName     = "get-pipeline-config"
	GetPipelineConfigCommandUsage    = "Get a Pipeline Configuration"
)

// CreatePipelineConfigAction handles the interaction between the cli flags and the action handler for
// create-pipeline-config-action
func createPipelineConfigAction(c *cli.Context) cli.ExitCoder {
	group := c.String("group")
	if group == "" {
		return NewCliError("CreatePipelineConfig", nil, errors.New("'--group' is missing"))
	}

	pipeline := c.String("pipeline")
	pipelineFile := c.String("pipeline-file")
	if pipeline == "" && pipelineFile == "" {
		return NewCliError(
			"CreatePipelineConfig", nil,
			errors.New("One of '--pipeline-file' or '--pipeline' must be specified"),
		)
	}

	if pipeline != "" && pipelineFile != "" {
		return NewCliError(
			"CreatePipelineConfig", nil,
			errors.New("Only one of '--pipeline-file' or '--pipeline' can be specified"),
		)
	}

	var pf []byte
	var err error
	if pipelineFile != "" {
		pf, err = ioutil.ReadFile(pipelineFile)
		if err != nil {
			return NewCliError("CreatePipelineConfig", nil, err)
		}
	} else {
		pf = []byte(pipeline)
	}
	p := &gocd.Pipeline{}
	err = json.Unmarshal(pf, &p)
	if err != nil {
		return NewCliError("CreatePipelineConfig", nil, err)
	}

	pc, r, err := cliAgent(c).PipelineConfigs.Create(context.Background(), group, p)
	if err != nil {
		return NewCliError("CreatePipelineConfig", r, err)
	}
	return handleOutput(pc, "CreatePipelineConfig")
}

// UpdatePipelineConfigAction handles the interaction between the cli flags and the action handler for
// update-pipeline-config-action
func updatePipelineConfigAction(c *cli.Context) cli.ExitCoder {
	name := c.String("name")
	if name == "" {
		return NewCliError("UpdatePipelineConfig", nil, errors.New("'--name' is missing"))
	}

	version := c.String("pipeline-version")
	if version == "" {
		return NewCliError("UpdatePipelineConfig", nil, errors.New("'--pipeline-version' is missing"))
	}

	pipeline := c.String("pipeline")
	pipelineFile := c.String("pipeline-file")
	if pipeline == "" && pipelineFile == "" {
		return NewCliError(
			"UpdatePipelineConfig", nil,
			errors.New("One of '--pipeline-file' or '--pipeline' must be specified"),
		)
	}

	if pipeline != "" && pipelineFile != "" {
		return NewCliError(
			"UpdatePipelineConfig", nil,
			errors.New("Only one of '--pipeline-file' or '--pipeline' can be specified"),
		)
	}

	var pf []byte
	var err error
	if pipelineFile != "" {
		pf, err = ioutil.ReadFile(pipelineFile)
		if err != nil {
			return NewCliError("UpdatePipelineConfig", nil, err)
		}
	} else {
		pf = []byte(pipeline)
	}
	p := &gocd.Pipeline{
		Version: version,
	}
	err = json.Unmarshal(pf, &p)
	if err != nil {
		return NewCliError("UpdatePipelineConfig", nil, err)
	}

	pc, r, err := cliAgent(c).PipelineConfigs.Update(context.Background(), name, p)
	if err != nil {
		return NewCliError("CreatePipelineConfig", r, err)
	}
	return handleOutput(pc, "CreatePipelineConfig")

}

// DeletePipelineConfigAction handles the interaction between the cli flags and the action handler for
// delete-pipeline-config-action
func deletePipelineConfigAction(c *cli.Context) cli.ExitCoder {
	name := c.String("name")
	if name == "" {
		return NewCliError("CreatePipelineConfig", nil, errors.New("'--name' is missing"))
	}

	deleteResponse, r, err := cliAgent(c).PipelineConfigs.Delete(context.Background(), name)
	if r.HTTP.StatusCode == 406 {
		err = errors.New(deleteResponse)
	}
	if err != nil {
		return NewCliError("CreatePipelineConfig", r, err)
	}
	return handleOutput(deleteResponse, "DeletePipelineTemplate")
}

// GetPipelineConfigAction handles the interaction between the cli flags and the action handler for get-pipeline-config
func getPipelineConfigAction(c *cli.Context) cli.ExitCoder {
	name := c.String("name")
	if name == "" {
		return NewCliError("GetPipelineConfig", nil, errors.New("'--name' is missing"))
	}

	getResponse, r, err := cliAgent(c).PipelineConfigs.Get(context.Background(), name)
	if r.HTTP.StatusCode != 404 {
		getResponse.RemoveLinks()
	}
	if err != nil {
		return NewCliError("GetPipelineConfig", r, err)
	}

	return handleOutput(getResponse, "GetPipelineConfig")
}

// CreatePipelineConfigCommand handles the interaction between the cli flags and the action handler for create-pipeline-config
func createPipelineConfigCommand() *cli.Command {
	return &cli.Command{
		Name:     CreatePipelineConfigCommandName,
		Usage:    CreatePipelineConfigCommandUsage,
		Action:   createPipelineConfigAction,
		Category: "Pipeline Configs",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "group"},
			cli.StringFlag{Name: "pipeline"},
			cli.StringFlag{Name: "pipeline-file"},
		},
	}
}

// UpdatePipelineConfigCommand handles the interaction between the cli flags and the action handler for update-pipeline-config
func updatePipelineConfigCommand() *cli.Command {
	return &cli.Command{
		Name:     UpdatePipelineConfigCommandName,
		Usage:    UpdatePipelineConfigCommandUsage,
		Action:   updatePipelineConfigAction,
		Category: "Pipeline Configs",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "name"},
			cli.StringFlag{Name: "pipeline-version"},
			cli.StringFlag{Name: "pipeline"},
			cli.StringFlag{Name: "pipeline-file"},
		},
	}
}

// DeletePipelineConfigCommand handles the interaction between the cli flags and the action handler for delete-pipeline-config
func deletePipelineConfigCommand() *cli.Command {
	return &cli.Command{
		Name:     DeletePipelineConfigCommandName,
		Usage:    DeletePipelineConfigCommandUsage,
		Category: "Pipeline Configs",
		Action:   deletePipelineConfigAction,
		Flags: []cli.Flag{
			cli.StringFlag{Name: "name"},
		},
	}
}

// GetPipelineConfigCommand handles the interaction between the cli flags and the action handler for get-pipeline-config
func getPipelineConfigCommand() *cli.Command {
	return &cli.Command{
		Name:     GetPipelineConfigCommandName,
		Usage:    GetPipelineConfigCommandUsage,
		Action:   getPipelineConfigAction,
		Category: "Pipeline Configs",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "name"},
		},
	}
}
