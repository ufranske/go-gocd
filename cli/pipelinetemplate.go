package cli

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/drewsonne/go-gocd/gocd"
	"github.com/urfave/cli"
)

// List of command name and descriptions
const (
	ListPipelineTemplatesCommandName   = "list-pipeline-templates"
	ListPipelineTemplatesCommandUsage  = "List Pipeline Templates"
	GetPipelineTemplateCommandName     = "get-pipeline-template"
	GetPipelineTemplateCommandUsage    = "Get Pipeline Templates"
	CreatePipelineTemplateCommandName  = "create-pipeline-template"
	CreatePipelineTemplateCommandUsage = "Create Pipeline Templates"
	UpdatePipelineTemplateCommandName  = "update-pipeline-template"
	UpdatePipelineTemplateCommandUsage = "Update Pipeline template"
	DeletePipelineTemplateCommandName  = "delete-pipeline-template"
	DeletePipelineTemplateCommandUsage = "Delete Pipeline template"
)

// ListPipelineTemplatesAction lists all pipeline templates.
func listPipelineTemplatesAction(c *cli.Context) cli.ExitCoder {
	ts, r, err := cliAgent(c).PipelineTemplates.List(context.Background())
	if err != nil {
		return NewCliError("ListPipelineTemplates", r, err)
	}

	type p struct {
		Name string `json:"name"`
	}

	type ptr struct {
		Name      string `json:"name"`
		Pipelines []*p   `json:"pipelines"`
	}
	responses := []ptr{}
	for _, pt := range ts {
		pt.RemoveLinks()
		ps := []*p{}
		for _, pipe := range pt.Pipelines() {
			ps = append(ps, &p{pipe.Name})
		}

		responses = append(responses, ptr{
			Name:      pt.Name,
			Pipelines: ps,
		})
	}
	return handleOutput(responses, "ListPipelineTemplates")
}

// GetPipelineTemplateAction checks template-name is provided, and that the response is 2xx.
func getPipelineTemplateAction(c *cli.Context) cli.ExitCoder {
	var name string
	if name = c.String("template-name"); name == "" {
		return NewCliError("GetPipelineTemplate", nil, errors.New("'--template-name' is missing"))
	}

	pt, r, err := cliAgent(c).PipelineTemplates.Get(context.Background(), name)
	if r.HTTP.StatusCode != 404 {
		pt.RemoveLinks()
	}
	if err != nil {
		return NewCliError("GetPipelineTemplate", r, err)
	}
	return handleOutput(pt, "GetPipelineTemplate")
}

// CreatePipelineTemplateAction checks stages and template-name is provided, and that the response is 2xx.
func createPipelineTemplateAction(c *cli.Context) cli.ExitCoder {
	if c.String("template-name") == "" {
		return NewCliError("CreatePipelineTemplate", nil, errors.New("'--template-name' is missing"))
	}
	if len(c.StringSlice("stage")) < 1 {
		return NewCliError("CreatePipelineTemplate", nil, errors.New("At least 1 '--stage' must be set"))
	}

	stages := []*gocd.Stage{}
	for _, stage := range c.StringSlice("stage") {
		st := gocd.Stage{}
		json.Unmarshal([]byte(stage), &st)

		if err := st.Validate(); err != nil {
			return NewCliError("CreatePipelineTemplate", nil, err)
		}
		stages = append(stages, &st)
	}

	pt, r, err := cliAgent(c).PipelineTemplates.Create(context.Background(), c.String("template-name"), stages)
	if err != nil {
		return NewCliError("CreatePipelineTemplate", r, err)
	}
	return handleOutput(pt, "CreatePipelineTemplate")
}

// UpdatePipelineTemplateAction checks stages, template-name and template-version is provided, and that the response is
// 2xx.
func updatePipelineTemplateAction(c *cli.Context) cli.ExitCoder {
	if c.String("template-name") == "" {
		return NewCliError("UpdatePipelineTemplate", nil, errors.New("'--template-name' is missing"))
	}
	if c.String("template-version") == "" {
		return NewCliError("UpdatePipelineTemplate", nil, errors.New("'--version' is missing"))
	}
	if len(c.StringSlice("stage")) < 1 {
		return NewCliError("UpdatePipelineTemplate", nil, errors.New("At least 1 '--stage' must be set"))
	}

	stages := []*gocd.Stage{}
	for _, stage := range c.StringSlice("stage") {
		st := gocd.Stage{}
		json.Unmarshal([]byte(stage), &st)

		if err := st.Validate(); err != nil {
			return NewCliError("UpdatePipelineTemplate", nil, err)
		}
		stages = append(stages, &st)
	}

	ptr := gocd.PipelineTemplate{
		Version: c.String("template-version"),
		Stages:  stages,
	}

	pt, r, err := cliAgent(c).PipelineTemplates.Update(context.Background(), c.String("template-name"), &ptr)
	if err != nil {
		return NewCliError("UpdatePipelineTemplate", r, err)
	}
	return handleOutput(pt, "UpdatePipelineTemplate")
}

// DeletePipelineTemplateCommand handles the interaction between the cli flags and the action handler for
// delete-pipeline-template and checks a template-name is provided and that the response is a 2xx response.
func deletePipelineTemplateCommand() *cli.Command {
	return &cli.Command{
		Name:     DeletePipelineTemplateCommandName,
		Usage:    DeletePipelineTemplateCommandUsage,
		Category: "Pipeline Templates",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "template-name", Usage: "Pipeline Template name."}},
		Action: func(c *cli.Context) cli.ExitCoder {
			if c.String("template-name") == "" {
				return NewCliError("DeletePipelineTemplate", nil, errors.New("'--template-name' is missing"))
			}

			deleteResponse, r, err := cliAgent(c).PipelineTemplates.Delete(context.Background(), c.String("template-name"))
			if r.HTTP.StatusCode == 406 {
				err = errors.New(deleteResponse)
			}
			if err != nil {
				return NewCliError("DeletePipelineTemplate", r, err)
			}
			return handleOutput(deleteResponse, "DeletePipelineTemplate")
		},
	}
}

// ListPipelineTemplatesCommand handles the interaction between the cli flags and the action handler for
// list-pipeline-templates
func listPipelineTemplatesCommand() *cli.Command {
	return &cli.Command{
		Name:     ListPipelineTemplatesCommandName,
		Usage:    ListPipelineTemplatesCommandUsage,
		Action:   listPipelineTemplatesAction,
		Category: "Pipeline Templates",
	}
}

// GetPipelineTemplateCommand handles the interaction between the cli flags and the action handler for
// get-pipeline-template
func getPipelineTemplateCommand() *cli.Command {
	return &cli.Command{
		Name:     GetPipelineTemplateCommandName,
		Usage:    GetPipelineTemplateCommandUsage,
		Action:   getPipelineTemplateAction,
		Category: "Pipeline Templates",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "template-name", Usage: "Name of the Pipeline Template configuration."},
		},
	}
}

// CreatePipelineTemplateCommand handles the interaction between the cli flags and the action handler for
// create-pipeline-template
func createPipelineTemplateCommand() *cli.Command {
	return &cli.Command{
		Name:     CreatePipelineTemplateCommandName,
		Usage:    CreatePipelineTemplateCommandUsage,
		Action:   createPipelineTemplateAction,
		Category: "Pipeline Templates",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "template-name", Usage: "Pipeline Template name."},
			cli.StringSliceFlag{Name: "stage", Usage: "JSON encoded stage object."},
		},
	}
}

// UpdatePipelineTemplateCommand handles the interaction between the cli flags and the action handler for
// update-pipeline-template
func updatePipelineTemplateCommand() *cli.Command {
	return &cli.Command{
		Name:     UpdatePipelineTemplateCommandName,
		Usage:    UpdatePipelineTemplateCommandUsage,
		Action:   updatePipelineTemplateAction,
		Category: "Pipeline Templates",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "template-version", Usage: "Pipeline template version."},
			cli.StringFlag{Name: "template-name", Usage: "Pipeline Template name."},
			cli.StringSliceFlag{Name: "stage", Usage: "JSON encoded stage object."},
		},
	}
}
