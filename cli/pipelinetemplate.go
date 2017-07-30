package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/drewsonne/go-gocd/gocd"
	"github.com/urfave/cli"
)

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

func ListPipelineTemplatesAction(c *cli.Context) error {
	ts, r, err := cliAgent().PipelineTemplates.List(context.Background())
	if err != nil {
		return handleOutput(nil, r, "ListPipelineTemplates", err)
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
	return handleOutput(responses, r, "ListPipelineTemplates", err)
}

func GetPipelineTemplateAction(c *cli.Context) error {
	if c.String("template-name") == "" {
		return handleOutput(nil, nil, "GetPipelineTemplate", errors.New("'--template-name' is missing."))
	}

	pt, r, err := cliAgent().PipelineTemplates.Get(context.Background(), c.String("template-name"))
	if r.Http.StatusCode != 404 {
		pt.RemoveLinks()
	}
	return handleOutput(pt, r, "GetPipelineTemplate", err)
}

func CreatePipelineTemplateAction(c *cli.Context) error {
	if c.String("template-name") == "" {
		return handleOutput(nil, nil, "CreatePipelineTemplate", errors.New("'--template-name' is missing."))
	}
	if len(c.StringSlice("stage")) < 1 {
		return handleOutput(nil, nil, "CreatePipelineTemplate", errors.New("At least 1 '--stage' must be set."))
	}

	stages := []*gocd.Stage{}
	for _, stage := range c.StringSlice("stage") {
		st := gocd.Stage{}
		json.Unmarshal([]byte(stage), &st)

		if err := st.Validate(); err != nil {
			return handleOutput(nil, nil, "CreatePipelineTemplate", err)
		}
		stages = append(stages, &st)
	}

	pt, r, err := cliAgent().PipelineTemplates.Create(context.Background(), c.String("template-name"), stages)
	return handleOutput(pt, r, "CreatePipelineTemplate", err)
}

func UpdatePipelineTemplateAction(c *cli.Context) error {
	if c.String("template-name") == "" {
		return handeErrOutput("UpdatePipelineTemplate", errors.New("'--template-name' is missing."))
	}
	if c.String("template-version") == "" {
		return handeErrOutput("UpdatePipelineTemplate", errors.New("'--version' is missing."))
	}
	if len(c.StringSlice("stage")) < 1 {
		return handeErrOutput("UpdatePipelineTemplate", errors.New("At least 1 '--stage' must be set."))
	}

	stages := []*gocd.Stage{}
	for _, stage := range c.StringSlice("stage") {
		st := gocd.Stage{}
		json.Unmarshal([]byte(stage), &st)

		if err := st.Validate(); err != nil {
			return handeErrOutput("UpdatePipelineTemplate", err)
		}
		stages = append(stages, &st)
	}

	pt, r, err := cliAgent().PipelineTemplates.Update(context.Background(), c.String("template-name"), c.String("template-version"), stages)
	return handleOutput(pt, r, "UpdatePipelineTemplate", err)
}

func DeletePipelineTemplateAction(c *cli.Context) error {
	if c.String("template-name") == "" {
		return handleOutput(nil, nil, "DeletePipelineTemplate", errors.New("'--template-name' is missing."))
	}

	deleteResponse, r, err := cliAgent().PipelineTemplates.Delete(context.Background(), c.String("template-name"))
	if r.Http.StatusCode == 406 {
		err = errors.New(deleteResponse)
	}
	return handleOutput(deleteResponse, r, "DeletePipelineTemplate", err)
}

func ListPipelineTemplatesCommand() *cli.Command {
	return &cli.Command{
		Name:     ListPipelineTemplatesCommandName,
		Usage:    ListPipelineTemplatesCommandUsage,
		Action:   ListPipelineTemplatesAction,
		Category: "Pipeline Templates",
	}
}

func GetPipelineTemplateCommand() *cli.Command {
	return &cli.Command{
		Name:     GetPipelineTemplateCommandName,
		Usage:    GetPipelineTemplateCommandUsage,
		Action:   GetPipelineTemplateAction,
		Category: "Pipeline Templates",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "template-name", Usage: "Name of the Pipeline Template configuration."},
		},
	}
}

func CreatePipelineTemplateCommand() *cli.Command {
	return &cli.Command{
		Name:     CreatePipelineTemplateCommandName,
		Usage:    CreatePipelineTemplateCommandUsage,
		Action:   CreatePipelineTemplateAction,
		Category: "Pipeline Templates",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "template-name", Usage: "Pipeline Template name."},
			cli.StringSliceFlag{Name: "stage", Usage: "JSON encoded stage object."},
		},
	}
}

func UpdatePipelineTemplateCommand() *cli.Command {
	return &cli.Command{
		Name:     UpdatePipelineTemplateCommandName,
		Usage:    UpdatePipelineTemplateCommandUsage,
		Action:   UpdatePipelineTemplateAction,
		Category: "Pipeline Templates",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "template-version", Usage: "Pipeline template version."},
			cli.StringFlag{Name: "template-name", Usage: "Pipeline Template name."},
			cli.StringSliceFlag{Name: "stage", Usage: "JSON encoded stage object."},
		},
	}
}
func DeletePipelineTemplateCommand() *cli.Command {
	return &cli.Command{
		Name:     DeletePipelineTemplateCommandName,
		Usage:    DeletePipelineTemplateCommandUsage,
		Action:   DeletePipelineTemplateAction,
		Category: "Pipeline Templates",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "template-name", Usage: "Pipeline Template name."},
		},
	}
}
